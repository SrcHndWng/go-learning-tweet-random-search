package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var httpClient *http.Client
var client *twitter.Client

func createHTTPClient() *http.Client {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	return config.Client(oauth1.NoContext, token)
}

func filterTweets(client *twitter.Client) (stream *twitter.Stream, err error) {
	// FILTER
	keyWord := fmt.Sprintf("@%s", os.Getenv("TWITTER_ACCOUNT"))
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{keyWord},
		StallWarnings: twitter.Bool(true),
	}
	stream, err = client.Streams.Filter(filterParams)
	return
}

func receiveTweet(tweet *twitter.Tweet) {
	text := tweet.Text
	log.Printf("receive tweet = %v\n", tweet)
	searchKeyword := strings.Split(text, " ")[1]
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
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
	_, _, err = client.Statuses.Update(url, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	httpClient = createHTTPClient()
	client = twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = receiveTweet

	log.Println("Starting Stream...")

	stream, err := filterTweets(client)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	log.Println("Stopping Stream...")
	stream.Stop()
}
