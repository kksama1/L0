package main

import (
	"level0/internal/cash"
	"level0/internal/handlers"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	cash.Init()
	router.HandleFunc("/all", handlers.ShowAllOrders)
	router.HandleFunc("/", handlers.SearchById)
	router.HandleFunc("/postform", handlers.ShowOrder)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
