package main

import (
	"fmt"
	"net/smtp"
)

//sendmail
func SendMail(mail string, name string, surname string, subject string, cmdNumber string, message string, pro bool) bool {

	to := ""

	//set pro or not mailbox
	if pro == true {
		to = "service.pro.orderndrink@gmail.com"
	} else {
		to = "service.orderndrink@gmail.com"
	}

	//configure sending mailbox
	from := "sav.orderndrink@gmail.com"
	pass := "OrderNDrink2020"

	//set message
	message = name + " " + surname + "\n" + mail + "\n" + "Objet: " + subject + "\n" + "N commande: " + cmdNumber + "\n" + "Message: " + message

	msg := "From: " + mail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("smtp error: %s", err)
		fmt.Println("failed")
		return false
	} else {
		fmt.Println("passed")
		return true
	}
}
