package main

import (
	//"log"
	"net/http"
)

func (s *ApiServer) HandlerGetInputRules(w http.ResponseWriter, r *http.Request) error {
	rules, err := GetRules(TableFilter, ChainInput)
	if err != nil {
		return err
	}

	WriteJSON(w, http.StatusOK, rules)
	return nil
}
