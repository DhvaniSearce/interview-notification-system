package main

import (
	"fmt"
	"hello/golang_training/hackaton/mongodb"
	"hello/golang_training/interview-notification-system/sender"
)

func main() {

	p := mongodb.Run()
	sender.Sendmail(p, "./test.html")

	fmt.Println("The email has been sent.")
}
