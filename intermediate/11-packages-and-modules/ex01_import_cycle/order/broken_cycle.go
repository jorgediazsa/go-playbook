//go:build broken

package order

// This file intentionally demonstrates a cyclic dependency.
// It is excluded from normal builds. To see the cycle, build with:
//   go test -tags broken ./...
//
// BUG: This package imports "user", creating a cycle!
import "go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/user"

type Order struct {
	ID     string
	Amount float64
	Owner  *user.User
}

func CreateOrder(amount float64, owner *user.User) Order {
	return Order{
		ID:     "ord-123",
		Amount: amount,
		Owner:  owner,
	}
}
