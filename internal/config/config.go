package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

var (
	port            string
	dbURL           string
	redisHost       string
	redisUsername   string
	redisPassword   string
	cacheUserExpiry int
	tokenExpiry     int
	tokenSecret     string
)

func init() {
	port = getEnv("PORT", "8080")
	dbURL = getEnv("DB_URL", "host=localhost auth=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	redisHost = getEnv("CACHE_REDIS_HOST", "localhost:6379")
	redisUsername = getEnv("CACHE_REDIS_USERNAME", "")
	redisPassword = getEnv("CACHE_REDIS_PASSWORD", "")
	cacheUserExpiry = getEnvInt("CACHE_USER_EXPIRY", 3600)
	tokenExpiry = getEnvInt("TOKEN_EXPIRY", 86400)
	tokenSecret = getEnv("TOKEN_SECRET", "secret")
}

func getEnv(key string, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func getEnvInt(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return def
		}
		return intVal
	}
	return def
}

func Init() error {
	initConfigFuncs := []func() error{
		initCache,
		initDatabase,
	}

	for _, initConfigFunc := range initConfigFuncs {
		err := initConfigFunc()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPort() string {
	return port
}

func GetCacheUserExpiry() int {
	return cacheUserExpiry
}

func GetTokenExpiry() int {
	return tokenExpiry
}

func GetTokenSecret() string {
	return tokenSecret
}

func GetDBURL() string {
	return dbURL
}
