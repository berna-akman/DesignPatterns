package bridge

import "fmt"

type MessageSender interface {
	SendMessage(message, recipient string) error
}

type MessagingPlatform interface {
	SetMessageSender(sender MessageSender)
	SendNotification(message, recipient string) error
}

type SMS struct {
	sender MessageSender
}

func (s *SMS) SetMessageSender(sender MessageSender) {
	s.sender = sender
}

func (s *SMS) SendNotification(message, recipient string) error {
	fmt.Printf("Sending SMS notification to %s: %s\n", recipient, message)
	return s.sender.SendMessage(message, recipient)
}

type Email struct {
	sender MessageSender
}

func (e *Email) SetMessageSender(sender MessageSender) {
	e.sender = sender
}

func (e *Email) SendNotification(message, recipient string) error {
	fmt.Printf("Sending Email notification to %s: %s\n", recipient, message)
	return e.sender.SendMessage(message, recipient)
}

type MessageSenderImpl struct{}

func (m *MessageSenderImpl) SendMessage(message, recipient string) error {
	fmt.Printf("Sending message to %s: %s\n", recipient, message)
	return nil
}

type Notification struct {
	message   string
	recipient string
	platform  MessagingPlatform
}

func (n *Notification) SetPlatform(platform MessagingPlatform) {
	n.platform = platform
}

func (n *Notification) Send() error {
	return n.platform.SendNotification(n.message, n.recipient)
}

func Send() {
	messageSender := &MessageSenderImpl{}

	sms := &SMS{}
	sms.SetMessageSender(messageSender)

	email := &Email{}
	email.SetMessageSender(messageSender)

	normalSMSNotification := &Notification{
		message:   "Hello from SMS",
		recipient: "Berna",
		platform:  sms,
	}

	normalEmailNotification := &Notification{
		message:   "Hello from Email",
		recipient: "Akman",
		platform:  email,
	}

	urgentSMSNotification := &Notification{
		message:   "Urgent message from SMS",
		recipient: "Berna 2",
		platform:  sms,
	}

	urgentEmailNotification := &Notification{
		message:   "Urgent message from Email",
		recipient: "Akman 2",
		platform:  email,
	}

	normalSMSNotification.Send()
	normalEmailNotification.Send()
	urgentSMSNotification.Send()
	urgentEmailNotification.Send()
}
