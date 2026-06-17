package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                     string
	DBHost                   string
	DBPort                   string
	DBUser                   string
	DBPassword               string
	DBName                   string
	JWTSecret                string
	JWTExpire                int
	RefreshJWTExpire         int
	ISDev                    string
	ISLocal                  string
	OtpAppPassword           string
	OtpEmailSender           string
	OTPNumberOfRetries       int
	OTPNumberOfResend        int
	EmailReverificationAfter int
	OtpLenght                int
	OTPSalt                  int
	OtpExpire                int
	// temporary
	jwtExpire                string
	refJWTExpier             string
	otpNumberOfRetries       string
	otpNumberOfResend        string
	emailReverificationAfter string
	otpLenght                string
	otpSalt                  string
	otpExpire                string
}

func LoadConfig() *Config {
	var err error
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("error on load env file", "error", err)
	}
	conf := Config{
		Port:                     getEnv("PORT", "8000"),
		DBHost:                   getEnv("DB_HOST", "localhost"),
		DBPort:                   getEnv("DB_PORT", "5432"),
		DBUser:                   getEnv("DB_USER", "ahmed"),
		DBPassword:               getEnv("DB_PASSWORD", "admin"),
		DBName:                   getEnv("DB_NAME", "akalni"),
		JWTSecret:                getEnv("JWT_SECRET", "secret"),
		jwtExpire:                getEnv("JWT_EXPIRE", "1"),
		refJWTExpier:             getEnv("REF_JWT_EXPIER", "24"),
		ISDev:                    getEnv("IS_DEV", "false"),
		ISLocal:                  getEnv("IS_LOCAL", "false"),
		OtpAppPassword:           getEnv("GMAIL_APP_PASSWORD", ""),
		OtpEmailSender:           getEnv("OTP_EMAIL_SENDER", ""),
		otpNumberOfRetries:       getEnv("OTP_NUMBER_OF_RETRIES", "3"),
		otpNumberOfResend:        getEnv("OTP_NUMBER_OF_RESEND", "3"),
		emailReverificationAfter: getEnv("EMAIL_REVERIFICATION_AFTER", "30"),
	}
	if conf.jwtExpire != "" {
		conf.JWTExpire, err = strconv.Atoi(conf.jwtExpire)
		if err != nil {
			log.Printf("error on convert token expire: %v", err)
			conf.JWTExpire = 1
		}
	}
	if conf.refJWTExpier != "" {
		conf.RefreshJWTExpire, err = strconv.Atoi(conf.refJWTExpier)
		if err != nil {
			log.Printf("error on convert token expire: %v", err)
			conf.RefreshJWTExpire = 24
		}
	}
	if conf.otpNumberOfRetries != "" {
		conf.OTPNumberOfRetries, err = strconv.Atoi(conf.otpNumberOfRetries)
		if err != nil {
			log.Printf("error on convert otp number of retries: %v", err)
			conf.OTPNumberOfRetries = 3
		}
	}
	if conf.otpNumberOfResend != "" {
		conf.OTPNumberOfResend, err = strconv.Atoi(conf.otpNumberOfResend)
		if err != nil {
			log.Printf("error on convert otp number of resend: %v", err)
			conf.OTPNumberOfResend = 3
		}
	}
	if conf.emailReverificationAfter != "" {
		conf.EmailReverificationAfter, err = strconv.Atoi(conf.emailReverificationAfter)
		if err != nil {
			log.Printf("error on convert email reverification after: %v", err)
			conf.EmailReverificationAfter = 30
		}
	}
	if conf.otpLenght != "" {
		conf.OtpLenght, err = strconv.Atoi(conf.otpLenght)
		if err != nil {
			log.Printf("error on convert otp lenght: %v", err)
			conf.OtpLenght = 6
		}
	}
	if conf.otpSalt != "" {
		conf.OTPSalt, err = strconv.Atoi(conf.otpSalt)
		if err != nil {
			log.Printf("error on convert otp salt: %v", err)
			conf.OTPSalt = 26
		}
	}
	if conf.otpExpire != "" {
		conf.OtpExpire, err = strconv.Atoi(conf.otpExpire)
		if err != nil {
			log.Printf("error on convert otp expire: %v", err)
			conf.OtpExpire = 5
		}
	}

	return &conf
}
func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	// log.Printf(" the key :%s,and value :%s\n", key, val)
	if len(val) == 0 {
		log.Printf("warning the value of this key: %s will set to default %v", key, defaultValue)
		return defaultValue
	}
	return val
}
