package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	sapliyio "github.com/sapliy/fintech-sdk-go"
)

func main() {
	client := sapliyio.NewClient(os.Getenv("SAPLIY_API_KEY"))
	webhookSecret := os.Getenv("SAPLIY_WEBHOOK_SECRET")

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		payment, err := client.Payments.CreateIntent(context.Background(), &sapliyio.PaymentIntentRequest{
			Amount:   1000,
			Currency: "USD",
			ZoneID:   "zone_123", // Optional zone scoping
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(payment)
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		sig := r.Header.Get("X-Sapliy-Signature")
		event, err := client.Webhooks.ConstructEvent(payload, sig, webhookSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received event: %s", event.Type)

		if event.Type == "payment.succeeded" {
			// Handle payment success
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
