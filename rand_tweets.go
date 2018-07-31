package main

import (
	"math/rand"
	"sort"

	"github.com/dghubble/go-twitter/twitter"
)

type randTweet struct {
	Tweet twitter.Tweet
	Rand  int
}

type randTweets []randTweet

func (rts randTweets) Len() int           { return len(rts) }
func (rts randTweets) Less(i, j int) bool { return rts[i].Rand > rts[j].Rand }
func (rts randTweets) Swap(i, j int)      { rts[i], rts[j] = rts[j], rts[i] }

func (rts randTweets) SelectTweet() randTweet {
	for i, t := range rts {
		rts[i].Rand = t.Tweet.RetweetCount * t.Tweet.FavoriteCount * rand.Intn(10)
	}
	sort.Sort(rts)
	return rts[0]
}
