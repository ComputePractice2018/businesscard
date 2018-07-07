package main

import (

	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/server"
)

func main() {
	//var name = flag.String("name", "Ilay", "имя для приветствия")
	//flag.Parse()

	http.HandleFunc("/api/businesscard/vcards", server.VcardsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
