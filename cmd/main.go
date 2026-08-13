package main

import (
	"benchmarking-scylladb-dynamodb/domain"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	testCustomer := domain.Order{
		ID:         "123",
		CustomerID: "1",
		Status:     "COMPRADO",
		Total:      251.52,
	}

	fmt.Println("Teste Customer: ID: " + testCustomer.ID)
}
