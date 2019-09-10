package models

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewPetWithNoName(t *testing.T) {
	pet := Pet{
		PhotoUrls: []string{"Text"},
	}
	err := pet.Save()
	if err == nil || err.Error() != "name can not be blank" {
		t.Errorf("Expected name can not be blank but got %s", err.Error())
	}
}

func TestNewPetWithNilPhotoUrls(t *testing.T) {
	pet := Pet{
		Name: "NewPet",
	}
	err := pet.Save()
	if err == nil || err.Error() != "at least one photoUrls must be defined" {
		t.Errorf("Expected at least one photoUrls must be defined got %s", err.Error())
	}
}

func TestNewPetWithEmptyPhotoUrls(t *testing.T) {
	pet := Pet{
		Name: "NewPet",
		PhotoUrls: []string{},
	}
	err := pet.Save()
	if err == nil || err.Error() != "at least one photoUrls must be defined" {
		t.Errorf("Expected at least one photoUrls must be defined got %s", err.Error())
	}
}

func TestNewPetWithValidInput(t *testing.T) {
	pet := Pet{
		Name: "NewPet",
		PhotoUrls: []string{"Text"},
	}
	err := pet.Save()
	if err != nil {
		t.Errorf("Expected no errors but got %s", err.Error())
	}
}

func TestPet_Marshal(t *testing.T) {
	jsonInput := []byte(`{
		  "id": 0,
		  "category": {
			"id": 0,
			"name": "category"
		  },
		  "name": "doggie",
		  "photoUrls": [
			"url"
		  ],
		  "tags": [
			{
			  "id": 0,
			  "name": "tag"
			}
		  ],
		  "status": "what",
		  "unknown": "ignore"
	}`)

   petObject := Pet{
   	Name: "doggie",
   	PhotoUrls: []string{"url"},
   	Tags: Tags{{Name: "tag"}},
   	Category: Category{Name: "category"},
   }
	var output Pet
	err := json.Unmarshal(jsonInput, &output)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}
	if !reflect.DeepEqual(petObject, output) {
		t.Errorf("Expected %v but got %v", petObject, output)
	}
}

func TestPet_UnMarshal(t *testing.T) {
	jsonInput := []byte(`{"id":0,"category":{"id":0,"name":"category"},"name":"doggie","photoUrls":["url"],"tags":[{"id":0,"name":"tag"}]}`)
	petObject := Pet{
		Name: "doggie",
		PhotoUrls: []string{"url"},
		Tags: Tags{{Name: "tag"}},
		Category: Category{Name: "category"},
	}
	jsonOutput, err := json.Marshal(petObject)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err.Error())
	}
	if !bytes.Equal(jsonInput, jsonOutput) {
		t.Errorf("Expected %s but got %s", jsonInput, jsonOutput)
	}
}
