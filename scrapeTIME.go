package main

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

// Function to remove all unecessary punctuation and character
func parseStringTIME(text string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9'.]+")
	if err != nil {
		log.Fatal(err)
	}

	newString := reg.ReplaceAllString(text, " ")

	return newString
}

// Function to find words
func WordCountTIME(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

// Takes a map and returns a slice of keyValue sorted by highest frequency
func sortMapTIME(myMap map[string]int) []keyValue {

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
		if stringInSliceTIME(ss[i].Key, fillerWords) {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}

	return ss
}

// Checks if string is contained in array
func stringInSliceTIME(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
