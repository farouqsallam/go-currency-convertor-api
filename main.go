package main

import (
	"convertor/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", router.Getcurrencies)

	http.HandleFunc("/currency", router.Currency)
	

	//Server connecting
	fmt.Println("Listening 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Error connecting to the server")
	}
}

