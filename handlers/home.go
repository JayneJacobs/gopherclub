package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("home page handler reached")
	http.Redirect(w, r, "/login", 302)
}
