package main

import (
	"net/http"
	"strconv"
)

func (s *ApiServer) HandlerDeleteInputRule(w http.ResponseWriter, r *http.Request) error {

	// Get ruleNumber from url
	ruleNumberStr := GetUrlParameter(r, "ruleNumber")
	ruleNumber, err := strconv.Atoi(ruleNumberStr)
	if err != nil {
		return &ApiError{ErrorMsg: "ruleNumber is not an integer value"}
	}

	// Validate ruleNumber
	err = ValidateRuleNumber(TableFilter, ChainInput, &ruleNumber, true)
	if err != nil {
		return err
	}

	// Delete the rule using ruleNumber
	err = DeleteRule(TableFilter, ChainInput, &ruleNumber)
	if err != nil {
		return err
	}

	// Write Response
	WriteJSON(w, http.StatusAccepted, struct{}{})
	return nil
}
