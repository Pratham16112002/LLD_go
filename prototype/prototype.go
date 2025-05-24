package prototype

import "fmt"

type EmailTemplate interface {
	Clone() EmailTemplate
	SetContent(content string)
	SetSubject(subject string)
	Send(to string)
}

type WelcomeEmail struct {
	subject string
	content string
}

func NewWelcomeEmail() *WelcomeEmail {
	return &WelcomeEmail{
		subject: "Welcome to TUF+",
		content: "Hi there! I am using WhatsApp",
	}
}

func (wl *WelcomeEmail) Clone() EmailTemplate {
	return &WelcomeEmail{
		subject: wl.subject,
		content: wl.content,
	}
}

func (wl *WelcomeEmail) SetContent(content string) {
	wl.content = content
}

func (wl *WelcomeEmail) SetSubject(subject string) {
	wl.subject = subject
}
func (wl *WelcomeEmail) Send(to string) {
	fmt.Printf("Sending email to %s\nSubject : %s\nContent : %s\n", to, wl.subject, wl.content)
}

type PasswordResetEmail struct {
	content string
	subject string
}

func NewPasswordResetEmail() *PasswordResetEmail {
	return &PasswordResetEmail{
		subject: "Password Reset Request",
		content: "Click on the link below to Reset your password",
	}
}

func (e *PasswordResetEmail) Clone() EmailTemplate {
	return &PasswordResetEmail{
		subject: e.content,
		content: e.subject,
	}
}

func (e *PasswordResetEmail) SetContent(content string) {
	e.content = content
}

func (e *PasswordResetEmail) SetSubject(subject string) {
	e.subject = subject
}
func (e *PasswordResetEmail) Send(to string) {
	fmt.Printf("Sending email to %s\nSubject : %s\nContent : %s\n", to, e.subject, e.content)
}
