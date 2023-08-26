package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type FilterTableRule struct {
    SourceAdress *string `json:"source-address,omitempty"`
    DestinationAdress *string `json:"destination-address,omitempty"`
    Protocol *string `json:"protocol,omitempty"`
    SourcePort *string `json:"source-port,omitempty"`
    DestinationPort *string `json:"destination-port,omitempty"`
    Target *string `json:"target,omitempty"`
}

type Table string

const (
    TableFilter Table = "filter"
)

type Chain string

const (
    ChainInput Chain = "INPUT"
)

const (
    cmdIpTables string = "iptables"
)


const (
    sourceAdressOption string= "-s"
    destinationAdressOption string = "-d"
    protocolOption string = "-p"
    sourcePortOption string = "--sport"
    destinationPortOption string = "--dport"
    targetOption string = "-j"
)

type RuleNumber struct {
    RuleNum *int `json:"rule-number,omitempty"`
}


func filterEmptyString(str *[]string) []string {
    filteredStringSlice := make([]string, 0)
    for _, strValue :=  range *str {
        if strValue == "" {
            continue
        }
        filteredStringSlice = append(filteredStringSlice, strValue)
    }
    return filteredStringSlice
}

func getRules(table Table, chain Chain) ([]FilterTableRule, error) {
    var args []string = []string{ "-S", string(chain) }
    rulesInBytes, err := exec.Command(cmdIpTables, args...).CombinedOutput()
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
        splittedRule = filterEmptyString(&splittedRule)
        for i := 0; i < len(splittedRule); i += 2 {
            option := splittedRule[i]
            value := splittedRule[i + 1]
            switch option {
                case sourceAdressOption:
                    filterTableRule.SourceAdress = &value
                case destinationAdressOption:
                    filterTableRule.DestinationAdress = &value
                case protocolOption:
                    filterTableRule.Protocol = &value
                case sourcePortOption:
                    filterTableRule.SourcePort = &value
                case destinationPortOption:
                    filterTableRule.DestinationPort = &value
                case targetOption:
                    filterTableRule.Target = &value
                default:
                    fmt.Println("default case : ", option, ": ", value)
            }
        }
        parsedRules = append(parsedRules, filterTableRule)
    }
    return parsedRules, nil
}

func getDefaultPolicy(table Table, chain Chain) (string, error) {
    var args []string = []string { "-S", string(chain), "-t", string(table) } 
    output, err := exec.Command(cmdIpTables, args...).CombinedOutput()
    if err != nil {
        return "", err
    }
    
    rules := strings.Split(string(output), "\n")
    chainDefaultPolicy := strings.Split(rules[0], " ")
    return chainDefaultPolicy[2], nil
}

func addRule(table Table, chain Chain, filterTableRule *FilterTableRule, numRule *int) error {
    if numRule == nil {
        numRule = new(int)
        *numRule = 1
    }
    args := []string{"-I", string(chain), strconv.Itoa(*numRule), "-t", string(table)}

    if filterTableRule.SourceAdress != nil {
        args = append(args, sourceAdressOption, *filterTableRule.SourceAdress)
    }
    if filterTableRule.DestinationAdress != nil {
        args = append(args, destinationAdressOption, *filterTableRule.DestinationAdress)
    }
    if filterTableRule.Protocol != nil {
        args = append(args, protocolOption, *filterTableRule.Protocol)
    }
    if filterTableRule.SourcePort != nil {
        args = append(args, sourcePortOption, *filterTableRule.SourcePort)
    }
    if filterTableRule.DestinationPort != nil {
        args = append(args, destinationPortOption, *filterTableRule.DestinationPort)
    }
    if filterTableRule.Target != nil {
        args = append(args, targetOption, *filterTableRule.Target)
    }
    _, err := exec.Command(cmdIpTables, args...).CombinedOutput()
    if err != nil {
        return err
    }
    return nil
}
