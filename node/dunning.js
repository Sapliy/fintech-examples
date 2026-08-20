#!/usr/bin/env node
// Sapliy Operational Playbooks — example.
// Prints the Revenue Recovery & Dunning retry schedule and a sample
// refund-approval policy decision. Mirrors the sane defaults from
// sapliy-ecosystem's internal/playbook. No dependencies — Node only.

const FIRST_RETRY_HOURS = 5;
const RETRY_STEP_HOURS = 48;
const FINAL_RETRY_HOURS = 96;
const MAX_RETRIES = 4;

const AUTO_APPROVE_UNDER_CENTS = 100000;
const REQUIRE_APPROVAL_OVER_CENTS = 100000;
const MAX_REFUND_DAYS = 90;

function evaluateRefund(amountCents, daysSinceCharge) {
  // Deterministic refund-approval gate matching refund-approval-policy.
  if (daysSinceCharge > MAX_REFUND_DAYS) {
    return `DENIED — beyond ${MAX_REFUND_DAYS}-day refund window`;
  }
  if (amountCents > REQUIRE_APPROVAL_OVER_CENTS) {
    return 'REQUIRE APPROVAL — over $1,000 manager-approval threshold';
  }
  if (amountCents <= AUTO_APPROVE_UNDER_CENTS) {
    return 'AUTO-APPROVE — under $1,000 auto-approve threshold';
  }
  return 'REQUIRE APPROVAL';
}

function main() {
  console.log('Sapliy Operational Playbooks — Example');
  console.log('======================================');

  console.log('\n[1] Revenue Recovery & Dunning — retry schedule');
  const first = FIRST_RETRY_HOURS;
  const step = RETRY_STEP_HOURS;
  console.log(`First retry:       +${first}h (~22% recovery expected)`);
  console.log(`Second retry:      +${first + step}h`);
  console.log(`Third retry:       +${first + 2 * step}h`);
  console.log(`Final retry:       +${FINAL_RETRY_HOURS}h`);
  console.log(`Max retries:       ${MAX_RETRIES}`);
  console.log(`Channels:          email`);
  console.log(`Magic-link dunning: true`);

  console.log('\n[2] Refund approval — policy gate decision');
  const amountCents = 125000; // $1,250.00
  const daysSinceCharge = 45;
  console.log(`Refund amount:     $${(amountCents / 100).toFixed(2)} (${amountCents} cents)`);
  console.log(`Days since charge: ${daysSinceCharge}`);
  console.log(`Decision:          ${evaluateRefund(amountCents, daysSinceCharge)}`);
}

main();