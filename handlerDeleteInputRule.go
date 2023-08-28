package main

import (
	"net/http"
	"strconv"
)

func (s *ApiServer) HandlerDeleteInputRule(w http.ResponseWriter, r *http.Request) error {

    // Get numRule from url
    numRuleStr, err := GetUrlParameter(r, "numRule")
    numRule, err := strconv.Atoi(numRuleStr)
    if err != nil {
        return &ApiError{ErrorMsg: "numRule is not an integer value"}
    }

    // Validate numRule
    ValidateRuleNumber(TableFilter, ChainInput, &numRule)

    // Delete the rule using numRule
    err = DeleteRule(TableFilter, ChainInput, &numRule)
    if err != nil {
        return err
    }

    // Write Response
    WriteJSON(w, http.StatusAccepted, struct{}{})
    return nil
}
