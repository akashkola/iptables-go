package main

import (
	//"log"
	"net/http"
)

func (s *ApiServer) HandlerAddInputRule(w http.ResponseWriter, r *http.Request) error {

    // read request body into the buffer
    body, err := ReadRequestBody(r)
    if err != nil {
        return err
    }

    // TODO: validate filter table rule options
    // get filter table rule
    var filterTableRule FilterTableRule
    err = GetFilterTableRuleFromRequest(body, &filterTableRule)
    if err != nil {
        return err
    }

    // validate requsted rule number 
    err = ValidateRuleNumber(TableFilter, ChainInput, filterTableRule.RuleNumber, false)
    if err != nil {
        return err
    }

    // add rule to the chain
    err = AddRule(TableFilter, ChainInput, &filterTableRule)
    if err != nil {
        return err
    }
    
    WriteJSON(w, http.StatusOK, filterTableRule)
    return nil
}
