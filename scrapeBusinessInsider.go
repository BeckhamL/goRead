package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var cABusinessInsider = new(currentArticle)

func getMostFrequentWordsBusinessInsider(url string) [5]string {

	c := colly.NewCollector()

	//var titleWords []string
	var paragraphWords [5]string

	c.OnHTML(".collapse-container", func(e *colly.HTMLElement) {

		paragraph := e.ChildText(".summary-list")
		//title := e.ChildText(".story-body__h1")

		// if title != "" {

		// 	counter := len(strings.Fields(title))
		// 	titleWords = make([]string, counter)
		// 	for i, words := range strings.Fields(title) {
		// 		titleWords[i] = parseStringBBCNews(words)
		// 	}
		// }

		// if paragraph != "" {

		// 	parsedString := parseStringBBCNews(paragraph)
		// 	m := make(map[string]int)
		// 	var elements []keyValue
		// 	m = WordCountBBCNews(parsedString)
		// 	elements = sortMapBBCNews(m)

		// 	for i := 0; i < 5; i++ {
		// 		paragraphWords[i] = elements[i].Key
		// 	}

		// cABusinessInsider.frequentWords = paragraphWords
		// cABusinessInsider.parsedString = parsedString
		// words := len(strings.Fields(paragraph))
		// cABusinessInsider.totalWords = float64(words)
		// cABusinessInsider.titleWords = titleWords

		//getSummaryBuzzFeed()
		//}

		fmt.Println(paragraph)
	})

	c.Visit(url)

	return paragraphWords
}
