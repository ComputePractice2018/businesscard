package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/businesscard/backend/data"
	"github.com/ComputePractice2018/businesscard/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	connection := flag.String("connection", "businesscard:SuperSecretPassword@tcp(db:3306)/addressbook", "mysql connection string")
	flag.Parse()

	vcardList, err := data.NewDBVcardList(*connection)
	defer vcardList.Close()
	//vcardList := data.NewVcardList()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(vcardList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
