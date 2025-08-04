// internal/snowflake/client.go
package snowflake

import (
	"database/sql"
	"fmt"
	"log"

	"salesforce-etl-ai/config"

	_ "github.com/snowflakedb/gosnowflake"
)

// Client holds the snowflake db connection and config

type Client struct {
	DB   *sql.DB
	Conf *config.Config
}

// Connect opens a Snowflake database connection
func Connect(conf *config.Config) (*Client, error) {
	dsn := conf.DSN()
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open Snowflake connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Snowflake: %w", err)
	}

	log.Println("✅ Connected to Snowflake")
	return &Client{DB: db, Conf: conf}, nil
}
