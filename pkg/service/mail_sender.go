package service

import (
	"context"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiNGQyN2VjMGE0YmEwMTQxMGRiMTI1OGEzYjQxMTgzYjJjOTg3OGQxOWMxZjBmNWQwZDFjMDlmYTczMjhmN2QzMDQ1Njc0MDg1ZTI4NWNiMTQiLCJpYXQiOjE2MTkzODY2MzUsIm5iZiI6MTYxOTM4NjYzNSwiZXhwIjo0Nzc1MDYwMjM1LCJzdWIiOiI1MTQ0Iiwic2NvcGVzIjpbImVtYWlsX2Z1bGwiLCJkb21haW5zX2Z1bGwiLCJhY3Rpdml0eV9mdWxsIiwiYW5hbHl0aWNzX2Z1bGwiLCJ0b2tlbnNfZnVsbCIsIndlYmhvb2tzX2Z1bGwiLCJ0ZW1wbGF0ZXNfZnVsbCJdfQ.mZ6aZXbtiMnp4g9LHoZgB5LSauOMPnsNoCVwqTPWgBHvcEafz8VTmqQDnyV4bnfAyFrROo_U_iZbEdoCYfKip7ksObe5P24FTNsQCehfsF_IB5ok2T_y8ZKNY9adiNkL1_DUaGtjS2TpQfK33nzXoQ6BQrAH9a0KjMw7t61SC1rWCd5_CvHLGtHyuvr2ssyuyDUC_BnE6mkckmcylymeoKLKNNMutWk9CDFw7_YuTPvdoRwNRZU9n7VMMUDSNLLpSeLkaBsc7EhmXkx4sJsus8484suCn_TFH5iMsz3Q3OHL2TVBH0U1pd8CyVSvn7PmPaDfQ6FsAFG_t55YyjEOP2BG2KSE6REjU-QWoQ95WRUZNUQ5g1rKIAaLBmVMeEHGqfUsrSxo7i72-JYelDyEiiyrn_W5DZCNncRlfSUZqwZP714Kxn7pX9HVsJ2TiR3e44ClvYVrLxwvhCvPFDwUsRfW7P4PyxC70_x7H1NA4URIsuB3zDdGfqTdSz2Q_21zKE5rcXgmrNp9fzTMGUuxJFXt5rf8EKLbRucFSfRzmwsENgvWdOIDLs0JYM94hxXq4TiRL9jSuRx4llUSOjz9KOb36WVKfobevD6yCCHh41PIwr6Mw7lH6K02SOhQCYeZVJwjSFPK9rAd0bRV1xOYXRrWGIDQNFzXO7HIW-JHAu8"

func SendMail(subject string, html string, name string, email string) error {
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject = subject
	html = html

	from := mailersend.From{
		Name:  "RHINODESIGN",
		Email: "hello@app.rhinodesign.ru",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  name,
			Email: email,
		},
	}

	//variables := []mailersend.Variables{
	//	{
	//		Email: email,
	//		Substitutions: []mailersend.Substitution{
	//			{
	//				Var:   "foo",
	//				Value: "bar",
	//			},
	//		},
	//	},
	//}

	//tags := []string{"foo", "bar"}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	//message.SetText(text)
	//message.SetTemplateID("testtemplateid")
	//message.SetSubstitutions(variables)
	//
	//message.SetTags(tags)

	_, err := ms.Send(ctx, message)

	return err
}
