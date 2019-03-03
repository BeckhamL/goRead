package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML(".l-container", func(e *colly.HTMLElement) {
		text := e.ChildText(".zn-body__paragraph")
		parseString(text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.cnn.com/2019/03/03/politics/rand-paul-trump-national-emergency-declaration/index.html")
}

func parseString(text string) {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	WordCount(newString)
	//fmt.Println(newString)
}

// Function to found words
func WordCount(s string) {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}

	sortMap(m)
}

func sortMap(myMap map[string]int) map[string]int {

	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println(myMap)

	return myMap
}
