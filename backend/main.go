package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/data"
	"github.com/ComputePractice2018/businesscard/backend/server"
)

func main() {
	//name := flag.String("name", "Ilay", "имя для приветствия")
	//flag.Parse()
	vcardList := data.NewVcardList()
	router := server.NewRouter(vcardList)
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*	http.HandleFunc("/api/businesscard/vcards", server.VcardsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}*/
