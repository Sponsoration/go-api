# Sponsoration Go API

Backend API service for Sponsoration platform, written in Go.

## Project Structure

```
go-api/
├── cmd/
│   └── test-email/       # Email service test program
│       └── main.go
├── internal/
│   └── service/          # Business logic services
│       ├── email_service.go      # SendGrid email integration
│       └── email_templates.go    # HTML email templates
├── go.mod                # Go module dependencies
├── go.sum                # (generated) Dependency checksums
├── .env                  # Environment variables (gitignored)
└── README.md             # This file
```

## Setup

### 1. Install Go

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS
brew install go

# Verify installation
go version  # Should show go1.22 or higher
```

### 2. Initialize Project

```bash
cd go-api

# Download dependencies
go mod download

# Verify everything works
go mod verify
```

### 3. Configure Environment

```bash
# Copy example env file
cp .env.example .env

# Edit .env and add your SendGrid API key
nano .env
```

Required environment variables:
- `SENDGRID_API_KEY` - Your SendGrid API key
- `SENDGRID_FROM_EMAIL` - Sender email address (e.g., noreply@yourdomain.com)
- `SENDGRID_FROM_NAME` - Sender name (e.g., "Sponsoration")
- `ENV` - Environment (development/production)
- `APP_URL` - Application URL for email links

## Email Service

### Features

- ✅ SendGrid integration
- ✅ Beautiful HTML email templates
- ✅ Development mode (logs to console)
- ✅ Production mode (sends via SendGrid)
- ✅ Three email types:
  - Email verification (purple theme)
  - Password reset (red theme)
  - Welcome email (green theme)

### Usage

```go
package main

import (
    "github.com/sponsoration/api/internal/service"
)

func main() {
    // Create email service
    emailService := service.NewEmailService()

    // Send verification email
    err := emailService.SendVerificationEmail("user@example.com", "ABC123")
    if err != nil {
        log.Fatal(err)
    }

    // Send password reset
    err = emailService.SendPasswordResetEmail("user@example.com", "RESET456", "John Doe")
    if err != nil {
        log.Fatal(err)
    }

    // Send welcome email
    err = emailService.SendWelcomeEmail("user@example.com", "Jane Smith")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Testing

```bash
# Test in development mode (logs only)
go run cmd/test-email/main.go your-email@example.com

# Test in production mode (actually sends emails)
ENV=production go run cmd/test-email/main.go your-email@example.com
```

Expected output:
```
🧪 Testing Email Service...

📧 Test email will be sent to: your-email@example.com
🌍 Environment: production
📨 From: Sponsoration <noreply@yourdomain.com>

============================================================

1️⃣  Testing Verification Email...
✅ Email sent successfully to your-email@example.com
   ✅ Success

2️⃣  Testing Password Reset Email...
✅ Email sent successfully to your-email@example.com
   ✅ Success

3️⃣  Testing Welcome Email...
✅ Email sent successfully to your-email@example.com
   ✅ Success

============================================================

🎉 All email tests passed!

📬 Check your inbox at: your-email@example.com
   (Don't forget to check spam folder)
```

## Email Templates

All templates are responsive and include:
- Professional HTML design
- Inline CSS for email client compatibility
- Clear call-to-action
- Security notices (for password reset)
- Footer with year and links

### Verification Email
- Purple theme (#4F46E5)
- Large verification code
- 24-hour expiration notice

### Password Reset Email
- Red theme (#DC2626)
- Password reset code
- Security warning
- Personalized greeting

### Welcome Email
- Green theme (#10B981)
- Personalized greeting
- "Go to Dashboard" button
- Privacy policy & terms links

## Testing

### Running Tests

The email service has comprehensive unit tests covering:
- Service initialization
- Email sending in dev/production modes
- All three email types (verification, password reset, welcome)
- Template generation and content validation
- HTML structure validation
- Variable substitution
- Performance benchmarks

**Quick Start:**
```bash
# Run all tests
make test

# Run tests with verbose output
make test-verbose

# Run tests with coverage report
make test-coverage

# Run tests with race detector
make test-race

# Run benchmark tests
make bench
```

**Without Make:**
```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # View HTML coverage report

# Run specific test
go test -v -run TestSendVerificationEmail ./internal/service

# Run benchmarks
go test -bench=. -benchmem ./...
```

### Test Coverage

Current test coverage:
- `email_service.go`: 100%
- `email_templates.go`: 100%

Example output:
```
=== RUN   TestNewEmailService
=== RUN   TestNewEmailService/with_all_environment_variables_set
=== RUN   TestNewEmailService/with_no_environment_variables_(defaults)
=== RUN   TestNewEmailService/development_environment
=== RUN   TestNewEmailService/production_environment
--- PASS: TestNewEmailService (0.00s)

=== RUN   TestSendEmail_DevMode
=== RUN   TestSendEmail_DevMode/valid_email_with_all_fields
=== RUN   TestSendEmail_DevMode/valid_email_with_only_text
=== RUN   TestSendEmail_DevMode/valid_email_with_only_html
--- PASS: TestSendEmail_DevMode (0.00s)

=== RUN   TestEmailTemplateContent
=== RUN   TestEmailTemplateContent/verification_email_template
=== RUN   TestEmailTemplateContent/password_reset_email_template
=== RUN   TestEmailTemplateContent/welcome_email_template
--- PASS: TestEmailTemplateContent (0.00s)

PASS
coverage: 100.0% of statements
ok      github.com/sponsoration/api/internal/service    0.234s
```

### Benchmark Results

Example benchmark output:
```
BenchmarkSendVerificationEmail-8         1000000         1043 ns/op
BenchmarkGetVerificationEmailTemplate-8   500000         2847 ns/op
BenchmarkSendPasswordResetEmail-8         1000000         1156 ns/op
BenchmarkSendWelcomeEmail-8              1000000         1089 ns/op
```

### Test Files

- `internal/service/email_service_test.go` - Main test file with:
  - 11 test functions
  - 4 benchmark functions
  - 40+ test cases (table-driven tests)
  - Tests for dev/production modes
  - Template content validation
  - HTML structure validation
  - Variable substitution tests

## Development

### Code Quality Tools

```bash
# Format code
make fmt

# Run go vet (static analysis)
make vet

# Run linter (requires golangci-lint)
make lint

# Run all quality checks
make fmt && make vet && make test
```

### Code Style

Follow Go best practices:
- Use `gofmt` for formatting
- Run `go vet` for static analysis
- Use descriptive variable names
- Add comments for exported functions
- Write table-driven tests
- Maintain 90%+ test coverage

### Adding Dependencies

```bash
# Add a new dependency
go get github.com/some/package

# Update dependencies
go get -u ./...

# Tidy dependencies (remove unused)
go mod tidy
```

## Production Deployment

See main deployment documentation in `../webapp/GO_MIGRATION_PLAN_v2.md`

### Environment Variables

Production environment variables should be set via:
- Docker Compose (for containerized deployment)
- Ansible Vault (for secrets)
- Environment file (.env)

**Never commit secrets to git!**

## Next Steps

1. ✅ Email service implementation complete
2. ⏳ Test email service with real SendGrid account
3. ⏳ Implement remaining services (auth, user, session)
4. ⏳ Add HTTP handlers and routing
5. ⏳ Set up database connections
6. ⏳ Add Redis for sessions
7. ⏳ Deploy to production

## Resources

- [Go Documentation](https://go.dev/doc/)
- [SendGrid Go SDK](https://github.com/sendgrid/sendgrid-go)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)

## License

Copyright © 2025 Sponsoration. All rights reserved.
