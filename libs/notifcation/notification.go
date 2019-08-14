package notifcation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	mailgun "github.com/mailgun/mailgun-go"
)

func mailGun(sender, subject, body, recipient string) {

	yourDomain := viper.GetString("mailgun_domain")

	privateAPIKey := viper.GetString("mailgun_privateKey")

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	message.SetHtml(body)

	// Send the message	with a 10 second timeout

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

//SendSms takes in sender, subject, body, recipient and sends an smtp mail to the described destination
func SendSms(sender, subject, body, recipient string) {
	mailGun(sender, subject, body, recipient)
}
