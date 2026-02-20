package user

import "go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/domain"

// WithOrders models a user + its order history without importing the order package.
// The order history uses domain.Order, which lives in a shared package to avoid cycles.
type WithOrders struct {
	User   domain.User
	Orders []domain.Order
}

func New(id domain.UserID, name string) domain.User {
	return domain.User{ID: id, Name: name}
}

func GetUserOrders(u WithOrders) []domain.Order {
	return u.Orders
}
