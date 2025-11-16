package service

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailService handles sending emails via SendGrid
type EmailService struct {
	apiKey    string
	fromEmail string
	fromName  string
	isDev     bool
}

// EmailOptions contains email parameters
type EmailOptions struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

// NewEmailService creates a new email service instance
func NewEmailService() *EmailService {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	fromEmail := os.Getenv("SENDGRID_FROM_EMAIL")
	fromName := os.Getenv("SENDGRID_FROM_NAME")
	env := os.Getenv("ENV")

	if fromEmail == "" {
		fromEmail = "noreply@yourdomain.com"
	}
	if fromName == "" {
		fromName = "Sponsoration"
	}

	isDev := env == "development" || env == ""

	if apiKey == "" && !isDev {
		log.Println("⚠️  SENDGRID_API_KEY not set in production!")
	}

	return &EmailService{
		apiKey:    apiKey,
		fromEmail: fromEmail,
		fromName:  fromName,
		isDev:     isDev,
	}
}

// SendEmail sends an email via SendGrid
func (s *EmailService) SendEmail(opts EmailOptions) error {
	// Development mode: log instead of sending
	if s.isDev {
		log.Println("📧 Email (DEV MODE - Not actually sent):")
		log.Printf("From: %s <%s>", s.fromName, s.fromEmail)
		log.Printf("To: %s", opts.To)
		log.Printf("Subject: %s", opts.Subject)
		if len(opts.HTML) > 200 {
			log.Printf("Content: %s...", opts.HTML[:200])
		} else {
			log.Printf("Content: %s", opts.Text)
		}
		return nil
	}

	// Production mode: send via SendGrid
	from := mail.NewEmail(s.fromName, s.fromEmail)
	to := mail.NewEmail("", opts.To)

	message := mail.NewSingleEmail(from, opts.Subject, to, opts.Text, opts.HTML)
	client := sendgrid.NewSendClient(s.apiKey)

	response, err := client.Send(message)
	if err != nil {
		log.Printf("❌ Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		log.Printf("❌ SendGrid error: %d - %s", response.StatusCode, response.Body)
		return fmt.Errorf("sendgrid error: %d - %s", response.StatusCode, response.Body)
	}

	log.Printf("✅ Email sent successfully to %s", opts.To)
	return nil
}

// SendVerificationEmail sends an email verification code
func (s *EmailService) SendVerificationEmail(email, code string) error {
	return s.SendEmail(EmailOptions{
		To:      email,
		Subject: "Verify Your Email Address",
		Text:    fmt.Sprintf("Your verification code is: %s", code),
		HTML:    getVerificationEmailTemplate(code),
	})
}

// SendPasswordResetEmail sends a password reset code
func (s *EmailService) SendPasswordResetEmail(email, code string, userName ...string) error {
	greeting := "Hello,"
	if len(userName) > 0 && userName[0] != "" {
		greeting = fmt.Sprintf("Hi %s,", userName[0])
	}

	return s.SendEmail(EmailOptions{
		To:      email,
		Subject: "Reset Your Password",
		Text:    fmt.Sprintf("Your password reset code is: %s", code),
		HTML:    getPasswordResetEmailTemplate(code, greeting),
	})
}

// SendWelcomeEmail sends a welcome email to a new user
func (s *EmailService) SendWelcomeEmail(email, name string) error {
	appURL := os.Getenv("APP_URL")
	if appURL == "" {
		appURL = "http://localhost:8082"
	}

	return s.SendEmail(EmailOptions{
		To:      email,
		Subject: "Welcome to Sponsoration!",
		Text:    fmt.Sprintf("Welcome %s! Thank you for joining Sponsoration.", name),
		HTML:    getWelcomeEmailTemplate(name, appURL),
	})
}
