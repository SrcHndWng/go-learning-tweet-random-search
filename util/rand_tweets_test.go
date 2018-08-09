package util

import (
	"fmt"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
)

func TestSelectTweet(t *testing.T) {
	t1 := randTweet{Tweet: twitter.Tweet{FavoriteCount: 1, RetweetCount: 2, Text: "aaa"}}
	t2 := randTweet{Tweet: twitter.Tweet{FavoriteCount: 3, RetweetCount: 4, Text: "bbb"}}
	t3 := randTweet{Tweet: twitter.Tweet{FavoriteCount: 5, RetweetCount: 6, Text: "ccc"}}
	var rts randTweets
	rts = []randTweet{t1, t2, t3}
	result := rts.SelectTweet()
	fmt.Printf("randTweets = %v\n", rts)
	fmt.Printf("result text = %v\n", result.Tweet.Text)
}
