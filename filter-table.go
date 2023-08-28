package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRules(table Table, chain Chain) ([]FilterTableRule, error) {
	var args []string = []string{ListRulesOption, string(chain)}
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
	var parsedRules []FilterTableRule = []FilterTableRule{}
	rules := strings.Split(string(rulesInBytes), "\n")
	for _, rule := range rules[1 : len(rules)-1] {
		var filterTableRule FilterTableRule
		splittedRule := strings.Split(rule, " ")
		splittedRule = FilterEmptyString(&splittedRule)
		for i := 0; i < len(splittedRule); i += 2 {
			option := splittedRule[i]
			value := splittedRule[i+1]
			switch option {
			case SourceAdressOption:
				filterTableRule.SourceAdress = &value
			case DestinationAdressOption:
				filterTableRule.DestinationAdress = &value
			case ProtocolOption:
				filterTableRule.Protocol = &value
			case SourcePortOption:
				sourcePort := new(int32)
				convertedSourcePort, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				*sourcePort = int32(convertedSourcePort)
				filterTableRule.SourcePort = sourcePort
			case DestinationPortOption:
				destinationPort := new(int32)
				convertedDestinationPort, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				*destinationPort = int32(convertedDestinationPort)
				filterTableRule.DestinationPort = destinationPort
			case TargetOption:
				filterTableRule.Target = &value
			}
		}
		parsedRules = append(parsedRules, filterTableRule)
	}
	return parsedRules, nil
}

func GetDefaultPolicy(table Table, chain Chain) (string, error) {
	var args []string = []string{ListRulesOption, string(chain), TableOption, string(table)}
	output, err := exec.Command(CmdIpTables, args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	rules := strings.Split(string(output), "\n")
	chainDefaultPolicy := strings.Split(rules[0], " ")
	return chainDefaultPolicy[2], nil
}

func AddRule(table Table, chain Chain, filterTableRule *FilterTableRule) error {
	if filterTableRule.RuleNumber == nil {
		filterTableRule.RuleNumber = new(int)
		*filterTableRule.RuleNumber = 1
	}

	args := []string{InsertOption, string(chain), strconv.Itoa(*filterTableRule.RuleNumber), TableOption, string(table)}

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
		args = append(args, SourcePortOption, strconv.Itoa(int(*filterTableRule.SourcePort)))
	}
	if filterTableRule.DestinationPort != nil {
		args = append(args, DestinationPortOption, strconv.Itoa(int(*filterTableRule.DestinationPort)))
	}
	if filterTableRule.Target != nil {
		args = append(args, TargetOption, *filterTableRule.Target)
	}

	if IsEmptyRule(filterTableRule, true) {
		return &ApiError{ErrorMsg: "empty rule"}
	}

	_, err := exec.Command(CmdIpTables, args...).CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func DeleteRule(table Table, chain Chain, ruleNumber *int) error {
	args := []string{DeleteOption, string(chain), strconv.Itoa(*ruleNumber)}

	_, err := exec.Command(CmdIpTables, args...).CombinedOutput()
	return err
}

func GetRuleByRuleNumber(table Table, chain Chain, ruleNumber *int) (*FilterTableRule, error) {
	rules, err := GetRules(table, chain)
	if err != nil {
		return nil, err
	}
	filterTableRule := new(FilterTableRule)
	for i, rule := range rules {
		if i+1 == *ruleNumber {
			filterTableRule = &rule
			break
		}
	}

	return filterTableRule, nil

}
