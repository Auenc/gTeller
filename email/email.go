package email

import (
	"bytes"
	m "github.com/keighl/mandrill"
)

const (
	KEY = "O9QtV8MopAUbY8QY9b8DFQ"
)

type Email struct {
	ID            string
	ReceiverEmail string
	ReceiverName  string
	Template      EmailTemplate
	Sent          bool
}

func (email *Email) Send(data interface{}) error {
	client := m.ClientWithKey(KEY)

	//Executing template
	var temp bytes.Buffer
	t, err := email.Template.Parse()
	if err != nil {
		return err
	}
	t.Execute(&temp, data)

	message := &m.Message{}
	message.AddRecipient(email.ReceiverEmail, email.ReceiverName, "to")
	message.FromEmail = email.Template.SenderEmail
	message.FromName = email.Template.SenderName
	message.Subject = email.Template.Subject
	message.HTML = temp.String()
	message.Text = temp.String() //May want a way to get plan text version of an email

	_, err = client.MessagesSend(message)
	if err != nil {
		return err
	}

	return nil
}
