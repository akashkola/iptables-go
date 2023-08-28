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

func ValidateRuleNumber(table Table, chain Chain, ruleNumber *int, deleteRequest bool) error {
    if ruleNumber == nil {
        ruleNumber = new(int)
        *ruleNumber = 1
    }

    rules, err := GetRules(table, chain)
	if err != nil {
		return err
	}

	rulesLength := len(rules)
    
    if !deleteRequest {
        rulesLength += 1
    }

    if *ruleNumber > rulesLength {
        return &ApiError{ErrorMsg: "rule number is too big"}
	}

	if *ruleNumber <= 0 {
		return &ApiError{ErrorMsg: "invalid rule number"}
	}

	return nil
}

func GetUrlParameter(r *http.Request, parameter string) string {
    paramValue := chi.URLParam(r, parameter)
    return paramValue
}

func CopyRule(srcRule *FilterTableRule, dstRule *FilterTableRule) {
    if srcRule == nil {
        return
    }

    if srcRule.SourceAdress != nil {
        dstRule.SourceAdress = srcRule.SourceAdress
    }
    if srcRule.DestinationAdress != nil {
        dstRule.DestinationAdress = srcRule.DestinationAdress
    }
    if srcRule.Protocol != nil {
        dstRule.Protocol = srcRule.Protocol
    }
    if srcRule.SourcePort != nil {
        dstRule.SourcePort = srcRule.SourcePort
    }
    if srcRule.DestinationPort != nil {
        dstRule.DestinationPort = srcRule.DestinationPort
    }
    if srcRule.Target != nil {
        dstRule.Target = srcRule.Target
    }
}

func IsEmptyRule(rule *FilterTableRule, ignoreTarget bool) bool { 
    if rule.SourceAdress != nil {
        return false
    }
    if rule.DestinationAdress != nil {
        return false
    }
    if rule.Protocol != nil {
        return false
    }
    if rule.SourcePort != nil {
        return false
    }
    if rule.DestinationPort != nil {
        return false
    }
    if !ignoreTarget {
        if rule.Target != nil {
            return false
        }
    }

    return true
}

