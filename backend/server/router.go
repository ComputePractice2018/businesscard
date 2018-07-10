package server

import (
	"github.com/ComputePractice2018/businesscard/backend/data"

	"github.com/gorilla/mux"
)

//NewRouter делает что-то
func NewRouter(vcardList data.Editable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/businesscard/vcards", GetVcards(vcardList)).Methods("GET")
	router.HandleFunc("/api/businesscard/vcards", AddVcard(vcardList)).Methods("POST")
	router.HandleFunc("/api/businesscard/vcards/{id}", EditVcard(vcardList)).Methods("PUT")
	router.HandleFunc("/api/businesscard/vcards/{id}", DeleteVcard(vcardList)).Methods("DELETE")
	return router
}
