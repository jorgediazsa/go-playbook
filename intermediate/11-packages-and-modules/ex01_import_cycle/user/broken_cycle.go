//go:build broken

package user

// This file intentionally demonstrates a cyclic dependency.
// It is excluded from normal builds. To see the cycle, build with:
//   go test -tags broken ./...
//
// BUG: This package imports "order", creating a cycle.
import "go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/order"

type User struct {
	ID     string
	Name   string
	Orders []order.Order
}

func GetUserOrders(u User) []order.Order {
	return u.Orders
}
