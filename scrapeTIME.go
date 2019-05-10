package main

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

var cATIMES = new(currentArticle)

func getMostFrequentWordsTIMES(url string) [5]string {

	c := colly.NewCollector()

	var titleWords []string
	var paragraphWords [5]string

	c.OnHTML(".l-container", func(e *colly.HTMLElement) {

		paragraph := e.ChildText("story-body__introduction")
		title := e.ChildText("story-body__h1")

		if title != "" {

			counter := len(strings.Fields(title))
			titleWords = make([]string, counter)
			for i, words := range strings.Fields(title) {
				titleWords[i] = parseStringTIMES(words)
			}
		}

		if paragraph != "" {

			parsedString := parseStringTIMES(paragraph)
			m := make(map[string]int)
			var elements []keyValue
			m = WordCountTIMES(parsedString)
			elements = sortMapTIMES(m)

			for i := 0; i < 5; i++ {
				paragraphWords[i] = elements[i].Key
			}

			cATIMES.frequentWords = paragraphWords
			cATIMES.parsedString = parsedString
			words := len(strings.Fields(paragraph))
			cATIMES.totalWords = float64(words)
			cATIMES.titleWords = titleWords

			getSummaryTIMES()
		}
	})

	c.Visit(url)

	return paragraphWords
}

// Function to remove all unecessary punctuation and character
func parseStringTIMES(text string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	return newString
}

// Function to find words
func WordCountTIMES(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

// Takes a map and returns a slice of keyValue sorted by highest frequency
func sortMapTIMES(myMap map[string]int) []keyValue {

	fillerWords := []string{"the", "to", "of", "a", "in", "and", "were", "they", "that", "have",
		"for", "been", "said", "but", "by", "is", "at", "how", "why", "many", "in", "on", "go", "of", "he", "was", "this", "or",
		"as", "if", "his", "also", "not", "it", "He", "She", "an", "able", "with", "I", "The", "will", "him", "be", "who", "has",
		"We", "are", "like", "than", "what", "your", "us", "had", "from", "would", "which", "now", "other", "we", "into", "could", "she",
		"her", "about", "you", "said."}

	var ss []keyValue

	for k, v := range myMap {
		ss = append(ss, keyValue{k, v})
	}

	// Sorting words by frequency
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i := 0; i < len(ss); i++ {
		if stringInSliceTIMES(ss[i].Key, fillerWords) {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}

	return ss
}

// Checks if string is contained in array
func stringInSliceTIMES(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Checks if there is at least one similar string in both arrays
func intersectFrequentTIMES(arr1 []string, arr2 [5]string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func intersectTitleTIMES(arr1 []string, arr2 []string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

// Extract summary from text
func getSummaryTIMES() string {

	var response string
	var sentences []string
	var words []string
	sentences = strings.Split(cATIMES.parsedString, ".")

	for i := 0; i < len(sentences); i++ {
		words = strings.Split(sentences[i], " ")
		if intersectFrequentCNN(words, cATIMES.frequentWords) {
			response = response + "." + sentences[i]
		}
	}

	cATIMES.summaryLength = float64(len(response))

	return response
}

// Get the reduction percentage
func getReductionPercentageTIMES() float64 {

	return math.Round((cATIMES.totalWords / cATIMES.summaryLength) * 100)
}
