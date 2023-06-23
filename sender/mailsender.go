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
			Name     string
			StartTme string
			EmailID  string
		}{Name: i.CandidateName, StartTme: i.ScheduledTme, EmailID: i.EmailID})

		// layout := "02-01-2006 15:04"

		// // Parse the scheduled time string
		// scheduledTime, err := time.Parse(layout, i.ScheduledTme)
		// if err != nil {
		// 	fmt.Println("Error parsing scheduled time:", err)
		// 	return
		// }

		// // Calculate the duration between the current time and the scheduled time
		// duration := scheduledTime.Sub(time.Now().Add(-15 * time.Minute))

		// // Sleep for the duration
		// time.Sleep(duration)

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
	m.SetHeader("Subject", "Interview Slot")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 587, "dhvanishah2501@gmail.com", "cbutisjvrovqbyzq")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
