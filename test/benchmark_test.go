package main

import (
    "testing"
    "salesforce-etl-ai/snowflake"
)

func BenchmarkUpsertRecord(b *testing.B) {
    dummyFields := []string{"Id", "Name", "StageName", "Amount", "CloseDate"}
    dummyValues := map[string]interface{}{
        "Id":        "0061X000009uN8YQAU",
        "Name":      "Benchmark Deal",
        "StageName": "Prospecting",
        "Amount":    10000,
        "CloseDate": "2025-09-30",
    }

    client := &snowflake.Client{DB: nil}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = client.UpsertRecord("Opportunity", dummyFields, dummyValues)
    }
}
