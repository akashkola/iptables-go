package main

type Table string

type Chain string

type FilterTableRule struct {
    SourceAdress *string `json:"source-address,omitempty"`
    DestinationAdress *string `json:"destination-address,omitempty"`
    Protocol *string `json:"protocol,omitempty"`
    SourcePort *string `json:"source-port,omitempty"`
    DestinationPort *string `json:"destination-port,omitempty"`
    Target *string `json:"target,omitempty"`
}

// TODO: Add this rule number to FilterTableRule Type
type RuleNumber struct {
    RuleNum *int `json:"rule-number,omitempty"`
}


