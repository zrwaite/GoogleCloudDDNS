package mail

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/zrwaite/google-cloud-ddns/models"
)

func SendMessage(toEmail string, toName string, subject string, content string, params *models.Params) (success bool) {
	from := mail.NewEmail("Zac Waite", params.FromEmail)
	to := mail.NewEmail(toName, toEmail)
	plainTextContent := content
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(params.SendgridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func ErrorMessage(content string, params *models.Params) (success bool) {
	subject := "Error on google-cloud-ddns"
	return SendMessage(params.ContactEmail, "Zac Waite", subject, content, params)
}
