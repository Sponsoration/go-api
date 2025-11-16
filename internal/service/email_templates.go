package service

import (
	"fmt"
	"time"
)

// getVerificationEmailTemplate returns the HTML template for email verification
func getVerificationEmailTemplate(code string) string {
	year := time.Now().Year()
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verify Your Email</title>
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color: #f4f4f4; padding: 20px;">
    <tr>
      <td align="center">
        <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 8px; overflow: hidden;">
          <!-- Header -->
          <tr>
            <td style="background-color: #4F46E5; padding: 30px 40px; text-align: center;">
              <h1 style="margin: 0; color: #ffffff; font-size: 28px;">Sponsoration</h1>
            </td>
          </tr>

          <!-- Content -->
          <tr>
            <td style="padding: 40px;">
              <h2 style="margin: 0 0 20px 0; color: #1F2937; font-size: 24px;">Verify Your Email Address</h2>
              <p style="margin: 0 0 20px 0; color: #4B5563; font-size: 16px; line-height: 1.5;">
                Thank you for registering! Please use the following code to verify your email address:
              </p>

              <!-- Code Box -->
              <div style="background-color: #F3F4F6; border-radius: 8px; padding: 30px; text-align: center; margin: 30px 0;">
                <div style="font-size: 32px; font-weight: bold; letter-spacing: 8px; color: #4F46E5; font-family: 'Courier New', monospace;">
                  %s
                </div>
              </div>

              <p style="margin: 20px 0 0 0; color: #6B7280; font-size: 14px; line-height: 1.5;">
                This code will expire in <strong>24 hours</strong>.
              </p>
              <p style="margin: 10px 0 0 0; color: #6B7280; font-size: 14px; line-height: 1.5;">
                If you didn't request this verification, please ignore this email.
              </p>
            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="background-color: #F9FAFB; padding: 30px 40px; text-align: center; border-top: 1px solid #E5E7EB;">
              <p style="margin: 0; color: #9CA3AF; font-size: 12px;">
                © %d Sponsoration. All rights reserved.
              </p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
    `, code, year)
}

// getPasswordResetEmailTemplate returns the HTML template for password reset
func getPasswordResetEmailTemplate(code, greeting string) string {
	year := time.Now().Year()
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Reset Your Password</title>
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color: #f4f4f4; padding: 20px;">
    <tr>
      <td align="center">
        <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 8px; overflow: hidden;">
          <!-- Header -->
          <tr>
            <td style="background-color: #DC2626; padding: 30px 40px; text-align: center;">
              <h1 style="margin: 0; color: #ffffff; font-size: 28px;">🔒 Password Reset</h1>
            </td>
          </tr>

          <!-- Content -->
          <tr>
            <td style="padding: 40px;">
              <h2 style="margin: 0 0 20px 0; color: #1F2937; font-size: 24px;">Reset Your Password</h2>
              <p style="margin: 0 0 20px 0; color: #4B5563; font-size: 16px; line-height: 1.5;">
                %s
              </p>
              <p style="margin: 0 0 20px 0; color: #4B5563; font-size: 16px; line-height: 1.5;">
                You requested to reset your password. Please use the following code:
              </p>

              <!-- Code Box -->
              <div style="background-color: #FEF2F2; border: 2px solid #FCA5A5; border-radius: 8px; padding: 30px; text-align: center; margin: 30px 0;">
                <div style="font-size: 32px; font-weight: bold; letter-spacing: 8px; color: #DC2626; font-family: 'Courier New', monospace;">
                  %s
                </div>
              </div>

              <p style="margin: 20px 0 0 0; color: #6B7280; font-size: 14px; line-height: 1.5;">
                This code will expire in <strong>24 hours</strong>.
              </p>
              <p style="margin: 10px 0 0 0; color: #6B7280; font-size: 14px; line-height: 1.5;">
                If you didn't request a password reset, please ignore this email and your password will remain unchanged.
              </p>

              <!-- Security Notice -->
              <div style="background-color: #FFFBEB; border-left: 4px solid #F59E0B; padding: 15px; margin-top: 30px;">
                <p style="margin: 0; color: #92400E; font-size: 13px; line-height: 1.5;">
                  <strong>Security Tip:</strong> Never share your password reset code with anyone. Sponsoration staff will never ask for this code.
                </p>
              </div>
            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="background-color: #F9FAFB; padding: 30px 40px; text-align: center; border-top: 1px solid #E5E7EB;">
              <p style="margin: 0; color: #9CA3AF; font-size: 12px;">
                © %d Sponsoration. All rights reserved.
              </p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
    `, greeting, code, year)
}

// getWelcomeEmailTemplate returns the HTML template for welcome email
func getWelcomeEmailTemplate(name, appURL string) string {
	year := time.Now().Year()
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Welcome to Sponsoration</title>
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color: #f4f4f4; padding: 20px;">
    <tr>
      <td align="center">
        <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 8px; overflow: hidden;">
          <!-- Header -->
          <tr>
            <td style="background-color: #10B981; padding: 30px 40px; text-align: center;">
              <h1 style="margin: 0; color: #ffffff; font-size: 28px;">🎉 Welcome to Sponsoration!</h1>
            </td>
          </tr>

          <!-- Content -->
          <tr>
            <td style="padding: 40px;">
              <h2 style="margin: 0 0 20px 0; color: #1F2937; font-size: 24px;">Hi %s,</h2>
              <p style="margin: 0 0 20px 0; color: #4B5563; font-size: 16px; line-height: 1.5;">
                Thank you for joining our community! We're excited to have you on board.
              </p>
              <p style="margin: 0 0 30px 0; color: #4B5563; font-size: 16px; line-height: 1.5;">
                Get started by completing your profile and exploring the platform.
              </p>

              <!-- CTA Button -->
              <div style="text-align: center; margin: 30px 0;">
                <a href="%s"
                   style="display: inline-block; background-color: #10B981; color: #ffffff; text-decoration: none; padding: 15px 30px; border-radius: 6px; font-weight: bold; font-size: 16px;">
                  Go to Dashboard
                </a>
              </div>

              <p style="margin: 30px 0 0 0; color: #6B7280; font-size: 14px; line-height: 1.5;">
                Best regards,<br>
                <strong>The Sponsoration Team</strong>
              </p>
            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="background-color: #F9FAFB; padding: 30px 40px; text-align: center; border-top: 1px solid #E5E7EB;">
              <p style="margin: 0 0 10px 0; color: #9CA3AF; font-size: 12px;">
                © %d Sponsoration. All rights reserved.
              </p>
              <p style="margin: 0; color: #9CA3AF; font-size: 12px;">
                <a href="%s/privacy/policy" style="color: #6B7280; text-decoration: none;">Privacy Policy</a> •
                <a href="%s/privacy/terms" style="color: #6B7280; text-decoration: none;">Terms of Service</a>
              </p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
    `, name, appURL, year, appURL, appURL)
}
