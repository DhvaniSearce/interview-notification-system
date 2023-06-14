package main

import (
	"bytes"
	"fmt"
	"hello/golang_training/hackaton/mongodb"
	"text/template"

	"gopkg.in/gomail.v2"
)

func main() {

	p := mongodb.Run()
	var body bytes.Buffer
	t, err := template.ParseFiles("./test.html")
	t.Execute(&body, struct {
		Name          string
		ScheduledTime string
	}{Name: p[0].CandidateName, ScheduledTime: p[0].ScheduledTme})
	if err != nil {
		fmt.Println(err)
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "dhvanishah2501@gmail.com")
	m.SetHeader("To", p[0].EmailID)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "dhvanishah2501@gmail.com", "cbutisjvrovqbyzq")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("The email has been sent.")

}
