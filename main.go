package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

// Option represents an option with an "id" and "details" field
type Fact struct {
	ID      int  `json:"id"`
	Description string  `json:"description"`
}
func main (){
 allFacts:= make([]Fact, 0)

 collector:= colly.NewCollector( colly.AllowedDomains("factretriever.com","www.factretriever.com"),)

 collector.OnHTML(".factsList li", func(element *colly.HTMLElement){
	factId, err:= strconv.Atoi(element.Attr("id"))
	if err!=nil {
		log.Println("could not find ID")
	}
   
	factDesc := element.Text

	fact := Fact{
		ID : factId,
		Description : factDesc,
	}

	allFacts = append(allFacts, fact)

 })

 collector.OnRequest( func(request *colly.Request) {
	fmt.Println("Visiting", request.URL.String())
 })

 collector.Visit("https://www.factretriever.com/rhino-facts")

 writeJson(allFacts)
}

func writeJson(data []Fact){
	file,err:=json.MarshalIndent(data, ""," ")
	if err!=nil{
		log.Println("unable to create json file")
		return
	}

	_ = ioutil.WriteFile("rhimofacts.json", file, 0644)

}

