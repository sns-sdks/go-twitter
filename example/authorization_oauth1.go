package main

import (
	"fmt"
	"github.com/sns-sdks/go-twitter/twitter"
)

const (
	ConsumerKey       = "Your app consumer key"
	ConsumerSecret    = "Your app consumer secret"
	OAuth1CallbackURL = "https://localhost/" // Your redirect uri
)

func main() {
	app := twitter.AuthorizationAPP{
		ConsumerKey:    ConsumerKey,
		ConsumerSecret: ConsumerSecret,
		CallbackURL:    OAuth1CallbackURL,
	}
	authUrl, _, err := app.GetAuthorizationURL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Click the authorization url: " + authUrl)
	fmt.Println("Enter redirect response: ")

	var resp string
	//resp := "https://localhost/?oauth_token=oauth_token&oauth_verifier=oauth_verifier"
	fmt.Scanln(&resp)
	token, err := app.GenerateAccessToken(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get user token: %v", token)

	cli := app.GetUserClient()
	followers, err := cli.Users.GetFollowers("Your id", twitter.FollowsOpts{})
	fmt.Println(followers, err)
}
