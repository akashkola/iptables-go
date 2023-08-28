package main

import "github.com/go-chi/chi/v5"

func getV1Router(s *ApiServer) *chi.Mux {
    router := chi.NewRouter()

    router.Get("/input", MakeHttpHandler(s.HandlerGetInputRules))
    router.Post("/input", MakeHttpHandler(s.HandlerAddInputRule))
    router.Delete("/input/{numRule}", MakeHttpHandler(s.HandlerDeleteInputRule))

    return router

}
