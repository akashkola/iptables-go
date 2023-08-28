package main

type Table string

type Chain string

type FilterTableRule struct {
	SourceAdress      *string `json:"sourceAddress,omitempty"`
	DestinationAdress *string `json:"destinationAddress,omitempty"`
	Protocol          *string `json:"protocol,omitempty"`
	SourcePort        *int32  `json:"sourcePort,omitempty"`
	DestinationPort   *int32  `json:"destinationPort,omitempty"`
	Target            *string `json:"target,omitempty"`
	RuleNumber        *int    `json:"ruleNumber,omitempty"`
}
