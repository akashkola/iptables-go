package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerGetInputRules(w http.ResponseWriter, r *http.Request) {
    rules, err := getRules(TableFilter, ChainInput)
    if err != nil {
        log.Fatal(err)
    }
    
    w.Header().Add("Content-Type", "application/json")
    jsonRules, err := json.Marshal(rules)
    if err != nil{
        log.Fatal(err)
    }
    w.Write(jsonRules)
}
