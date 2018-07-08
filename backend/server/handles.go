package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

	"github.com/ComputePractice2018/businesscard/backend/data"
)

//GetVcards обрабатывает запросы на получение списка контактов
func GetVcards(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "plain/text; charset=utf-8") 
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetVcards())
	if err !=nil {
		message := fmt.Sprintf("Unable to encode data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}
}

//AddVcard обрабатывает POST запрос
func AddVcard(w http.ResponseWriter, r *http.Request) {
	var vcard data.Vcard
	err := json.NewDecoder(r.Body).Decode(&vcard)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}

	id := data.AddVcard(vcard)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}
//EditVcard обрабатывает PUT запрос 
func EditVcard(w http.ResponseWriter, r *http.Request) {
	var vcard data.Vcard
	err := json.NewDecoder(r.Body).Decode(&vcard)
	if err != nil {
		message := fmt.Sprintf("Unable to decode PUT data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}

	err = data.EditVcard(vcard, id )
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)

}

//DeleteVcard обрабатывает  DELETE запрос
func DeleteVcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}

	err = data.RemoveVcard(id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}