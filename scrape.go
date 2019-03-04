package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

// better algorithm, get the title, break the paragraph into sentences, if the sentence contains a word in the title, add it
func main() {

	c := colly.NewCollector()

	c.OnHTML(".l-container", func(e *colly.HTMLElement) {

		paragraph := e.ChildText(".zn-body__paragraph")
		title := e.ChildText(".pg-headline")

		if title != "" {

			counter := len(strings.Fields(title))
			titleWords := make([]string, counter)
			for i, words := range strings.Fields(title) {
				titleWords[i] = words
			}
		}

		if paragraph != "" {
			parseString(paragraph)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.cnn.com/2019/03/03/us/tornadoes-alabama-georgia-wxc/index.html")
}

func parseString(text string) {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	WordCount(newString)
}

// Function to find words
func WordCount(s string) {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	sortMap(m)
}

func sortMap(myMap map[string]int) map[string]int {

	fillerWords := []string{"the", "to", "of", "a", "in", "and", "were", "they", "that", "have",
		"for", "been", "said", "but", "by", "is", "at", "how", "why", "many", "in", "on", "go", "of"}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range myMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	newMap := make(map[string]int)

	for _, kv := range ss {

		if stringInSlice(kv.Key, fillerWords) {
			delete(myMap, kv.Key)
		}

		newMap[kv.Key] = kv.Value
	}

	for k, v := range newMap {
		fmt.Println(k, v)
	}

	return myMap
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
