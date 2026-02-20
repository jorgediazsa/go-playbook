package production

// Context: Structured Logging (`log/slog`)
// You are maintaining a payment gateway. The old monolithic logger just calls
// `fmt.Printf` or `log.Println` with massive unstructured strings.
//
// Why this matters: When a customer complains that transaction "TX-999" failed
// at 3:00 AM, you cannot easily search for "TX-999" if your logs are just text blobs.
// Structured logs (JSON) allow log aggregators (Datadog/Elasticsearch) to index
// fields natively, allowing you to query `transaction_id : "TX-999" AND level : "ERROR"`.
//
// Requirements:
// 1. Refactor `ChargeCustomer` to stop using `log.Printf`.
// 2. Use `log/slog`.
// 3. Create a JSON logger: `logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))`.
//    (In production, you'd configure this once globally or pass it via Dependency Injection).
// 4. Record the success or failure using `logger.Info` or `logger.Error`.
// 5. You MUST include structured key-value pairs in the log:
//    `"customer_id"` and `"amount"`, plus `"error"` if it fails.

import (
	"errors"
	"log"
)

func ChargeCustomer(customerID string, amount int) error {
	// BUG: Unstructured logging is basically useless at scale.
	// TODO: Replace with `slog` JSON logging. Include the customerID and amount as structured keys!

	if amount <= 0 {
		log.Printf("[ERROR] Failed to charge customer %s for amount %d: invalid amount", customerID, amount)
		return errors.New("invalid amount")
	}

	log.Printf("[INFO] Successfully charged customer %s for amount %d", customerID, amount)
	return nil
}
