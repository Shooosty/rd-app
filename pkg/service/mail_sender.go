package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpI" +
	"joiYzhkYzA5OWQzN2VmNTQxMTQ5NmE2YTI1Y2QxNWE0OGRiMjMwOTRiMmJlZTQ2ZmI5M2E2" +
	"ODllYjhkZmZjMjQ0MTBiNTRjZDE5NjQzYWNkZDIiLCJpYXQiOjE2MTkyNzQ1MDAsIm5iZiI6" +
	"MTYxOTI3NDUwMCwiZXhwIjo0Nzc0OTQ4MTAwLCJzdWIiOiI1MTQ0Iiwic2NvcGVzIjpbImVtYWlsX" +
	"2Z1bGwiLCJkb21haW5zX2Z1bGwiLCJhY3Rpdml0eV9mdWxsIiwiYW5hbHl0aWNzX2Z1bGwiLCJ0b2t" +
	"lbnNfZnVsbCIsIndlYmhvb2tzX2Z1bGwiLCJ0ZW1wbGF0ZXNfZnVsbCJdfQ.enD94v_KUGa3yQfFHXuX" +
	"pOgiolks67NoT4f3MtneBzu6jr72_VRW8-FtJEAzOKbBLqsp8ERfutnKC_Pt4YK6L2KHTN2rAErz4nZ1I" +
	"RUEAB4k1stbwFW2eIRDklinl7L2-L6AaK67VKzGllCly0so2SV9wIq93IeDnkbjC5c9mKFJ3sGEo5wOLtajJ1" +
	"O6FV2sdDzIbGM_3Z7IsaLzNZGR4kxJwbge4KzSn_C4iMdtJPLALxxbJQIh8TjgaA17iLQ3uh9QaY8umaoMdGip" +
	"3CNffB_YSuaDLx9qzF4aJqBfA9RTqTLXa2MIPOTY9DulN-qlfxRlP7Aa6TMVJ2j8PiRCCZGn6wrMATOluzFFRTbQ" +
	"BRdMI3IPbpxkdt9nJ4ZEYM6K8K7IttRXcNeruQ-LVVpL8FvrvHSkmvJzYsCWXubgW6PnLVAJIBFzPDsag1w9tIFmB-F" +
	"_y2llTJrFsbOFGJpe56F0KQjsQ9gub9o59G8s0-sTA80hKou7nfE1EUuI1_5_OrG3w0w0Q_GM-oKNZbdHakXNLrpL-g_t" +
	"9kubmu2a89jOq-vXj3aQ998-LJVlrPGdpSZArt9q88X8vpU2Jp6_NzOkQpGiw9tYaPvfKiWmn55QTiBNQq1ZbQI3rMmtahB5H" +
	"X6AUGvII1hyh08IeeCD1-_-dglG0F6pmpp_eKwqNiA"

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
			Email: "your@client.com",
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
