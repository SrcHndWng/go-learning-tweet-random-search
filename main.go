package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SrcHndWng/go-learning-tweet-random-search/util"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	httpClient := util.CreateHTTPClient()
	twitterClient := util.NewTwitterClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = twitterClient.ReceiveTweet

	log.Println("Starting Stream...")

	stream, err := twitterClient.FilterTweets()
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
