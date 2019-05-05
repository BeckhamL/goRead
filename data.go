package main

import (
	"context"
	"net/http"

	"github.com/barthr/newsapi"
)

const (
	dateCAN = "January 2, 2006"
)

type Article struct {
	Title     string
	Author    string
	Keywords  [5]string
	URL       string
	URLImage  string
	Website   string
	Date      string
	Summary   string
	Reduction float64
}

func topHeadlines() []Article {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"cnn"},
	})

	if err != nil {
		panic(err)
	}

	var articles = []Article{}

	for _, s := range sources.Articles {

		t := s.PublishedAt
		t.Format("Monday Jan 2, 3:04pm")
		timeString := t.String()
		timeString = timeString[0 : len(timeString)-10]

		a := Article{
			Title:     s.Title,
			Author:    s.Author,
			Keywords:  getMostFrequentWordsCNN(s.URL),
			URL:       s.URL,
			URLImage:  s.URLToImage,
			Website:   s.Source.Name,
			Date:      timeString,
			Summary:   getSummaryCNN(),
			Reduction: getReductionPercentageCNN(),
		}

		articles = append(articles, a)

	}

	return articles
}

// func sports() []Article {

// 	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
// 	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
// 		Category: "sports",
// 		Country:  "us",
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	var articles = []Article{}

// 	for _, s := range sources.Articles {

// 		a := Article{
// 			Title:    s.Title,
// 			Author:   s.Author,
// 			Desc:     s.Content,
// 			URL:      s.URL,
// 			URLImage: s.URLToImage,
// 		}

// 		articles = append(articles, a)

// 	}

// 	return articles
// }
