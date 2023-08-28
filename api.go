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
    // TODO: add cache store
}

type ApiError struct {
    ErrorMsg string `json:"error,omitempty"`
}

type ApiFunc func (http.ResponseWriter, *http.Request) error

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

    router.Mount("/v1", getV1Router(s))

    log.Printf("Server starting on port %s", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(apiFunc ApiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := apiFunc(w, r); err != nil {
            // TODO: handle errors
            WriteJSON(w, http.StatusBadRequest, err)
        }
    }
}
