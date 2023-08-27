package main

import (
	"errors"
	"log"
	"net/http"
)

func handlerAddInputRule(w http.ResponseWriter, r *http.Request) {

    // read request body into the buffer
    body, err := ReadRequestBody(r)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, errors.New("Error: Unable to read body from request"))
        log.Println(err)
        return
    }

    // get filter table rule
    var filterTableRule FilterTableRule
    err = GetFilterTableRuleFromRequest(body, &filterTableRule)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, errors.New("Error: Unable to read rule from request"))
        log.Println(err)
        return
    }

    // TODO: validate filter table rule value

    // get requsted rule number
    var ruleNumber RuleNumber
    err = GetRequestedRuleNumber(body, &ruleNumber)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, errors.New("Error: Unable to read rule number from request"))
        log.Println(err)
        return
    }

    err = ValidateRuleNumber(TableFilter, ChainInput, *ruleNumber.RuleNum)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err.Error())
        log.Println(err)
        return
    }

    // add rule to the chain
    err = AddRule(TableFilter, ChainInput, &filterTableRule, ruleNumber.RuleNum)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err.Error())
        log.Println(err)
        return
    }
    
    WriteJSON(w, http.StatusOK, filterTableRule)
}
