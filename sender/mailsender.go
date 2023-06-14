package sender

import (
	"bytes"
	"fmt"
	"hello/golang_training/hackaton/mongodb"
	"html/template"

	"gopkg.in/gomail.v2"
)

func Sendmail(p []mongodb.Fields, templatePath string) {

	for _, i := range p {
		var body bytes.Buffer
		t, err := template.ParseFiles(templatePath)
		t.Execute(&body, struct {
			Name          string
			ScheduledTime string
			EmailID       string
		}{Name: i.CandidateName, ScheduledTime: i.ScheduledTme, EmailID: i.EmailID})
		if err != nil {
			fmt.Println(err)
			return
		}
		mail_sender(i.EmailID, body)
	}
}
func mail_sender(EmailID string, body bytes.Buffer) {
	m := gomail.NewMessage()
	m.SetHeader("From", "dhvanishah2501@gmail.com")
	m.SetHeader("To", EmailID)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 587, "dhvanishah2501@gmail.com", "cbutisjvrovqbyzq")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
