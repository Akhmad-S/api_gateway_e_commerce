package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string //development, staging, production

	HTTPPort    string

	CatalogServiceGrpcHost string
	CatalogServiceGrpcPort string

	OrderServiceGrpcHost string
	OrderServiceGrpcPort string

	AuthServiceGrpcHost string
	AuthServiceGrpcPort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "e-commerce"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7071"))

	config.CatalogServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATALOG_SERVICE_GRPC_HOST", "localhost"))
	config.CatalogServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATALOG_SERVICE_GRPC_PORT", ":9001"))

	config.OrderServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_HOST", "localhost"))
	config.OrderServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_PORT", ":9002"))

	config.AuthServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_GRPC_HOST", "localhost"))
	config.AuthServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_GRPC_PORT", ":9003"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
