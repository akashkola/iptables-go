package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRules(table Table, chain Chain) ([]FilterTableRule, error) {
    var args []string = []string{ "-S", string(chain) }
    rulesInBytes, err := exec.Command(CmdIpTables, args...).CombinedOutput()
    if err != nil {
        return nil, err
    }
    rules, err := parseRules(chain, rulesInBytes)
    if err != nil {
        return nil, err
    }
    return rules, nil
}

func parseRules(chain Chain, rulesInBytes []byte) ([]FilterTableRule, error) { 
    var parsedRules []FilterTableRule = make([]FilterTableRule, 0)
    rules := strings.Split(string(rulesInBytes), "\n")
    for _, rule := range rules[1:len(rules) - 1] {
        var filterTableRule FilterTableRule
        splittedRule := strings.Split(rule, " ")
        splittedRule = FilterEmptyString(&splittedRule)
        for i := 0; i < len(splittedRule); i += 2 {
            option := splittedRule[i]
            value := splittedRule[i + 1]
            switch option {
                case SourceAdressOption:
                    filterTableRule.SourceAdress = &value
                case DestinationAdressOption:
                    filterTableRule.DestinationAdress = &value
                case ProtocolOption:
                    filterTableRule.Protocol = &value
                case SourcePortOption:
                    filterTableRule.SourcePort = &value
                case DestinationPortOption:
                    filterTableRule.DestinationPort = &value
                case TargetOption:
                    filterTableRule.Target = &value
            }
        }
        parsedRules = append(parsedRules, filterTableRule)
    }
    return parsedRules, nil
}

func GetDefaultPolicy(table Table, chain Chain) (string, error) {
    var args []string = []string { "-S", string(chain), "-t", string(table) } 
    output, err := exec.Command(CmdIpTables, args...).CombinedOutput()
    if err != nil {
        return "", err
    }
    
    rules := strings.Split(string(output), "\n")
    chainDefaultPolicy := strings.Split(rules[0], " ")
    return chainDefaultPolicy[2], nil
}

func AddRule(table Table, chain Chain, filterTableRule *FilterTableRule, numRule *int) error {
    if numRule == nil {
        numRule = new(int)
        *numRule = 1
    }    

    args := []string{"-I", string(chain), strconv.Itoa(*numRule), "-t", string(table)}

    if filterTableRule.SourceAdress != nil {
        args = append(args, SourceAdressOption, *filterTableRule.SourceAdress)
    }
    if filterTableRule.DestinationAdress != nil {
        args = append(args, DestinationAdressOption, *filterTableRule.DestinationAdress)
    }
    if filterTableRule.Protocol != nil {
        args = append(args, ProtocolOption, *filterTableRule.Protocol)
    }
    if filterTableRule.SourcePort != nil {
        args = append(args, SourcePortOption, *filterTableRule.SourcePort)
    }
    if filterTableRule.DestinationPort != nil {
        args = append(args, DestinationPortOption, *filterTableRule.DestinationPort)
    }
    if filterTableRule.Target != nil {
        args = append(args, TargetOption, *filterTableRule.Target)
    }
    _, err := exec.Command(CmdIpTables, args...).CombinedOutput()
    if err != nil {
        return err
    }
    return nil
}

func ValidateRuleNumber(table Table, chain Chain, ruleNumber int) error {
    rules, err := GetRules(table, chain)
    if err != nil {
        return err
    }

    rulesLength := len(rules)
    if ruleNumber > rulesLength + 1 {
        return &ApiError{error: "rule number is too big"}
    } 
    if ruleNumber <= 0 {
        return &ApiError{error: "invalid rule number"}
    }

    return nil

} 
