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

func topHeadlinesCNN() []Article {

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
		timeString = timeString[0 : len(timeString)-18]

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

func topHeadlinesBBCNews() []Article {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"bbc-news"},
	})

	if err != nil {
		panic(err)
	}

	var articles = []Article{}

	for _, s := range sources.Articles {

		t := s.PublishedAt
		t.Format("Monday Jan 2, 3:04pm")
		timeString := t.String()
		timeString = timeString[0 : len(timeString)-18]

		a := Article{
			Title:     s.Title,
			Author:    s.Author,
			Keywords:  getMostFrequentWordsBBCNews(s.URL),
			URL:       s.URL,
			URLImage:  s.URLToImage,
			Website:   s.Source.Name,
			Date:      timeString,
			Summary:   getSummaryBBCNews(),
			Reduction: getReductionPercentageBBCNews(),
		}

		articles = append(articles, a)

	}

	return articles
}

func topHeadlinesBusinessInsider() []Article {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"business-insider"},
	})

	if err != nil {
		panic(err)
	}

	var articles = []Article{}

	for _, s := range sources.Articles {

		t := s.PublishedAt
		t.Format("Monday Jan 2, 3:04pm")
		timeString := t.String()
		timeString = timeString[0 : len(timeString)-18]

		a := Article{
			Title:     s.Title,
			Author:    s.Author,
			Keywords:  getMostFrequentWordsBusinessInsider(s.URL),
			URL:       s.URL,
			URLImage:  s.URLToImage,
			Website:   s.Source.Name,
			Date:      timeString,
			Summary:   getSummaryBusinessInsider(),
			Reduction: getReductionPercentageBusinessInsider(),
		}

		articles = append(articles, a)

	}

	return articles
}

func topHeadlinesCBC() []Article {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))
	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"cbc-news"},
	})

	if err != nil {
		panic(err)
	}

	var articles = []Article{}

	for _, s := range sources.Articles {

		t := s.PublishedAt
		t.Format("Monday Jan 2, 3:04pm")
		timeString := t.String()
		timeString = timeString[0 : len(timeString)-18]

		a := Article{
			Title:     s.Title,
			Author:    s.Author,
			Keywords:  getMostFrequentWordsCBC(s.URL),
			URL:       s.URL,
			URLImage:  s.URLToImage,
			Website:   s.Source.Name,
			Date:      timeString,
			Summary:   getSummaryCBC(),
			Reduction: getReductionPercentageCBC(),
		}

		articles = append(articles, a)

	}

	return articles
}
