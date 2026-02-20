package domain

// UserID and OrderID are defined types to avoid accidental mixing.
// This also reinforces how a small shared package can prevent import cycles.
type UserID string
type OrderID string

type User struct {
	ID   UserID
	Name string
}

type Order struct {
	ID     OrderID
	Amount float64
	Owner  User
}
