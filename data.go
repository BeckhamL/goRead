package main

import (
	"context"
	"net/http"

	"github.com/barthr/newsapi"
)

type Article struct {
	Title    string
	Author   string
	Desc     string
	URL      string
	URLImage string
}

func topHeadlines() []Article {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"cnn", "time"},
	})

	if err != nil {
		panic(err)
	}

	var articles = []Article{}

	for _, s := range sources.Articles {

		a := Article{
			Title:    s.Title,
			Author:   s.Author,
			Desc:     s.Content,
			URL:      s.URL,
			URLImage: s.URLToImage,
		}

		articles = append(articles, a)

	}

	return articles
}

// func sports() []Article {

// 	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
// 	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
// 		category: []string{"sports"},
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
