package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/data"
	"github.com/ComputePractice2018/businesscard/backend/server"
	"github.com/gorilla/mux"
)

func main() {
	//name := flag.String("name", "Ilay", "имя для приветствия")
	//flag.Parse()
	vcardList := data.NewVcardList()
	router := mux.NewRouter()
	router.HandleFunc("/api/businesscard/vcards", server.GetVcards(vcardList)).Methods("GET")
	router.HandleFunc("/api/businesscard/vcards", server.AddVcard(vcardList)).Methods("POST")
	router.HandleFunc("/api/businesscard/vcards/{id}", server.EditVcard(vcardList)).Methods("PUT")
	router.HandleFunc("/api/businesscard/vcards/{id}", server.DeleteVcard(vcardList)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*	http.HandleFunc("/api/businesscard/vcards", server.VcardsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}*/
