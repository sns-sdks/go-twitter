# go-twitter

A simple Go wrapper for Twitter API v2 ✨ 🍰 ✨

[![Test Status](https://github.com/sns-sdks/go-twitter/workflows/tests/badge.svg)](https://github.com/sns-sdks/go-twitter/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/sns-sdks/go-twitter/branch/main/graph/badge.svg)](https://codecov.io/gh/sns-sdks/go-twitter)
[![Go Report Card](https://goreportcard.com/badge/github.com/sns-sdks/go-twitter)](https://goreportcard.com/report/github.com/sns-sdks/go-twitter)
[![Go Reference](https://pkg.go.dev/badge/github.com/sns-sdks/go-twitter.svg)](https://pkg.go.dev/github.com/sns-sdks/go-twitter)

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

[OAuth 2.0](https://developer.twitter.com/en/docs/twitter-api/oauth2) has released with beta.

So you can do `OAuth2.0` with twitter. More see the [`Example for OAuth2.0`](https://github.com/sns-sdks/go-twitter/tree/master/example/authorization_oauth2.go)


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

Or you can handle authentication contact with twitter. See the [`Example for OAuth 1.0a`](https://github.com/sns-sdks/go-twitter/tree/master/example/authorization_oauth1.go)

## Features

Tweets:
- Tweets lookup
- Manage Tweets
- Timelines
- search Tweets
- Tweet counts
- Retweets
- Likes
- Hide replies

Users:
- Users lookup
- Follows
- Blocks
- Mutes

Spaces:
- Lookup
- Search

Lists:
- List lookup
- Manage lists
- List Tweets lookup
- List members
- List follows
- Pinned Lists

Compliance:
- Batch compliance
