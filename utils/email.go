package utils

import (
	"fmt"
	"net/smtp"
	"visitor-management-system/config"
	//"visitor-management-system/model"
)

func SendEmail(email string, password string, subdomain string) error {
	to := []string{email}

	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort

	subject := "Welcome to VMS"
	body := fmt.Sprintf("Your credentials for login are given below: \n")
	body += fmt.Sprintf("Username: %s \n", email)
	body += fmt.Sprintf("Password: %s\n", password)
	if subdomain != "" {
		body += fmt.Sprintf("Company SubDomain: %s\n", subdomain)
	}

	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

	fmt.Println("email processed")
	return err

}
