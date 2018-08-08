package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

type twitterClient struct {
	client *twitter.Client
}

// NewTwitterClient creates twitterClient struct
func NewTwitterClient(httpClient *http.Client) *twitterClient {
	tc := new(twitterClient)
	tc.client = twitter.NewClient(httpClient)
	return tc
}

func (tc twitterClient) receiveTweet(tweet *twitter.Tweet) {
	text := tweet.Text
	log.Printf("receive tweet = %v\n", tweet)
	searchKeyword := strings.Split(text, " ")[1]
	search, _, err := tc.client.Search.Tweets(&twitter.SearchTweetParams{
		Query: searchKeyword,
	})
	if err != nil {
		log.Fatal(err)
	}
	var rndTweets randTweets
	for _, t := range search.Statuses {
		rndTweets = append(rndTweets, randTweet{Tweet: t})
	}
	selected := rndTweets.SelectTweet().Tweet
	url := fmt.Sprintf("https://twitter.com/%s/status/%s\n", selected.User.ScreenName, selected.IDStr)
	log.Printf("selected url = %s\n", url)
	_, _, err = tc.client.Statuses.Update(url, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (tc twitterClient) filterTweets() (stream *twitter.Stream, err error) {
	// FILTER
	keyWord := fmt.Sprintf("@%s", os.Getenv("TWITTER_ACCOUNT"))
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{keyWord},
		StallWarnings: twitter.Bool(true),
	}
	stream, err = tc.client.Streams.Filter(filterParams)
	return
}
