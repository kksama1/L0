package handlers

import (
	"fmt"
	"level0/internal/cash"
	"net/http"
)

// ShowAllOrders displays list of all existing orders in cash.
func ShowAllOrders(w http.ResponseWriter, r *http.Request) {
	var output string
	cash.Init()
	for k, _ := range cash.Cash {
		output += fmt.Sprintf("%s\n\n", cash.Cash[k])
	}
	fmt.Fprintf(w, output)
}
