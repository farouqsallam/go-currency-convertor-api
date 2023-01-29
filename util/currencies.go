package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

//Scrap the available currencies use it once to make csv file
func main(){
	var data  []string
	url := "https://www.worlddata.info/currencies/"
	c := colly.NewCollector(
		colly.AllowedDomains("www.worlddata.info"),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.OnHTML("table.std100>tbody>tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, h *colly.HTMLElement) {
			if i == 0 {
				data = append(data, h.Text)
			}
		})
	})
	c.Visit(url)
	
	myFile, err := os.Create("Currences.txt")
	defer myFile.Close()
	if err != nil {
		log.Fatal("Cannot create new file currencies.txt")
	}

	
	for _, item := range data {
		myFile.WriteString(fmt.Sprintf("%s\n", item))
	}
}

