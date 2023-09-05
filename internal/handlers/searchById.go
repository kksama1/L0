package handlers

import "net/http"

// SearchById displays form to search order via id.
func SearchById(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
