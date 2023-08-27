package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiServer struct {
    listenAddr string
}

type ApiError struct {
    ErrorMsg string `json:"error,omitempty"`
}

func (e *ApiError) Error() string {
    return fmt.Sprintf("Error: %v", e.ErrorMsg)
}

func NewServer(listenAddr string) *ApiServer {
    return &ApiServer{
        listenAddr: listenAddr,
    }
}

func (s *ApiServer) Run() {
    router := chi.NewRouter()

    router.Mount("/v1", getV1Router())

    log.Printf("Server starting on port %s", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}
