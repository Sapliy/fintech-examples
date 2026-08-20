#!/usr/bin/env python3
"""Sapliy Operational Playbooks — example.

Prints the Revenue Recovery & Dunning retry schedule and a sample
refund-approval policy decision. Mirrors the sane defaults from
sapliy-ecosystem's internal/playbook. Standard library only.
"""

from __future__ import annotations

from dataclasses import dataclass, field

FIRST_RETRY_HOURS = 5
RETRY_STEP_HOURS = 48
FINAL_RETRY_HOURS = 96
MAX_RETRIES = 4

AUTO_APPROVE_UNDER_CENTS = 100_000
REQUIRE_APPROVAL_OVER_CENTS = 100_000
MAX_REFUND_DAYS = 90


@dataclass
class DunningConfig:
    max_retries: int = MAX_RETRIES
    first_retry_hours: int = FIRST_RETRY_HOURS
    retry_step_hours: int = RETRY_STEP_HOURS
    final_retry_hours: int = FINAL_RETRY_HOURS
    channels: list[str] = field(default_factory=lambda: ["email"])
    magic_link: bool = True


def evaluate_refund(amount_cents: int, days_since_charge: int) -> str:
    """Deterministic refund-approval gate matching refund-approval-policy."""
    if days_since_charge > MAX_REFUND_DAYS:
        return f"DENIED — beyond {MAX_REFUND_DAYS}-day refund window"
    if amount_cents > REQUIRE_APPROVAL_OVER_CENTS:
        return "REQUIRE APPROVAL — over $1,000 manager-approval threshold"
    if amount_cents <= AUTO_APPROVE_UNDER_CENTS:
        return "AUTO-APPROVE — under $1,000 auto-approve threshold"
    return "REQUIRE APPROVAL"


def main() -> None:
    cfg = DunningConfig()
    print("Sapliy Operational Playbooks — Example")
    print("======================================")

    print("\n[1] Revenue Recovery & Dunning — retry schedule")
    first = cfg.first_retry_hours
    step = cfg.retry_step_hours
    print(f"First retry:       +{first}h (~22% recovery expected)")
    print(f"Second retry:      +{first + step}h")
    print(f"Third retry:       +{first + 2 * step}h")
    print(f"Final retry:       +{cfg.final_retry_hours}h")
    print(f"Max retries:       {cfg.max_retries}")
    print(f"Channels:          {', '.join(cfg.channels)}")
    print(f"Magic-link dunning: {cfg.magic_link}")

    print("\n[2] Refund approval — policy gate decision")
    amount_cents = 125_000  # $1,250.00
    days_since_charge = 45
    print(f"Refund amount:     ${amount_cents / 100:,.2f} ({amount_cents} cents)")
    print(f"Days since charge: {days_since_charge}")
    print(f"Decision:          {evaluate_refund(amount_cents, days_since_charge)}")


if __name__ == "__main__":
    main()