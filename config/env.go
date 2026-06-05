package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	JWTSecret        string
	JWTExpire        int
	RefreshJWTExpire int
	jwtExpire        string
	refJWTExpier     string
	ISDev            string
	ISLocal          string
	OtpAppPassword   string
	OtpEmailSender   string
}

func LoadConfig() *Config {
	var err error
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("error on load env file", "error", err)
	}
	conf := Config{
		Port:           getEnv("PORT", "8000"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "ahmed"),
		DBPassword:     getEnv("DB_PASSWORD", "admin"),
		DBName:         getEnv("DB_NAME", "akalni"),
		JWTSecret:      getEnv("JWT_SECRET", "secret"),
		jwtExpire:      getEnv("JWT_EXPIRE", "1"),
		refJWTExpier:   getEnv("REF_JWT_EXPIER", "24"),
		ISDev:          getEnv("IS_DEV", "false"),
		ISLocal:        getEnv("IS_LOCAL", "false"),
		OtpAppPassword: getEnv("GMAIL_APP_PASSWORD", ""),
		OtpEmailSender: getEnv("OTP_EMAIL_SENDER", ""),
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
