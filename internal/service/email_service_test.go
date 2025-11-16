package service

import (
	"os"
	"strings"
	"testing"
)

func TestNewEmailService(t *testing.T) {
	tests := []struct {
		name              string
		envVars           map[string]string
		expectedFromEmail string
		expectedFromName  string
		expectedIsDev     bool
	}{
		{
			name: "with all environment variables set",
			envVars: map[string]string{
				"SENDGRID_API_KEY":    "test-api-key",
				"SENDGRID_FROM_EMAIL": "test@example.com",
				"SENDGRID_FROM_NAME":  "Test Service",
				"ENV":                 "production",
			},
			expectedFromEmail: "test@example.com",
			expectedFromName:  "Test Service",
			expectedIsDev:     false,
		},
		{
			name:              "with no environment variables (defaults)",
			envVars:           map[string]string{},
			expectedFromEmail: "noreply@yourdomain.com",
			expectedFromName:  "Sponsoration",
			expectedIsDev:     true,
		},
		{
			name: "development environment",
			envVars: map[string]string{
				"ENV": "development",
			},
			expectedFromEmail: "noreply@yourdomain.com",
			expectedFromName:  "Sponsoration",
			expectedIsDev:     true,
		},
		{
			name: "production environment",
			envVars: map[string]string{
				"ENV":              "production",
				"SENDGRID_API_KEY": "test-key",
			},
			expectedFromEmail: "noreply@yourdomain.com",
			expectedFromName:  "Sponsoration",
			expectedIsDev:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			os.Clearenv()

			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Create service
			service := NewEmailService()

			// Verify configuration
			if service.fromEmail != tt.expectedFromEmail {
				t.Errorf("fromEmail = %v, want %v", service.fromEmail, tt.expectedFromEmail)
			}
			if service.fromName != tt.expectedFromName {
				t.Errorf("fromName = %v, want %v", service.fromName, tt.expectedFromName)
			}
			if service.isDev != tt.expectedIsDev {
				t.Errorf("isDev = %v, want %v", service.isDev, tt.expectedIsDev)
			}
		})
	}
}

