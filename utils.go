package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func ReadRequestBody(r *http.Request) ([]byte, error) {
    return io.ReadAll(r.Body)
}

func GetFilterTableRuleFromRequest(body []byte, filterTableRule *FilterTableRule) error {
    reader := strings.NewReader(string(body))
    decoder := json.NewDecoder(reader)
    if err := decoder.Decode(filterTableRule); err != nil {
       return err 
    }
    return nil
}

func FilterEmptyString(str *[]string) []string {
    filteredStringSlice := make([]string, 0)
    for _, strValue :=  range *str {
        if strValue == "" {
            continue
        }
        filteredStringSlice = append(filteredStringSlice, strValue)
    }
    return filteredStringSlice
}

func ValidateRuleNumber(table Table, chain Chain, ruleNumber *int) error {

    if ruleNumber == nil {
        ruleNumber = new(int)
        *ruleNumber = 1
    }

    rules, err := GetRules(table, chain)
	if err != nil {
		return err
	}

	rulesLength := len(rules)

    if *ruleNumber > rulesLength + 1 {
		return &ApiError{ErrorMsg: "rule number is too big"}
	}

	if *ruleNumber <= 0 {
		return &ApiError{ErrorMsg: "invalid rule number"}
	}

	return nil
}

func GetUrlParameter(r *http.Request, parameter string) (string, error) {
    paramValue := chi.URLParam(r, parameter)
    return paramValue, nil
}













