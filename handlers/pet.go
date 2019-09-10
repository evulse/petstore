package handlers

import (
	"encoding/json"
	"github.com/evulse/petstore/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PetHandler struct {}

func (handler *PetHandler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.Split(r.URL.Path, "/")[2]

	if len(param) == 0 {
		switch r.Method {
		case http.MethodGet:
			handler.ListHandler(w,r)
			return
		case http.MethodPost:
			handler.CreateHandler(w,r)
			return
		case http.MethodPut:
			handler.UpdateHandler(w,r)
			return
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			handler.DetailHandler(w,r)
			return
		}
	}
	w.WriteHeader(404);
}
func (handler *PetHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	record, err := models.LoadPets()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Unable to find any matching records"}`))
		return
	}

	jsonBody, _ := json.Marshal(record)
	w.Write(jsonBody)
}

func (handler *PetHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	log.Println(string(body))
	var pet models.Pet
	err = json.Unmarshal(body, &pet)
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	pet.Id = 0
	err = pet.Save()
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	jsonBody, _ := json.Marshal(pet)
	w.Write(jsonBody)
}

func (handler *PetHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	log.Println(string(body))
	var pet models.Pet
	err = json.Unmarshal(body, &pet)
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	err = pet.Save()
	if err != nil {
		w.WriteHeader(404)
		errorMessage := models.Error{Message: err.Error()}
		jsonBody, _ := json.Marshal(errorMessage)
		w.Write(jsonBody)
		return
	}
	jsonBody, _ := json.Marshal(pet)
	w.Write(jsonBody)
}

func (handler *PetHandler) DetailHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.Split(r.URL.Path, "/")[2]
	recordId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"invalid id format"}`))
		return
	}
	record, err := models.LoadPet(recordId)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Unable to find any matching records"}`))
		return
	}

	jsonBody, _ := json.Marshal(record)
	w.Write(jsonBody)
}

