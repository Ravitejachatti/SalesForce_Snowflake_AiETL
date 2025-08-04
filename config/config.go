// internal/config/config.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all environment variables
type Config struct {
	SalesforceUsername string
	SalesforcePassword string
	SalesforceToken    string
	SalesforceInstanceURL      string

	SnowflakeUser      string
	SnowflakePassword  string
	SnowflakeAccount   string
	SnowflakeDatabase  string
	SnowflakeSchema    string
	SnowflakeWarehouse string

	EnableMetrics bool
	Port          string
}

// LoadConfig loads the .env file or uses system environment variables
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	get := func(key string) string {
		v := os.Getenv(key)
		if v == "" {
			log.Printf("⚠️  missing env: %s", key)
		}
		return v
	}

	enableMetrics := os.Getenv("ENABLE_METRICS") == "true"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := &Config{
		SalesforceUsername: get("SALESFORCE_USERNAME"),
		SalesforcePassword: get("SALESFORCE_PASSWORD"),
		SalesforceToken:    get("SALESFORCE_TOKEN"),
		SalesforceInstanceURL:      get("SALESFORCE_INSTANCE_URL"),

		SnowflakeUser:      get("SNOWFLAKE_USER"),
		SnowflakePassword:  get("SNOWFLAKE_PASSWORD"),
		SnowflakeAccount:   get("SNOWFLAKE_ACCOUNT"),
		SnowflakeDatabase:  get("SNOWFLAKE_DATABASE"),
		SnowflakeSchema:    get("SNOWFLAKE_SCHEMA"),
		SnowflakeWarehouse: get("SNOWFLAKE_WAREHOUSE"),

		EnableMetrics:      enableMetrics,
		Port:               port,

		
	}

	return cfg, nil
}

// DSN returns the Snowflake DSN string
func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s/%s?warehouse=%s",
		c.SnowflakeUser,
		c.SnowflakePassword,
		c.SnowflakeAccount,
		c.SnowflakeDatabase,
		c.SnowflakeSchema,
		c.SnowflakeWarehouse,
	)
}
