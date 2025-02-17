package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver = GetEnv("DB_DRIVER", "postgres")
	DBName   = GetEnv("DB_NAME", "local")
	DBHost   = GetEnv("DB_HOST", "localhost")
	DBPort   = GetEnv("DB_PORT", "5432")
	DBUser   = GetEnv("DB_USER", "root")
	DBPass   = GetEnv("DB_PASS", "")
	SSLMode  = GetEnv("SSL_MODE", "disable")

	REDISHost = GetEnv("REDIS_HOST")
	REDISPass = GetEnv("REDIS_PASS")
	REDISPort = GetEnv("REDIS_PORT")

	MONGOHost = GetEnv("MONGO_HOST")
	MONGOPort = GetEnv("MONGO_PORT")
	MONGODB   = GetEnv("MONGO_DB")

	URLCTE         = GetEnv("URLCTE")
	URLOTE         = GetEnv("URLOTE")
	URLBPBATAM     = GetEnv("URLBPBATAM")
	URLBPBATAMPROD = GetEnv("URLBPBATAMPROD")
	URLDITLALA     = GetEnv("URLDITLALA")
	RQID_PELNI     = GetEnv("RQID_PELNI")

	URLBILLER           = GetEnv("URL_BILLER")
	URLILCS             = GetEnv("URLILCS")
	BILLER_MERCHANT_KEY = GetEnv("BILLER_MERCHANT_KEY")
	PELNI_CID           = GetEnv("PELNI_CID")
	ACCOUNT_BATAM       = GetEnv("ACCOUNT_BATAM")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}
}
