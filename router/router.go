package router

import (
	"convertor/scrape"
	"net/http"
	"os"
	"fmt"
)


func Getcurrencies(w http.ResponseWriter, r *http.Request){
	data, err := os.ReadFile("../util/Currences.txt")
	if err != nil {
		fmt.Println("File not found")
	}

	w.Write(data)

}


func Currency(w http.ResponseWriter, r *http.Request){
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	info := scrape.GetInfo(from, to, amount)
	
	w.Write([]byte(info))
}

