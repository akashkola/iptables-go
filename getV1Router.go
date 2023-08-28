package main

import "github.com/go-chi/chi/v5"

func getV1Router() *chi.Mux {
    router := chi.NewRouter()

    router.Get("/input", MakeHttpHandler(handlerGetInputRules))
    router.Post("/input", MakeHttpHandler(handlerAddInputRule))

    return router

}
