/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * Icaro is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Icaro is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Icaro.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	gomail "gopkg.in/gomail.v2"

	"sun-api/configuration"
	"sun-api/database"
	"sun-api/models"
)


func GetSessionByKeyAndUnitId(key string, unitId int) models.Session {
        var session models.Session
        db := database.Database()
        db.Where("session_key = ? and unit_id = ?", key, unitId).First(&session)
        db.Close()

	return session;
}

func GetDeviceByMacAddress(mac string) models.Device {
        var unit models.Device
        db := database.Database()
        db.Where("mac_address = ?", mac).First(&unit)
        db.Close()

        return unit
}

func GetUnitByMacAddress(mac string) models.Unit {
        var unit models.Unit
        db := database.Database()
        db.Where("mac_address = ?", mac).First(&unit)
        db.Close()

        return unit
}

func GetUserByNameAndHotspotId(name string, hotspotId int) models.User {
        var unit models.User
        db := database.Database()
        db.Where("username = ? and hotspot_id = ?",name,hotspotId).First(&unit)
        db.Close()

        return unit
}

func GetUnitByUuid(uuid string) models.Unit {
	var unit models.Unit
	db := database.Database()
	db.Where("uuid = ?", uuid).First(&unit)
	db.Close()

	return unit
}

func GetUserByUsername(id string) models.User {
	var user models.User
	db := database.Database()
	db.Where("username = ?", id).First(&user)
	db.Close()

	return user
}

func CalcDigest(unit models.Unit) string {
	h := md5.New()
	io.WriteString(h, unit.Secret+unit.Uuid)
	digest := fmt.Sprintf("%x", h.Sum(nil))

	return digest
}

func GenerateCode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func SendSMSCode(number string, code string) {
	// retrieve account info and token
	accountSid := configuration.Config.Endpoints.Sms.AccountSid
	authToken := configuration.Config.Endpoints.Sms.AuthToken
	urlAPI := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// compose message data
	msgData := url.Values{}
	msgData.Set("To", number)
	msgData.Set("From", configuration.Config.Endpoints.Sms.Number)
	msgData.Set("Body", "Icaro - SMS Login code: "+code) // TODO: get message from hotspot preferences
	msgDataReader := *strings.NewReader(msgData.Encode())

	// create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlAPI, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(resp.Status)
	}
}

func SendEmailCode(email string, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", configuration.Config.Endpoints.Email.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Icaro")
	m.SetBody("text/html", "Email Login code: "+code)

	d := gomail.NewDialer(
		configuration.Config.Endpoints.Email.SMTPHost,
		configuration.Config.Endpoints.Email.SMTPPort,
		configuration.Config.Endpoints.Email.SMTPUser,
		configuration.Config.Endpoints.Email.SMTPPassword,
	)

	// send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
	}
}

func Contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}
