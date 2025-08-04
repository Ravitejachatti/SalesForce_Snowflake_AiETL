// internal/etl/ai_etl.go
package etl

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// GenerateSummary generates a plain-text summary of the record change
func GenerateSummary(entity string, values map[string]interface{}) string {
	id, _ := values["Id"].(string)
	summary := fmt.Sprintf("🔎 %s change summary for ID: %s\n", entity, id)

	for k, v := range values {
		if strings.EqualFold(k, "Id") {
			continue
		}
		valStr := fmt.Sprintf("%v", v)
		if len(valStr) > 80 {
			valStr = valStr[:80] + "..."
		}
		summary += fmt.Sprintf("- %s: %s\n", k, valStr)
	}

	return summary
}

// LogSummary outputs the AI summary to console or external sinks (Slack, etc.)
func LogSummary(entity string, values map[string]interface{}) {
	summary := GenerateSummary(entity, values)
	log.Println("🧠 AI Summary:\n" + summary)
}

// SerializeToJSON returns a clean JSON of the record change (optional for AI input or logs)
func SerializeToJSON(values map[string]interface{}) string {
	out, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(out)
}
