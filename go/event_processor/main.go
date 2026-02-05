package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	fintech "github.com/sapliy/fintech-sdk-go"
)

func main() {
	apiKey := os.Getenv("SAPLIY_API_KEY")
	if apiKey == "" {
		apiKey = "sk_test_123"
	}

	// Initialize the client
	client := fintech.NewClient(apiKey)

	// Simulate event processing loop
	// In a real app, this would consume from a queue (SQS, Kafka, RabbitMQ)
	// Here we just list events periodically

	fmt.Println("🚀 Go Event Processor started")
	fmt.Println("Waiting for events...")

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			// Poll for recent events (using our mock Listing-like behavior if available,
			// or just simulate processing for example purposes since we don't have ListEvents in SDK yet?
			// Actually Client structure has Webhooks, Payments etc.
			// Let's assume we are acting as a worker processing specific payment logic.

			// Mock: trigger a manual check
			processPendingPayments(client)

		case <-quit:
			fmt.Println("Shutting down...")
			return
		}
	}
}

func processPendingPayments(client *fintech.Client) {
	// Example: In a real scenario, we might list pending payments.
	// For this example, we'll just print a status check.
	// Since we don't have ListPayments exposed in the simple SDK yet (based on sapliyio.go),
	// we will simulate fetching a known payment.

	ctx := context.Background()

	// Try to fetch a dummy payment
	payment, err := client.Payments.Get(ctx, "pay_123456")
	if err != nil {
		// Expected if it doesn't exist
		// fmt.Printf("Poll: No new payments found (checked pay_123456: %v)\n", err)
	} else {
		fmt.Printf("Checked Payment %s: Status=%s\n", payment.ID, payment.Status)
	}
}

// Example of defining an event struct for decoding webhooks
type WebhookEvent struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func HandleWebhook(payload []byte) {
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Failed to parse event: %v", err)
		return
	}

	switch event.Type {
	case "payment.succeeded":
		fmt.Printf("Processing successful payment: %v\n", event.Data["id"])
	case "refund.created":
		fmt.Printf("Processing refund: %v\n", event.Data["id"])
	}
}
