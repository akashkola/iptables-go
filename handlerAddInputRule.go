package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func handlerAddInputRule(w http.ResponseWriter, r *http.Request) {
    bodyBytes, err := io.ReadAll(r.Body)
    
    // get filter table rule
    var filterTableRule FilterTableRule
    if err != nil {
        log.Fatal(err)
    }
    r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&filterTableRule); err != nil {
        log.Fatal(err)
    }

    // get requsted rule number
    requestedRuleNumber := getRequestedRuleNumber(bodyBytes, r)
    
    addRule(TableFilter, ChainInput, &filterTableRule, requestedRuleNumber.RuleNum)

    w.Header().Add("content-type", "application/json")
    jsonRule, err := json.Marshal(filterTableRule)
    if err != nil {
        log.Fatal(err)
    }
    w.Write(jsonRule)
}
