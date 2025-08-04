// internal/snowflake/writer.go
package snowflake

import (
	"fmt"
	"log"
	"strings"
)

// UpsertRecord deletes and reinserts a record by ID
func (c *Client) UpsertRecord(table string, fields []string, values map[string]interface{}) error {
	id, ok := values["Id"].(string)
	if !ok || id == "" {
		return fmt.Errorf("missing or invalid Id field in values")
	}

	// DELETE existing
	delQuery := fmt.Sprintf("DELETE FROM %s WHERE Id = ?", table)
	_, err := c.DB.Exec(delQuery, id)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	// INSERT new record
	placeholders := strings.Repeat("?,", len(fields))
	placeholders = strings.TrimRight(placeholders, ",")
	columns := strings.Join(fields, ", ")

	insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	args := make([]interface{}, len(fields))
	for i, f := range fields {
		args[i] = values[f]
	}

	_, err = c.DB.Exec(insertQuery, args...)
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	log.Printf("✅ Upserted %s: %s", table, id)
	return nil
}

// DeleteByID removes a record by Id
func (c *Client) DeleteByID(table, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = ?", table)
	_, err := c.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete by ID failed: %w", err)
	}
	log.Printf("🗑️ Deleted from %s: %s", table, id)
	return nil
}
