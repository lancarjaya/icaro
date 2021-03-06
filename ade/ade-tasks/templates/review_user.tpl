<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <style>
        h1 {
            color: {{ .TitleColor }};
            border: none;
            padding: 0 0;
            font-weight: 700;
            text-transform: none;
        }

        h2 {
            border: none;
            padding: 0 0;
            font-weight: 700;
            text-transform: none;
            margin: 20px;
        }

        h3 {
            border: none;
            padding: 0 0;
            font-weight: 700;
            text-transform: none;
        }

        h4 {
            border: none;
            padding: 0 0;
            font-weight: 700;
            text-transform: none;
        }

        h5 {
            border: none;
            padding: 0 0;
            font-weight: 700;
            text-transform: none;
        }

        .container div a {
            color: {{ .TextColor }} !important;
            text-decoration: underline;
        }

        .title {
            margin: 20px;
        }

        .container {
            background: {{ .ContainerBgColor }};
            margin-top: 20px !important;
            margin-bottom: 20px !important;
            margin-left: auto !important;
            margin-right: auto !important;
            border-radius: .28571429rem;
            padding: 10px;
            width: 70%;
        }

        .divider {
            padding: 5px;
        }
    </style>
    <title>Review</title>
</head>

<body style="margin: 0;text-align: center;font-family: Lato,'Helvetica Neue',Arial,Helvetica,sans-serif; color: {{ .TextColor }} !important;">
    <table style="width: 100%; text-align: center; border-collapse: collapse; background-color: {{ .BgColor }};">
        <tbody>
            <tr>
                <td style="width: 100%;">
                    <div class="container">
                        <div class="title">
                            <h1>{{ .HotspotName }}</h1>
                        </div>
                        <div>
                            {{ .HotspotSurveyBodyText }}
                        </div>
                        <h5>
                            {{ .HotspotDetails }}
                        </h5>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>
</body>

</html>
