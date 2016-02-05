package email

import (
	"html/template"
)

type EmailTemplate struct {
	ID          string
	Name        string
	Subject     string
	Content     string
	SenderEmail string
	SenderName  string
}

func (temp *EmailTemplate) Parse() (*template.Template, error) {
	return template.New(temp.ID).Parse(temp.Content)
}

func (temp *EmailTemplate) Prepare(id string, toEmail string, toName string) (Email, error) {
	var email Email

	email = Email{id, toEmail, toName, *temp, false}

	return email, nil
}
