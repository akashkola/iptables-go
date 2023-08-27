package main

import (
	"log"
	"net/http"
)

func handlerGetInputRules(w http.ResponseWriter, r *http.Request) {
    rules, err := GetRules(TableFilter, ChainInput)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, ApiError{ErrorMsg: "unable to read rules"})
        log.Println(err)
        return
    }
    
    WriteJSON(w, http.StatusOK, rules)
}
