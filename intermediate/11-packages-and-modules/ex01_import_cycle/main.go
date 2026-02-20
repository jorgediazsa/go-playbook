package main

// Context: Import Cycles
// You are building an e-commerce platform. The `user` package manages authentication
// and profile data. The `order` package manages shopping carts and purchases.
// An Order belongs to a User, and a User has a history of Orders.
//
// Why this matters: In Go, if Package A imports Package B, Package B CANNOT
// import Package A. This forces a strict Directed Acyclic Graph (DAG) of dependencies.
//
// This exercise ships with:
// - A FIXED design (default build) that avoids cycles by extracting shared models into `domain/`.
// - A BROKEN design (build tag `broken`) that recreates the import cycle for study.
//
// Try:
//   go test ./...
//   go test -tags broken ./...   # observe "import cycle not allowed"
//
// Requirements (for the learner):
// 1. Understand why the `broken` version fails.
// 2. Understand why extracting shared types into `domain` breaks the cycle.
// 3. Keep the dependency graph acyclic as you evolve the codebase.

import (
	"fmt"

	"go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/domain"
	"go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/order"
	"go-playbook/intermediate/11-packages-and-modules/ex01_import_cycle/user"
)

func main() {
	u := user.New(domain.UserID("user_001"), "Alice")
	o := order.CreateOrder(99.99, u)

	u2 := user.WithOrders{
		User:   u,
		Orders: []domain.Order{o},
	}

	fmt.Printf("Success! User %s created %d order(s).\n", u2.User.Name, len(u2.Orders))
}
