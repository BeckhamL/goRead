package main

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

var cACBC = new(currentArticle)

// Function that visits website and summarizes the text
func getMostFrequentWordsCBC(url string) [5]string {

	c := colly.NewCollector()

	var titleWords []string
	var paragraphWords [5]string

	c.OnHTML(".detailBodyContent", func(e *colly.HTMLElement) {

		paragraph := e.ChildText(".story")
		title := e.ChildText(".detailHeadline")

		if title != "" {

			counter := len(strings.Fields(title))
			titleWords = make([]string, counter)
			for i, words := range strings.Fields(title) {
				titleWords[i] = parseStringCBC(words)
			}
		}

		if paragraph != "" {
			parsedString := parseStringCBC(paragraph)
			m := make(map[string]int)
			var elements []keyValue
			m = WordCountCBC(parsedString)
			elements = sortWordsCBC(m)

			for i := 0; i < 5; i++ {
				paragraphWords[i] = elements[i].Key
			}

			cACBC.frequentWords = paragraphWords
			cACBC.parsedString = parsedString
			words := len(strings.Fields(paragraph))
			cACBC.totalWords = float64(words)
			cACBC.titleWords = titleWords

			getSummaryCBC()
		}
	})

	c.Visit(url)

	return paragraphWords
}

// Function to remove all unecessary punctuation and character
func parseStringCBC(text string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	return newString
}

// Function to find words
func WordCountCBC(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

// Takes a map and returns a slice of keyValue sorted by highest frequency
func sortWordsCBC(myMap map[string]int) []keyValue {

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
		if stringInSliceCBC(ss[i].Key, fillerWords) {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}

	return ss
}

// Checks if string is contained in array
func stringInSliceCBC(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Checks if there is at least one similar string in both arrays
func intersectFrequentCBC(arr1 []string, arr2 [5]string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func intersectTitleCBC(arr1 []string, arr2 []string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func countWordPriorityCBC(arr1 []string, arr2 []string) int {

	counter := 0

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				counter++
			}
		}
	}

	return counter
}

func sortSentencesCBC(myMap map[string]int) []keyValue {

	var kv []keyValue

	for k, v := range myMap {
		kv = append(kv, keyValue{k, v})
	}

	// Sorting words by frequency
	sort.Slice(kv, func(i, j int) bool {
		return kv[i].Value > kv[j].Value
	})

	return kv
}

// Extract summary from text
func getSummaryCBC() string {

	var response string
	var sentences []string
	var words []string

	sentenceWeight := make(map[string]int)

	sentences = strings.Split(cACBC.parsedString, ".")

	for i := 0; i < len(sentences); i++ {
		words = strings.Split(sentences[i], " ")
		sentenceWeight[sentences[i]] = countWordPriorityCBC(words, cACBC.titleWords)
	}

	summarizedSentences := sortSentencesCBC(sentenceWeight)
	response = summarizedSentences[0].Key

	if len(summarizedSentences) < 5 {
		cACBC.summaryLength = 0
		return response
	} else {
		for i := 1; i < 5; i++ {
			response = response + ". " + summarizedSentences[i].Key
		}
		cACBC.summaryLength = float64(len(response))
	}

	response = response + "."

	return response
}

// Get the reduction percentage
func getReductionPercentageCBC() float64 {

	return math.Round((cACBC.totalWords / cACBC.summaryLength) * 100)
}
