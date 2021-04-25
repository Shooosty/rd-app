package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiOTNhZDM3MDI2ZmNmZjE2ZDNhZGEyMzU4MDMzMTkyZTg1OWQ3MzFjNGVkOGVjMzI2Y2I5MDdhZmYxMTI5M2IyM2MzOTk0MmYyNjU1ZTc0ZTgiLCJpYXQiOjE2MTkzODYxMzIsIm5iZiI6MTYxOTM4NjEzMiwiZXhwIjo0Nzc1MDU5NzMyLCJzdWIiOiI1MTQ0Iiwic2NvcGVzIjpbImVtYWlsX2Z1bGwiLCJkb21haW5zX2Z1bGwiLCJhY3Rpdml0eV9mdWxsIiwiYW5hbHl0aWNzX2Z1bGwiLCJ0b2tlbnNfZnVsbCIsIndlYmhvb2tzX2Z1bGwiLCJ0ZW1wbGF0ZXNfZnVsbCJdfQ.dh9qZ5Yz6xKUYNM3oVEbQuDKfRntPcS4xH55Z1G4EFfbft8W3LgpyUHsVTYe_f39_borL2BZYDFQFks5ip2RoMvWBjSHdnL1HSugY9jOlNjz3cKiMsnVJ1HWPWjocbInlj-G0Jkzc6RdQ8PQi2cKwfalx8S5CCj3xfBiaWOg-QmwV6SBn725PA0eidrjFyUCDkWg89kOc5JhQtBAI4aJtnLz9PZCHblrLme8_7Oxjv_17FeZ_rl1wqZKqNVJerwWUrO4OZ5ZGJSSaZ-nAs2uuLE_eaCIl9geZhyEGWP5AtHK_8wyFMGxSI_TIKFOwEVBwOgxj9igJcFUHsKmzgfHHYHncNCx_SADap2d0zOe44CFWZPaADADE5_rnghLStzwSc_OA4oU-IaAeDT_tXa4YhN79_Gv3KuW5dX4XlfgJ4PECHu_utyyK6asblcUg9CV5MNMTDX-6jOBghZ5u66FpTT2rOIqU0AS-v4Sm97gUfbwYpAd9VhqqjpL8hHDk-dNPFAy34m2el0CDr2pWWuOKbXarUNELKI8LeOsjT7BPDRU4DzkppmqKfXS2kcY2mstRznQlnoVbXsTto5eIWU2_L36Mh3FRmbq_wP9r5oqhMsZN1ZK8kc1nqpKeIbocW-ZuTjliY4J0sgfLNofcGxBKswRgMk1Bn_uFoYAYK4FBzQ"

func SendMail(text string, html string, name string, email string) {
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text = text
	html = html

	from := mailersend.From{
		Name:  "RHINODESIGN",
		Email: "app.rhinodesign.ru",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  name,
			Email: email,
		},
	}

	variables := []mailersend.Variables{
		{
			Email: email,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "foo",
					Value: "bar",
				},
			},
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTemplateID("testtemplateid")
	message.SetSubstitutions(variables)

	message.SetTags(tags)

	res, _ := ms.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
