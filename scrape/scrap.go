//scrap and get the currency current value and make it update every hour
package scrape

import (
	"fmt"
	

	"github.com/gocolly/colly"
)

//Get the info about currencies by passing the required arguments
func GetInfo(from, to, amount string) string{
	url := fmt.Sprintf("https://www.xe.com/currencyconverter/convert/?Amount=%s&From=%s&To=%s", amount, from, to)
	
	
	var info string

	c := colly.NewCollector(
		colly.AllowedDomains("www.xe.com"),
	)

	c.OnHTML("p.result__BigRate-sc-1bsijpp-1", func(e *colly.HTMLElement) {
		info = e.Text
	})

	c.Visit(url)
	return info
}


