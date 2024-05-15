package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/rest"
)

func main() {
	http.HandleFunc("/", rest.RootHandler)
	http.HandleFunc("/view/", rest.MakeHandler(rest.ViewHandler))
	http.HandleFunc("/edit/", rest.MakeHandler(rest.EditHandler))
	http.HandleFunc("/save/", rest.MakeHandler(rest.SaveHandler))
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
