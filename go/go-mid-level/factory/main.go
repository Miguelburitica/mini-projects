package main

import "fmt"

// SMS, Email
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}


// ====================================> <===========================================
// INotificationFactory
type SMSNotification struct {

}
func (sms SMSNotification) SendNotification() {
	fmt.Println("Sending Notifiaction throught SMS")
}
func (SMSSender SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// ====================================> <===========================================
// ISender
type SMSNotificationSender struct {

}
func (SMSSender SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSSender SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}


// ====================================> <===========================================
// ====================================> <===========================================
// INotificationFactory
type EmailNotification struct {

}
func (sms EmailNotification) SendNotification() {
	fmt.Println("Sending Notifiaction throught Email")
}
func (SMSSender EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// ====================================> <===========================================
// ISender
type EmailNotificationSender struct {

}
func (EmailSender EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailSender EmailNotificationSender) GetSenderChannel() string {
	return "SNS Amazon"
}

// ====================================> <===========================================
func getNotificationFactory (notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}