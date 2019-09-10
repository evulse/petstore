package models

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestPetStatus_Validate(t *testing.T) {
	status := PetStatus(10)
	err := status.Validate()
	if err == nil || err.Error() != "invalid PetStatus" {
		t.Errorf("Expected name invalid PetStatus but got %s", err.Error())
	}

	status = PetStatusSold
	err = status.Validate()
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}

	status = 0
	err = status.Validate()
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}
}

func TestPetStatus_MarshalJSON(t *testing.T) {
	status := [...]PetStatus{PetStatusSold, PetStatusPending, PetStatusAvailable, PetStatusUndefined, 10}
	json, err := json.Marshal(status)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}
	if !bytes.Equal(json,[]byte(`["sold","pending","available","",""]`)) {
		t.Errorf("Expected err to be nil but got %s", json)
	}
}

func TestPetStatus_UnmarshalJSON(t *testing.T) {
	statusObjects := [...]PetStatus{PetStatusSold, PetStatusPending, PetStatusAvailable, PetStatusUndefined, PetStatusUndefined, PetStatusUndefined}
	status := []byte(`["sold","pending","available","","foo",null]`)
	var output [6]PetStatus
	err := json.Unmarshal(status, &output)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}
	if !reflect.DeepEqual(statusObjects, output) {
		t.Errorf("Expected %v but got %v", statusObjects, output)
	}
}