package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/sapliy/fintech-sdk-go"
)

func main() {
	client := fintech.NewClient(os.Getenv("SAPLIY_API_KEY"))

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		payment, err := client.Payments.Create(context.Background(), &fintech.CreateChargeRequest{
			Amount:   1000,
			Currency: "USD",
			SourceID: "tok_visa",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(payment)
	})

	log.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
