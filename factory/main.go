// tneemos un software en envia notificaciones
// sms, push notifications, Email
// crear un programa que pueda enviar cualquier tipo de notificaciones polimorficamente
package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//SMS

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("enviando notificación SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "twilio"
}

// Email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("enviando notificación por email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return ""
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
