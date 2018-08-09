package util

import (
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

// CreateHTTPClient creates a http client.
func CreateHTTPClient() *http.Client {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	return config.Client(oauth1.NoContext, token)
}
