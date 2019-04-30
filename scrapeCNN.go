package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

type keyValue struct {
	Key   string
	Value int
}

// Function that visits website and summarizes the text
func getSummary(url string) [10]string {

	c := colly.NewCollector()

	var titleWords []string
	var paragraphWords [10]string

	c.OnHTML(".l-container", func(e *colly.HTMLElement) {

		paragraph := e.ChildText(".zn-body__paragraph")
		title := e.ChildText(".pg-headline")

		if title != "" {

			counter := len(strings.Fields(title))
			titleWords = make([]string, counter)
			for i, words := range strings.Fields(title) {
				titleWords[i] = parseStringCNN(words)
			}
		}

		if paragraph != "" {
			parsedString := parseStringCNN(paragraph)
			m := make(map[string]int)
			var elements []keyValue
			m = WordCountCNN(parsedString)
			elements = sortMapCNN(m)

			for i := 0; i < 10; i++ {
				paragraphWords[i] = elements[i].Key
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	return paragraphWords
}

// Function to remove all unecessary punctuation and character
func parseStringCNN(text string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	return newString
}

// Function to find words
func WordCountCNN(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

// Takes a map and returns a slice of keyValue sorted by highest frequency
func sortMapCNN(myMap map[string]int) []keyValue {

	fillerWords := []string{"the", "to", "of", "a", "in", "and", "were", "they", "that", "have",
		"for", "been", "said", "but", "by", "is", "at", "how", "why", "many", "in", "on", "go", "of", "he", "was", "this", "or",
		"as", "if", "his", "also", "not", "it", "He", "She", "an", "able", "with", "I", "The"}

	var ss []keyValue

	for k, v := range myMap {
		ss = append(ss, keyValue{k, v})
	}

	// Sorting words by frequency
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i := 0; i < len(ss); i++ {
		if stringInSliceCNN(ss[i].Key, fillerWords) {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}

	return ss
}

// Checks if string is contained in array
func stringInSliceCNN(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
