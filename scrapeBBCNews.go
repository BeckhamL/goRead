package main

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

var cABBCNews = new(currentArticle)

func getMostFrequentWordsBBCNews(url string) [5]string {

	c := colly.NewCollector()

	var titleWords []string
	var paragraphWords [5]string

	// THIS IS FOR BBC-NEWS
	c.OnHTML("story-body", func(e *colly.HTMLElement) {

		paragraph := e.ChildText("story-body__introduction")
		title := e.ChildText("story-body__h1")

		if title != "" {

			counter := len(strings.Fields(title))
			titleWords = make([]string, counter)
			for i, words := range strings.Fields(title) {
				titleWords[i] = parseStringBBCNews(words)
			}
		}

		if paragraph != "" {

			parsedString := parseStringBBCNews(paragraph)
			m := make(map[string]int)
			var elements []keyValue
			m = WordCountBBCNews(parsedString)
			elements = sortMapBBCNews(m)

			for i := 0; i < 5; i++ {
				paragraphWords[i] = elements[i].Key
			}

			cABBCNews.frequentWords = paragraphWords
			cABBCNews.parsedString = parsedString
			words := len(strings.Fields(paragraph))
			cABBCNews.totalWords = float64(words)
			cABBCNews.titleWords = titleWords

			getSummaryBBCNews()
		}
	})

	c.Visit(url)

	return paragraphWords
}

// Function to remove all unecessary punctuation and character
func parseStringBBCNews(text string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	return newString
}

// Function to find words
func WordCountBBCNews(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

// Takes a map and returns a slice of keyValue sorted by highest frequency
func sortMapBBCNews(myMap map[string]int) []keyValue {

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
		if stringInSliceBBCNews(ss[i].Key, fillerWords) {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}

	return ss
}

// Checks if string is contained in array
func stringInSliceBBCNews(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Checks if there is at least one similar string in both arrays
func intersectFrequentBBCNews(arr1 []string, arr2 [5]string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func intersectTitleBBCNews(arr1 []string, arr2 []string) bool {

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
func getSummaryBBCNews() string {

	var response string
	var sentences []string
	var words []string
	sentences = strings.Split(cABBCNews.parsedString, ".")

	for i := 0; i < len(sentences); i++ {
		words = strings.Split(sentences[i], " ")
		if intersectFrequentBBCNews(words, cABBCNews.frequentWords) {
			response = response + "." + sentences[i]
		}
	}

	cABBCNews.summaryLength = float64(len(response))

	return response
}

// Get the reduction percentage
func getReductionPercentageBBCNews() float64 {

	return math.Round((cABBCNews.totalWords / cABBCNews.summaryLength) * 100)
}
