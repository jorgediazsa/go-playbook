package order

import "go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/domain"

// CreateOrder constructs a new domain.Order.
// In real systems, prefer IDs or shared domain models to avoid cyclic dependencies.
func CreateOrder(amount float64, owner domain.User) domain.Order {
	return domain.Order{
		ID:     domain.OrderID("ord-123"),
		Amount: amount,
		Owner:  owner,
	}
}
