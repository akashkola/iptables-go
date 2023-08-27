package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ReadRequestBody(r *http.Request) ([]byte, error) {
    return io.ReadAll(r.Body)
}

func GetRequestedRuleNumber(body []byte, ruleNumber *RuleNumber) error {
    reader := strings.NewReader(string(body))
    decoder := json.NewDecoder(reader)
    err := decoder.Decode(ruleNumber)
    if err != nil {
        return err
    }
    return nil
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
