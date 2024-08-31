package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

var (
	port                         string
	timeout                      int
	dbURL                        string
	redisHost                    string
	redisUsername                string
	redisPassword                string
	tokenExpiry                  int
	tokenSecret                  string
	cacheUserExpiry              int
	cacheTableExpiry             int
	cacheMerchantExpiry          int
	cacheItemExpiry              int
	cacheItemCategoryExpiry      int
	cacheItemVariantExpiry       int
	cacheItemVariantOptionExpiry int
)

func init() {
	port = getEnv("PORT", "8080")
	timeout = getEnvInt("TIMEOUT", 5)
	dbURL = getEnv("DB_URL", "host=localhost auth=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	redisHost = getEnv("CACHE_REDIS_HOST", "localhost:6379")
	redisUsername = getEnv("CACHE_REDIS_USERNAME", "")
	redisPassword = getEnv("CACHE_REDIS_PASSWORD", "")
	tokenExpiry = getEnvInt("TOKEN_EXPIRY", 86400)
	tokenSecret = getEnv("TOKEN_SECRET", "secret")
	cacheUserExpiry = getEnvInt("CACHE_USER_EXPIRY", 3600)
	cacheTableExpiry = getEnvInt("CACHE_TABLE_EXPIRY", 3600)
	cacheMerchantExpiry = getEnvInt("CACHE_MERCHANT_EXPIRY", 3600)
	cacheItemExpiry = getEnvInt("CACHE_ITEM_EXPIRY", 3600)
	cacheItemCategoryExpiry = getEnvInt("CACHE_ITEM_CATEGORY_EXPIRY", 3600)
	cacheItemVariantExpiry = getEnvInt("CACHE_ITEM_VARIANT_EXPIRY", 3600)
	cacheItemVariantOptionExpiry = getEnvInt("CACHE_ITEM_VARIANT_OPTION_EXPIRY", 3600)
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

func GetTokenExpiry() int {
	return tokenExpiry
}

func GetTokenSecret() string {
	return tokenSecret
}

func GetDBURL() string {
	return dbURL
}

func GetCacheUserExpiry() int {
	return cacheUserExpiry
}

func GetCacheTableExpiry() int {
	return cacheTableExpiry
}

func GetCacheMerchantExpiry() int {
	return cacheMerchantExpiry
}

func GetCacheItemExpiry() int {
	return cacheItemExpiry
}

func GetCacheItemCategoryExpiry() int {
	return cacheItemCategoryExpiry
}

func GetCacheItemVariantExpiry() int {
	return cacheItemVariantExpiry
}

func GetCacheItemVariantOptionExpiry() int {
	return cacheItemVariantOptionExpiry
}

func GetTimeout() int {
	return timeout
}
