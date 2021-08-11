# go-twitter

A simple Go wrapper for Twitter API v2 ✨ 🍰 ✨

[![Test Status](https://github.com/sns-sdks/go-twitter/workflows/tests/badge.svg)](https://github.com/sns-sdks/go-twitter/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/sns-sdks/go-twitter/branch/main/graph/badge.svg)](https://codecov.io/gh/sns-sdks/go-twitter)
[![Go Report Card](https://goreportcard.com/badge/github.com/sns-sdks/go-twitter)](https://goreportcard.com/report/github.com/sns-sdks/go-twitter)

## Installation

```shell
# Go Modules
require github.com/sns-sdks/go-twitter
```

## Usage

### App with OAuth 2.0

Construct a new client with bearer token, then use the api.

```go
cli := twitter.NewBearerClient("your bearer token")
u, err := cli.Tweets.LookupByID("tweet id", twitter.TweetOpts{})
fmt.Println(u, err)
```

### Authentication

Some apis need `OAuth 1.0a User context`, You can use initial a client with user access token.

```go
app := twitter.NewAuthorizationAPP(
	twitter.AuthorizationAPP{
		ConsumerKey:       "consumer key",
		ConsumerSecret:    "consumer secret",
		AccessTokenKey:    "user access token key",
		AccessTokenSecret: "user access toke secret",
    })
cli := app.GetUserClient()
bu, err := cli.Users.GetBlocking(uid, twitter.UserBlockingOpts{})
fmt.Println(bu, err)
```

Or you can handle authentication contact with twitter. See the [`demo for authorize`](https://github.com/sns-sdks/go-twitter/tree/master/example/authorization.go)

## Features

Tweets:
- Lookup
- Timelines
- search Tweets
- Retweets
- Likes
- Hide replies

Users:
- Lookup
- Follows
- Blocks
- Mutes
