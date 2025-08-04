# 🚀 AI-Powered Salesforce → Snowflake ETL Tool

## Overview

This project implements a real-time ETL (Extract, Transform, Load) pipeline that captures Salesforce Change Events (CDC) for `Case` and `Opportunity` objects and synchronizes them into Snowflake. The pipeline is designed with production-readiness in mind and includes optional AI-powered summarization for audit logs.

---

## 🔧 Features

* ✅ Real-time Salesforce CDC via WebSocket or CometD (plug-in layer)
* ✅ Upserts into Snowflake tables (`Case`, `Opportunity`)
* ✅ Automatic DELETE handling for removed records
* ✅ AI Summary Log (OpenAI/GPT optional for `AUDIT_LOG` table)
* ✅ Graceful shutdown and error recovery
* ✅ Retry logic and robust error handling
* ✅ Fully Dockerized & configurable via `.env`

---

## 📁 Project Structure

```
salesforce-etl-ai/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── salesforce/         # Login and SOQL queries
│   ├── snowflake/          # Insert/Upsert/Delete
│   ├── etl/                # ChangeEvent handler & AI
│   └── models/             # Structs for ChangeEvent
├── Dockerfile
├── go.mod
├── .env
└── README.md
```

---

## 🧠 AI-Powered Add-On (Optional)

Generates plain-English summaries of changes, e.g.:

```
"[Opportunity] ID=0061... was updated: Amount changed to $12,500, Stage moved to 'Closed Won'."
```

These can be inserted into a Snowflake `AUDIT_LOG` table or published to Slack via webhook.

---

## 🔧 Requirements

* Go 1.20+
* Snowflake Account & Table
* Salesforce Dev Account
* \[Optional] OpenAI Key for audit summaries

---

## 🚀 Quickstart

1. Clone and configure:

```bash
git clone https://github.com/yourname/salesforce-etl-ai
cd salesforce-etl-ai
```

2. Fill in your Salesforce and Snowflake credentials in `.env`

3. Run:

```bash
go run cmd/main.go
```

---

## 🐳 Docker

```bash
docker build -t etl-pipeline .
docker run --env-file .env etl-pipeline
```

---

## 🧪 Tests (Planned)

* Unit tests for SOQL parsing and inserts
* Integration test: mocked Salesforce stream
* E2E test via GitHub Actions

---

## ✨ Credits

Built by Venkata Ravi Teja Chatti using Go, Snowflake, Salesforce, OpenAI. Inspired by industry-ready ETL needs.

---

## 📬 Contact

For questions or feature requests, open an issue or contact \[[ravitejachatti@gmail.com]].

---

