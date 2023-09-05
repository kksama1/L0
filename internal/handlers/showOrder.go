package handlers

import (
	"fmt"
	"level0/internal/cash"
	"net/http"
)

// ShowOrder displays order with given id from the cash.
func ShowOrder(w http.ResponseWriter, r *http.Request) {
	cash.Init()
	key := r.FormValue("Id")
	if cash.Cash[key] != "" {
		fmt.Fprintf(w, cash.Cash[key])
	} else {
		fmt.Fprintf(w, "Неверный Id")
	}
}
