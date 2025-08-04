// internal/etl/processor.go
package etl

import (
	"context"
	
	"fmt"
	"log"

	"salesforce-etl-ai/salesforce"
	"salesforce-etl-ai/snowflake"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"   
)

type ChangeEventHeader struct {
	RecordIds  []string `json:"recordIds"`
	ChangeType string   `json:"changeType"`
	EntityName string   `json:"entityName"`
}

type Payload struct {
	ChangeEventHeader ChangeEventHeader `json:"ChangeEventHeader"`
}

type ChangeEvent struct {
	Payload Payload `json:"payload"`
}

// ListenAndServe connects to WebSocket and handles incoming CDC events
func ListenAndServe(ctx context.Context, sf *salesforce.Client, db *snowflake.Client) error {
	wsURL := "wss://your-streaming-endpoint" // replace with actual
	conn, _, err := websocket.Dial(ctx, wsURL, nil)
	if err != nil {
		return fmt.Errorf("WebSocket dial failed: %w", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "graceful shutdown")

	log.Println("📡 Listening for Salesforce CDC events...")

	for {
		var evt ChangeEvent
		err := wsjson.Read(ctx, conn, &evt)
		if err != nil {
			log.Printf("⚠️ WebSocket read failed: %v", err)
			continue
		}

		go processEvent(sf, db, evt)
	}
}

func processEvent(sf *salesforce.Client, db *snowflake.Client, evt ChangeEvent) {
	h := evt.Payload.ChangeEventHeader
	if len(h.RecordIds) == 0 {
		log.Println("⚠️ No recordIds in change event")
		return
	}
	recordId := h.RecordIds[0]
	entity := h.EntityName
	change := h.ChangeType

	log.Printf("🔔 Change Detected: %s [%s] - %s", entity, recordId, change)

	fields := defaultFields(entity)
	table := entity

	if change == "DELETE" {
		err := db.DeleteByID(table, recordId)
		if err != nil {
			log.Printf("❌ Delete failed: %v", err)
		}
		return
	}

	record, err := sf.GetRecord(entity, recordId)
	if err != nil {
		log.Printf("❌ Fetch from Salesforce failed: %v", err)
		return
	}

	values := map[string]interface{}{}
	fieldsMap, ok := (*record)["fields"].(map[string]interface{})
	if !ok {
		log.Printf("⚠️ Could not parse fields from record: %+v", record)
		return
	}
	for _, f := range fields {
		if val, ok := fieldsMap[f]; ok {
			values[f] = val
		} else {
			values[f] = nil
		}
	}

	err = db.UpsertRecord(table, fields, values)
	if err != nil {
		log.Printf("❌ Upsert failed: %v", err)
	}
}

func defaultFields(entity string) []string {
	switch entity {
	case "Case":
		return []string{"Id", "Subject", "Status", "CreatedDate"}
	case "Opportunity":
		return []string{"Id", "Name", "StageName", "Amount", "CloseDate"}
	default:
		return []string{"Id"}
	}
}
