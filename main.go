package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := "9IUX87ttXDDtemJ2WU49cuLBj"
	consumerSecret := "8oHF1IErmiQvSGgjxeJQqJiUQbLj0jbl1f8WLp5PJXS0EEovkW"
	accessToken := "1000696376232099840-9qxOVr2ytcl47695lblvhoM6BkQFMM"
	accessSecret := "ls6zZetlZImvgsp7QGHLkBUIFeS8T1zxbqoC8oczcfY2X"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query:      "#fandango",
		Count:      15,
		ResultType: "recent",
		Lang:       "en",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

	// Retweet
	for _, tweet := range searchResult.Statuses {
		tweet_id := tweet.ID
		client.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})

		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}
}
