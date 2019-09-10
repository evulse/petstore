package models

import (
	"encoding/json"
	"errors"
	"strings"
)

type PetStatus int

const (
	PetStatusUndefined PetStatus = iota
	PetStatusAvailable
	PetStatusPending
	PetStatusSold
)

func (p PetStatus) Validate() error {
	if p < 0 || p > 3 {
		return errors.New("invalid PetStatus")
	}
	return nil
}

func (p *PetStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*p = PetStatusUndefined
	case "available":
		*p = PetStatusAvailable
	case "pending":
		*p = PetStatusPending
	case "sold":
		*p = PetStatusSold
	}

	return nil
}

func (p PetStatus) MarshalJSON() ([]byte, error) {
	var s string
	switch p {
	default:
		s = ""
	case PetStatusAvailable:
		s = "available"
	case PetStatusPending:
		s = "pending"
	case PetStatusSold:
		s = "sold"
	}

	return json.Marshal(s)
}