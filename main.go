package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func main() {
    router := chi.NewRouter()

    router.Mount("/v1", getV1Router()) 
    
    fmt.Println("Server starting on port 3000")
    http.ListenAndServe(":3000", router)
}
