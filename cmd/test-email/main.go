package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sponsoration/api/internal/service"
)

func main() {
	// Get email from command line
	var testEmail string
	if len(os.Args) > 1 {
		testEmail = os.Args[1]
	} else {
		fmt.Println("⚠️  Warning: No email address provided!")
		fmt.Println("Usage: ENV=production go run cmd/test-email/main.go your-email@example.com")
		fmt.Println("\nProceeding with default test@example.com...\n")
		testEmail = "test@example.com"
	}

	fmt.Println("🧪 Testing Email Service...\n")
	fmt.Printf("📧 Test email will be sent to: %s\n", testEmail)
	fmt.Printf("🌍 Environment: %s\n", getEnv())
	fmt.Printf("📨 From: %s <%s>\n", os.Getenv("SENDGRID_FROM_NAME"), os.Getenv("SENDGRID_FROM_EMAIL"))
	fmt.Println("\n" + repeat("=", 60) + "\n")

	// Create email service
	emailService := service.NewEmailService()

	// Test 1: Verification email
	fmt.Println("1️⃣  Testing Verification Email...")
	err1 := emailService.SendVerificationEmail(testEmail, "TEST123")
	printResult(err1)

	// Wait between emails
	time.Sleep(1 * time.Second)

	// Test 2: Password reset email
	fmt.Println("2️⃣  Testing Password Reset Email...")
	err2 := emailService.SendPasswordResetEmail(testEmail, "RESET456", "Test User")
	printResult(err2)

	// Wait between emails
	time.Sleep(1 * time.Second)

	// Test 3: Welcome email
	fmt.Println("3️⃣  Testing Welcome Email...")
	err3 := emailService.SendWelcomeEmail(testEmail, "Test User")
	printResult(err3)

	// Summary
	fmt.Println(repeat("=", 60))
	allPassed := err1 == nil && err2 == nil && err3 == nil

	if allPassed {
		fmt.Println("\n🎉 All email tests passed!")
		fmt.Printf("\n📬 Check your inbox at: %s\n", testEmail)
		fmt.Println("   (Don't forget to check spam folder)")
	} else {
		fmt.Println("\n⚠️  Some email tests failed!")
		fmt.Println("   Check the error messages above for details.")
		os.Exit(1)
	}
}

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "development"
	}
	return env
}

func printResult(err error) {
	if err != nil {
		log.Printf("   ❌ Failed: %v\n", err)
	} else {
		fmt.Println("   ✅ Success\n")
	}
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
