package main

import (
	"context"
	"net/http"
	"time"

	"github.com/barthr/newsapi"
)

type Article struct {
	Title    string
	Author   string
	Keywords [10]string
	URL      string
	URLImage string
	Website  string
	Date     time.Time
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

		a := Article{
			Title:    s.Title,
			Author:   s.Author,
			Keywords: getSummary(s.URL),
			URL:      s.URL,
			URLImage: s.URLToImage,
			Website:  s.Source.Name,
			Date:     s.PublishedAt,
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
