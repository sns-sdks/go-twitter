# Examples

## User Auth By OAuth1

Many endpoints on the Twitter developer platform use the `OAuth 1.0a` method to act, or make API requests, on behalf of a Twitter account. For example, if you have a Twitter developer app, you can make API requests on behalf of any Twitter account as long as that user authenticates your app.

Please note: if you aren’t familiar with concepts such as HMAC-SHA1 and percent encoding, we recommend that you check out the "useful tools" section below that lists some API clients that greatly simplify the authentication process.

You can get more [information](https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code) for `OAuth 1.0a`

You need provide the `ConsumerKey`, `ConsumerSecret` configs.

Then run the demo

```shell
go run authorization_oauth1.go
```

## User Auth By OAuth2

`OAuth 2.0` is an industry-standard authorization protocol that allows for greater control over an application’s scope, and authorization flows across multiple devices. OAuth 2.0 allows you to pick specific fine-grained scopes which give you specific permissions on behalf of a user.

You can get more [information](https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code) for `OAuth 2.0`

You need provide the `ClientID` config.

Then run the demo

```shell
go run authorization_oauth2.go
```

## Use App-only Auth

Twitter offers applications the ability to issue authenticated requests on behalf of the application itself, as opposed to on behalf of a specific user. Twitter's implementation is based on the Client Credentials Grant flow of the OAuth 2 specification.

You can get more [information](https://developer.twitter.com/en/docs/authentication/oauth-2-0/application-only) for `App-only Auth`

You need provide the `App bearer access token` config.

Then run the demo

```shell
go run lookup_tweets.go
```
