package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getRequestedRuleNumber(body []byte, r *http.Request) (*RuleNumber) {
    var requestedRuleNumber RuleNumber
    r.Body = io.NopCloser(bytes.NewBuffer(body))
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&requestedRuleNumber)
    if err != nil {
        log.Fatal(err)
    }
    return &requestedRuleNumber
}
