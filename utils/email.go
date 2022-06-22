package utils

import (
	"fmt"
	"net/smtp"
	"visitor-management-system/config"
	"visitor-management-system/model"
)

func SendSubscriptionEmail(user *model.User, password string) error {
	to := []string{user.Email}

	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort

	subject := "Welcome to VMS"
	body := fmt.Sprintf("Your credentials for login are given below: \n")
	body += fmt.Sprintf("Username: %s \n", user.Email)
	body += fmt.Sprintf("Password: %s \n", password)

	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	//fmt.Println(message)
	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

	fmt.Println("email processed")
	return err

}

// func SendOfficialCreatedEmail(user *model.User) error {
// 	to := []string{user.Email}

// 	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort
// 	fmt.Println(config.GetConfig().Username)
// 	//build msg
// 	subject := "Welcome to VMS"
// 	body := fmt.Sprintf("dear %s ,\n", user.Name)
// 	//body += fmt.Sprintf("Your %s subscription is activated. \n", user.Subscription_type)
// 	body += fmt.Sprintf("Your credentials for login are given below: \n")
// 	body += fmt.Sprintf("Username: %s \n", user.Email)
// 	body += fmt.Sprintf("Password: %s \n", user.Password)

// 	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
// 	message += fmt.Sprintf("To: %s\r\n", to)
// 	message += fmt.Sprintf("Subject: %s\r\n", subject)
// 	message += fmt.Sprintf("\r\n%s\r\n", body)

// 	//fmt.Println(message)
// 	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
// 	fmt.Println(auth)
// 	// send mail
// 	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

// 	fmt.Println("email processed")
// 	return err

// }
