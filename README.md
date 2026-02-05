# Sapliy Fintech Examples

This directory contains sample applications demonstrating how to integrate with the Sapliy Fintech Ecosystem using various languages.

## 📦 Examples

### 1. Node.js Checkout (`/node`)
A complete checkout flow using Express.js and the Fintech Node.js SDK.

- **Features:** Payment Intents, Webhook Handling, Checkout UI
- **Run:**
  ```bash
  cd node
  npm install
  npm start
  ```
- **URL:** http://localhost:3000

### 2. Python Webhook Handler (`/python`)
A simple Flask application that verifies and processes webhook events.

- **Features:** Signature verification, Event typing
- **Run:**
  ```bash
  cd python
  pip install -r requirements.txt
  python webhook_handler.py
  ```
- **URL:** http://localhost:5000/webhook

### 3. Go Event Processor (`/go`)
A backend worker service that processes events using the Go SDK.

- **Features:** Event loop, API Client usage
- **Run:**
  ```bash
  cd go
  go mod tidy
  go run event_processor/main.go
  ```

## 🔑 Configuration

All examples require the following environment variables:

```bash
export SAPLIY_API_KEY="sk_test_..."
export SAPLIY_WEBHOOK_SECRET="whsec_..."
```
