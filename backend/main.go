package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/server"
	"github.com/gorilla/mux"
)

func main() {
	//name := flag.String("name", "Ilay", "имя для приветствия")
	//flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/businesscard/vcards", server.GetVcards).Methods("GET")
	router.HandleFunc("/api/businesscard/vcards", serverAddVcard).Methods("POST")
	router.HandleFunc("/api/businesscard/vcards/{id}", serverEditVcard).Methods("PUT")
	router.HandleFunc("/api/businesscard/vcards/{id}", serverDeleteVcard).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
