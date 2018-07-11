/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * Icaro is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Icaro is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Icaro.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package methods

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/nethesis/icaro/sun/sun-api/configuration"
	"github.com/nethesis/icaro/sun/sun-api/models"
	"github.com/nethesis/icaro/wax/utils"
)

func GetWingsPrefs(c *gin.Context) {
	uuid := c.Query("uuid")

	// extract unit info
	unit := utils.GetUnitByUuid(uuid)

	// extract hotspot info
	hotspot := utils.GetHotspotById(unit.HotspotId)

	// get hotspot preferences
	prefs := utils.GetHotspotPreferences(hotspot.Id)

	var wingsPrefs models.WingsPrefs
	wingsPrefs.HotspotId = hotspot.Id
	wingsPrefs.HotspotName = hotspot.Name

	prefsMap := make(map[string]string)
	for i := 0; i < len(prefs); i++ {
		prefsMap[prefs[i].Key] = prefs[i].Value
	}
	wingsPrefs.Preferences = prefsMap

	// get social ids
	wingsPrefs.Socials.FacebookClientId = configuration.Config.AuthSocial.Facebook.ClientId
	wingsPrefs.Socials.LinkedInClientId = configuration.Config.AuthSocial.LinkedIn.ClientId
	wingsPrefs.Socials.InstagramClientId = configuration.Config.AuthSocial.Instagram.ClientId

	// disclaimers
	terms := configuration.Config.Disclaimers.TermsOfUse
	terms = strings.Replace(terms, "$$COMPANY_NAME$$", hotspot.BusinessName, -1)
	terms = strings.Replace(terms, "$$COMPANY_VAT$$", hotspot.BusinessVAT, -1)
	terms = strings.Replace(terms, "$$COMPANY_ADDRESS$$", hotspot.BusinessAddress, -1)
	terms = strings.Replace(terms, "$$COMPANY_EMAIL$$", hotspot.BusinessEmail, -1)

	marketings := configuration.Config.Disclaimers.MarketingUse
	marketings = strings.Replace(marketings, "$$COMPANY_NAME$$", hotspot.BusinessName, -1)
	marketings = strings.Replace(marketings, "$$COMPANY_VAT$$", hotspot.BusinessVAT, -1)
	marketings = strings.Replace(marketings, "$$COMPANY_ADDRESS$$", hotspot.BusinessAddress, -1)
	marketings = strings.Replace(marketings, "$$COMPANY_EMAIL$$", hotspot.BusinessEmail, -1)

	wingsPrefs.Disclaimers.TermsOfUse = terms
	wingsPrefs.Disclaimers.MarketingUse = marketings

	c.JSON(http.StatusOK, wingsPrefs)
}
