package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/AhmedZeyad/Akalni/config"
	"gopkg.in/gomail.v2"
)

// gen otp
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// send otp
func SendOTP(conf *config.Config, To, otpCode string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", conf.OtpEmailSender)
	msg.SetHeader("To", To)
	msg.SetHeader("Subject", fmt.Sprintf("Akalni OTP %d", time.Now().Unix()))

	msg.Embed("assets/logo.png")
	emailTemplet, err := os.ReadFile("assets/email.html")
	if err != nil {
		slog.ErrorContext(context.TODO(), "failed to read email template", "Error", err)
		return err
	}

	msg.SetBody("text/html", fmt.Sprintf(string(emailTemplet), otpCode))

	Dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		conf.OtpEmailSender,
		conf.OtpAppPassword,
	)
	err = Dialer.DialAndSend(msg)
	if err != nil {
		slog.ErrorContext(context.TODO(), "failed to send otp email", "Error", err)
		return err
	}
	return nil
}
func updateUserOTP(otpCode string) error {

	return nil
}

func OTPHandler(conf *config.Config, to string) error {

	// gen otp
	otpCode := generateOTP()
	// save on db
	err := updateUserOTP(otpCode)
	if err != nil {
		slog.ErrorContext(context.TODO(), "failed to update user otp ", "Error", err)
		return err
	}
	// send email
	err = SendOTP(conf, to, otpCode)
	if err != nil {
		slog.ErrorContext(context.TODO(), "failed to send otp to user", "Error", err)
		return err
	}
	return nil
}
