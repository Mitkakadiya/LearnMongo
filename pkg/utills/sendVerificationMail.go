package utills

import (
	"bytes"
	"html/template"
	"net/smtp"
)

func SendVerificationEmail(toEmail string, htmlBody string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := "mitkakadiya08@gmail.com"
	smtpPass := "ccozmmruggcylrnx"

	subject := "Verify your email address"

	// Compose the email headers and body
	msg := "From: " + smtpUser + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		htmlBody

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, []byte(msg))
}

func RenderVerificationEmail(verificationLink string) (string, error) {
	// Load the template file
	tmpl, err := template.ParseFiles("C:/Users/CrawlApps Meet/Desktop/Mit kakadiya/Go/mongo_db/pkg/utills/email_verification.html")
	if err != nil {
		return "", err
	}

	// Prepare data to inject
	data := struct {
		VerificationLink string
	}{
		VerificationLink: verificationLink,
	}

	// Render template into a buffer
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
