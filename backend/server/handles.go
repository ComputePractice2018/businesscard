package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/data"
)

//VcardsHandler обрабатывет все запросы к /api/businesscard/vcards
func VcardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetVcards(w, r)
		return
	}
	if r.Method == "POST" {
		AddVcard(w, r)
		return
	}

	message := fmt.Sprintf("Method %s is not  allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Println(message)
}

//GetVcards обрабатывает запросы на получение списка контактов
func GetVcards(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.VcardList)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}
}

//AddVcard обрабатывает POST запрос
func AddVcard(w http.ResponseWriter, r *http.Request) {
	var Vcard data.Vcard
	err := json.NewDecoder(r.Body).Decode(&Vcard)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", Vcard)
	w.WriteHeader(http.StatusCreated)
}
