// Command dunning prints the Revenue Recovery & Dunning playbook schedule
// and a sample refund-approval decision. It mirrors the sane defaults from
// sapliy-ecosystem's internal/playbook (DunningConfig / RefundApprovalConfig)
// and uses only the standard library — no SDK dependency.
package main

import (
	"fmt"
	"time"
)

// DunningConfig mirrors internal/playbook DunningConfig sane defaults.
type DunningConfig struct {
	MaxRetries      int
	FirstRetryDelay time.Duration
	RetryStepDelay  time.Duration
	FinalRetryDelay time.Duration
	Channels        []string
	MagicLink       bool
}

func defaultDunningConfig() DunningConfig {
	return DunningConfig{
		MaxRetries:      4,
		FirstRetryDelay: 5 * time.Hour,
		RetryStepDelay:  48 * time.Hour,
		FinalRetryDelay: 96 * time.Hour,
		Channels:        []string{"email"},
		MagicLink:       true,
	}
}

// RefundApprovalConfig mirrors internal/playbook RefundApprovalConfig.
type RefundApprovalConfig struct {
	AutoApproveUnderCents    int64
	RequireApprovalOverCents int64
	MaxRefundDays            int
}

func defaultRefundApprovalConfig() RefundApprovalConfig {
	return RefundApprovalConfig{
		AutoApproveUnderCents:    100000,
		RequireApprovalOverCents: 100000,
		MaxRefundDays:            90,
	}
}

func main() {
	cfg := defaultDunningConfig()

	fmt.Println("Sapliy Operational Playbooks — Example")
	fmt.Println("======================================")

	fmt.Println("\n[1] Revenue Recovery & Dunning — retry schedule")
	fmt.Printf("First retry:      +%s (~22%% recovery expected)\n", cfg.FirstRetryDelay)
	fmt.Printf("Second retry:     +%s\n", cfg.FirstRetryDelay+cfg.RetryStepDelay)
	fmt.Printf("Third retry:      +%s\n", cfg.FirstRetryDelay+2*cfg.RetryStepDelay)
	fmt.Printf("Final retry:      +%s\n", cfg.FinalRetryDelay)
	fmt.Printf("Max retries:      %d\n", cfg.MaxRetries)
	fmt.Printf("Channels:         %v\n", cfg.Channels)
	fmt.Printf("Magic-link dunning: %v\n", cfg.MagicLink)

	fmt.Println("\n[2] Refund approval — policy gate decision")
	refund := defaultRefundApprovalConfig()
	amountCents := int64(125000) // $1,250.00
	daysSinceCharge := 45
	fmt.Printf("Refund amount:    $%d.%02d (%d cents)\n", amountCents/100, amountCents%100, amountCents)
	fmt.Printf("Days since charge: %d\n", daysSinceCharge)

	switch {
	case daysSinceCharge > refund.MaxRefundDays:
		fmt.Printf("Decision: DENIED — beyond %d-day refund window\n", refund.MaxRefundDays)
	case amountCents > refund.RequireApprovalOverCents:
		fmt.Printf("Decision: REQUIRE APPROVAL — over $1,000 manager-approval threshold\n")
	case amountCents <= refund.AutoApproveUnderCents:
		fmt.Printf("Decision: AUTO-APPROVE — under $1,000 auto-approve threshold\n")
	}
}