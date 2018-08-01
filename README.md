# go-learning-tweet-random-search

## About This

Twitterのアカウントに「@Account Keyword」形式でメッセージを投げると

キーワードで検索した結果からランダムにツイートを選び、返却してくれるBotです。

Golangの学習用に作成しました。

## Usage

環境変数に以下の値を設定し、バイナリを実行します。
停止はCtrl + Cで行います。

```
TWITTER_CONSUMER_KEY=Twitter APIのCunsumerKey
TWITTER_CONSUMER_SECRET=Twitter APIのConsumerSecret
TWITTER_ACCESS_TOKEN=Twitter APIのAccessToken
TWITTER_ACCESS_TOKEN_SECRET=Twitter APIのAccessTokenSecret
TWITTER_ACCOUNT=Twitter Account名
```
