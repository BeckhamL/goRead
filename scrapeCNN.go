package main

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

type keyValue struct {
	Key   string
	Value int
}

type currentArticle struct {
	parsedString  string
	frequentWords [5]string
	titleWords    []string
	totalWords    float64
	summaryLength float64
}

var cACNN = new(currentArticle)

// Function that visits website and summarizes the text
func getMostFrequentWordsCNN(url string) [5]string {

	c := colly.NewCollector()

	var titleWords []string
	var paragraphWords [5]string

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
			elements = sortWordsCNN(m)

			for i := 0; i < 5; i++ {
				paragraphWords[i] = elements[i].Key
			}

			cACNN.frequentWords = paragraphWords
			cACNN.parsedString = parsedString
			words := len(strings.Fields(paragraph))
			cACNN.totalWords = float64(words)
			cACNN.titleWords = titleWords

			getSummaryCNN()
		}
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
func sortWordsCNN(myMap map[string]int) []keyValue {

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

// Checks if there is at least one similar string in both arrays
func intersectFrequentCNN(arr1 []string, arr2 [5]string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func intersectTitleCNN(arr1 []string, arr2 []string) bool {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if strings.ToLower(arr1[i]) == strings.ToLower(arr2[j]) {
				return true
			}
		}
	}
	return false
}

func countWordPriority(arr1 []string, arr2 []string) int {

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

func sortSentencesCNN(myMap map[string]int) []keyValue {

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
func getSummaryCNN() string {

	var response string
	var sentences []string
	var words []string

	sentenceWeight := make(map[string]int)

	sentences = strings.Split(cACNN.parsedString, ".")

	for i := 0; i < len(sentences); i++ {
		words = strings.Split(sentences[i], " ")
		sentenceWeight[sentences[i]] = countWordPriority(words, cACNN.titleWords)
	}

	summarizedSentences := sortSentencesCNN(sentenceWeight)
	response = summarizedSentences[0].Key

	if len(summarizedSentences) < 5 {
		cACNN.summaryLength = 0
		return response
	} else {
		for i := 1; i < 5; i++ {
			response = response + ". " + summarizedSentences[i].Key
		}
		cACNN.summaryLength = float64(len(response))
	}

	response = response + "."

	return response
}

// Get the reduction percentage
func getReductionPercentageCNN() float64 {

	return math.Round((cACNN.totalWords / cACNN.summaryLength) * 100)
}
