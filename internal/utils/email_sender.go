package utils

import (
	"net/smtp"
	"os"
	"strconv"
)

func SendEmail(email string, code int) error {
	from := os.Getenv("SMTP_EMAIL_LOGIN")
	password := os.Getenv("SMTP_EMAIL_PASSWORD")

	toList := []string{email}

	host := "smtp.mail.ru"
	port := "587"

	msg := "Привет!\nТвой код для подтверждения регистрации:\n"
	msg += strconv.Itoa(code)

	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		return err
	}

	return nil
}
