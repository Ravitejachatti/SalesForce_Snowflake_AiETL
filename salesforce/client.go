package salesforce

import (
	"fmt"
	"log"

	"salesforce-etl-ai/config"
	"github.com/simpleforce/simpleforce"
)

type Client struct {
	SF   *simpleforce.Client
	Conf *config.Config
}

func NewClient(conf *config.Config) (*Client, error) {
	log.Println("🔧 Initializing Salesforce client...")

	log.Printf("📡 Using instance URL: %s", conf.SalesforceInstanceURL)
	client := simpleforce.NewClient(conf.SalesforceInstanceURL, "", "")

	log.Println("🔐 Attempting login to Salesforce...")
	err := client.LoginPassword(conf.SalesforceUsername, conf.SalesforcePassword, conf.SalesforceToken)
	if err != nil {
		log.Printf("❌ Login failed for user: %s", conf.SalesforceUsername)
		return nil, fmt.Errorf("Salesforce login failed: %w", err)
	}

	log.Println("✅ Connected to Salesforce API")
	return &Client{SF: client, Conf: conf}, nil
}

func (c *Client) GetRecord(entity, id string) (*simpleforce.SObject, error) {
	log.Printf("🔍 Fetching record from entity: %s, ID: %s", entity, id)

	sobj := c.SF.SObject(entity)
	record := sobj.Get(id)

	if record == nil {
		return nil, fmt.Errorf("failed to fetch %s record (%s): returned nil", entity, id)
	}

	log.Println("✅ Successfully fetched record")
	return record, nil
}