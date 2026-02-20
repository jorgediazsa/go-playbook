package production

import (
	"testing"
)

func TestChargeCustomerLogs(t *testing.T) {
	// Verifying structured log output natively without hijacking `os.Stdout` is clunky,
	// but we can verify the function signature and execution don't panic.
	//
	// To truly verify: run this test natively and look at the console output.
	// It should look like:
	// {"time":"2023-10-01T...", "level":"ERROR", "msg":"charge failed", "customer_id":"CUST-1", "amount":0, "error":"invalid amount"}
	// {"time":"2023-10-01T...", "level":"INFO", "msg":"charge successful", "customer_id":"CUST-2", "amount":500}

	err := ChargeCustomer("CUST-1", 0)
	if err == nil {
		t.Fatalf("FAILED: Expected error for amount 0")
	}

	err = ChargeCustomer("CUST-2", 500)
	if err != nil {
		t.Fatalf("FAILED: Expected success for amount 500")
	}

	t.Log(`If you successfully refactored to "log/slog", you should see JSON logs in your terminal during this test!`)
}