func TestSendEmail_DevMode(t *testing.T) {
	// Set up development environment
	os.Clearenv()
	os.Setenv("ENV", "development")

	service := NewEmailService()

	tests := []struct {
		name    string
		opts    EmailOptions
		wantErr bool
	}{
		{
			name: "valid email with all fields",
			opts: EmailOptions{
				To:      "test@example.com",
				Subject: "Test Email",
				Text:    "This is a test",
				HTML:    "<p>This is a test</p>",
			},
			wantErr: false,
		},
		{
			name: "valid email with only text",
			opts: EmailOptions{
				To:      "test@example.com",
				Subject: "Test Email",
				Text:    "This is a test",
			},
			wantErr: false,
		},
		{
			name: "valid email with only html",
			opts: EmailOptions{
				To:      "test@example.com",
				Subject: "Test Email",
				HTML:    "<p>This is a test</p>",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.SendEmail(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendVerificationEmail(t *testing.T) {
	// Set up development environment
	os.Clearenv()
	os.Setenv("ENV", "development")

	service := NewEmailService()

	tests := []struct {
		name    string
		email   string
		code    string
		wantErr bool
	}{
		{
			name:    "valid verification email",
			email:   "user@example.com",
			code:    "ABC123",
			wantErr: false,
		},
		{
			name:    "verification email with numeric code",
			email:   "user@example.com",
			code:    "123456",
			wantErr: false,
		},
		{
			name:    "verification email with mixed case code",
			email:   "user@example.com",
			code:    "AbC123",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.SendVerificationEmail(tt.email, tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendVerificationEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendPasswordResetEmail(t *testing.T) {
	// Set up development environment
	os.Clearenv()
	os.Setenv("ENV", "development")

	service := NewEmailService()

	tests := []struct {
		name     string
		email    string
		code     string
		userName []string
		wantErr  bool
	}{
		{
			name:     "password reset with user name",
			email:    "user@example.com",
			code:     "RESET123",
			userName: []string{"John Doe"},
			wantErr:  false,
		},
		{
			name:     "password reset without user name",
			email:    "user@example.com",
			code:     "RESET456",
			userName: nil,
			wantErr:  false,
		},
		{
			name:     "password reset with empty user name",
			email:    "user@example.com",
			code:     "RESET789",
			userName: []string{""},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.SendPasswordResetEmail(tt.email, tt.code, tt.userName...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPasswordResetEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendWelcomeEmail(t *testing.T) {
	// Set up development environment
	os.Clearenv()
	os.Setenv("ENV", "development")
	os.Setenv("APP_URL", "https://example.com")

	service := NewEmailService()

	tests := []struct {
		name    string
		email   string
		userName string
		wantErr bool
	}{
		{
			name:     "welcome email with user name",
			email:    "user@example.com",
			userName: "Jane Smith",
			wantErr:  false,
		},
		{
			name:     "welcome email with single name",
			email:    "user@example.com",
			userName: "John",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.SendWelcomeEmail(tt.email, tt.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendWelcomeEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailTemplateContent(t *testing.T) {
	tests := []struct {
		name           string
		templateFunc   func() string
		expectedParts  []string
		unexpectedParts []string
	}{
		{
			name: "verification email template",
			templateFunc: func() string {
				return getVerificationEmailTemplate("TEST123")
			},
			expectedParts: []string{
				"TEST123",
				"Verify Your Email Address",
				"24 hours",
				"Sponsoration",
				"<!DOCTYPE html>",
			},
			unexpectedParts: []string{
				"Password",
				"Welcome",
			},
		},
		{
			name: "password reset email template",
			templateFunc: func() string {
				return getPasswordResetEmailTemplate("RESET456", "Hi John,")
			},
			expectedParts: []string{
				"RESET456",
				"Hi John,",
				"Reset Your Password",
				"24 hours",
				"Security Tip",
				"🔒",
			},
			unexpectedParts: []string{
				"Verify",
				"Welcome",
			},
		},
		{
			name: "welcome email template",
			templateFunc: func() string {
				return getWelcomeEmailTemplate("Jane Smith", "https://app.example.com")
			},
			expectedParts: []string{
				"Jane Smith",
				"Welcome to Sponsoration",
				"Go to Dashboard",
				"https://app.example.com",
				"🎉",
			},
			unexpectedParts: []string{
				"Verify",
				"Reset",
				"code",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := tt.templateFunc()

			// Check expected parts are present
			for _, part := range tt.expectedParts {
				if !strings.Contains(template, part) {
					t.Errorf("Template missing expected content: %q", part)
				}
			}

			// Check unexpected parts are not present
			for _, part := range tt.unexpectedParts {
				if strings.Contains(template, part) {
					t.Errorf("Template contains unexpected content: %q", part)
				}
			}

			// Verify it's valid HTML
			if !strings.HasPrefix(template, "\n<!DOCTYPE html>") {
				t.Error("Template should start with <!DOCTYPE html>")
			}
			if !strings.Contains(template, "</html>") {
				t.Error("Template should end with </html>")
			}
		})
	}
}

func TestEmailTemplateVariableSubstitution(t *testing.T) {
	t.Run("verification code is properly substituted", func(t *testing.T) {
		code := "XYZ789"
		template := getVerificationEmailTemplate(code)

		// Should appear in the code box
		if !strings.Contains(template, code) {
			t.Errorf("Template should contain code %q", code)
		}

		// Should not contain placeholder
		if strings.Contains(template, "%s") {
			t.Error("Template should not contain unsubstituted placeholders")
		}
	})

	t.Run("password reset greeting is properly substituted", func(t *testing.T) {
		greeting := "Hi Test User,"
		code := "RESET999"
		template := getPasswordResetEmailTemplate(code, greeting)

		if !strings.Contains(template, greeting) {
			t.Errorf("Template should contain greeting %q", greeting)
		}
		if !strings.Contains(template, code) {
			t.Errorf("Template should contain code %q", code)
		}
	})

	t.Run("welcome email personalization", func(t *testing.T) {
		name := "Alice Johnson"
		appURL := "https://test.example.com"
		template := getWelcomeEmailTemplate(name, appURL)

		if !strings.Contains(template, name) {
			t.Errorf("Template should contain name %q", name)
		}
		if !strings.Contains(template, appURL) {
			t.Errorf("Template should contain app URL %q", appURL)
		}
	})
}

func TestEmailTemplateHTML(t *testing.T) {
	tests := []struct {
		name         string
		templateFunc func() string
	}{
		{
			name: "verification email",
			templateFunc: func() string {
				return getVerificationEmailTemplate("TEST")
			},
		},
		{
			name: "password reset email",
			templateFunc: func() string {
				return getPasswordResetEmailTemplate("RESET", "Hello,")
			},
		},
		{
			name: "welcome email",
			templateFunc: func() string {
				return getWelcomeEmailTemplate("User", "http://localhost:8082")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := tt.templateFunc()

			// Basic HTML structure checks
			requiredTags := []string{
				"<!DOCTYPE html>",
				"<html>",
				"</html>",
				"<head>",
				"</head>",
				"<body",
				"</body>",
				"<table",
				"</table>",
			}

			for _, tag := range requiredTags {
				if !strings.Contains(template, tag) {
					t.Errorf("Template missing required HTML tag: %s", tag)
				}
			}

			// Check for responsive meta tag
			if !strings.Contains(template, `<meta name="viewport"`) {
				t.Error("Template should include viewport meta tag for responsiveness")
			}

			// Check for UTF-8 charset
			if !strings.Contains(template, `charset="utf-8"`) {
				t.Error("Template should include UTF-8 charset declaration")
			}
		})
	}
}

// Benchmark tests
func BenchmarkSendVerificationEmail(b *testing.B) {
	os.Clearenv()
	os.Setenv("ENV", "development")
	service := NewEmailService()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.SendVerificationEmail("test@example.com", "CODE123")
	}
}

func BenchmarkGetVerificationEmailTemplate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getVerificationEmailTemplate("TEST123")
	}
}

func BenchmarkSendPasswordResetEmail(b *testing.B) {
	os.Clearenv()
	os.Setenv("ENV", "development")
	service := NewEmailService()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.SendPasswordResetEmail("test@example.com", "RESET123", "User")
	}
}

func BenchmarkSendWelcomeEmail(b *testing.B) {
	os.Clearenv()
	os.Setenv("ENV", "development")
	service := NewEmailService()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.SendWelcomeEmail("test@example.com", "User Name")
	}
}
