package main

import (
	"log"
	"net/http"
)

func handlerAddInputRule(w http.ResponseWriter, r *http.Request) {

    // read request body into the buffer
    body, err := ReadRequestBody(r)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, ApiError{ErrorMsg: "Error: Unable to read request"})
        log.Println(err)
        return
    }

    // TODO: validate filter table rule options
    // get filter table rule
    var filterTableRule FilterTableRule
    err = GetFilterTableRuleFromRequest(body, &filterTableRule)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, ApiError{ErrorMsg: "Error: Unable to read rule from request"})
        log.Println(err)
        return
    }

    // validate requsted rule number 
    err = ValidateRuleNumber(TableFilter, ChainInput, filterTableRule.RuleNumber)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err.Error())
        log.Println(err)
        return
    }

    // add rule to the chain
    err = AddRule(TableFilter, ChainInput, &filterTableRule)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, ApiError{ErrorMsg: "Error: Unable to add rule"})
        log.Println(err)
        return
    }
    
    WriteJSON(w, http.StatusOK, filterTableRule)
}
