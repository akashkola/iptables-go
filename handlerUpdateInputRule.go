package main

import (
	"net/http"
	"strconv"
)

func (s *ApiServer) HandlerUpdateInputRule(w http.ResponseWriter, r *http.Request) error {
    // Get ruleNumber param from URL
    ruleNumberStr := GetUrlParameter(r, "ruleNumber")
    ruleNumber, err := strconv.Atoi(ruleNumberStr)
    if err != nil {
        return &ApiError{ErrorMsg: "ruleNumber is not an integer value"}
    }

    // Validate the ruleNum
    err = ValidateRuleNumber(TableFilter, ChainInput, &ruleNumber, true)
    if err != nil {
        return err
    }

    // read filter table rule from request
    body, err := ReadRequestBody(r)
    if err != nil {
        return err
    }
    requestedRule := new(FilterTableRule)
    err = GetFilterTableRuleFromRequest(body, requestedRule)
    if err != nil {
        return err
    }

    if IsEmptyRule(requestedRule, false) {
        return &ApiError{ErrorMsg: "empty rule"}
    }

    // get existingRule using ruleNumber
    existingRule, err := GetRuleByNumRule(TableFilter, ChainInput, &ruleNumber)
    if err != nil {
        return err
    }

    // update rule by deleting and adding updated one with same ruleNumber
    err = DeleteRule(TableFilter, ChainInput, &ruleNumber)
    if err != nil {
        return err
    }
    CopyRule(requestedRule, existingRule)
    existingRule.RuleNumber = &ruleNumber
    AddRule(TableFilter, ChainInput, existingRule)

    return nil
}
