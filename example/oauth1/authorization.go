package main

import (
	"fmt"
	"github.com/sns-sdks/go-twitter/twitter"
)

func main() {
	app := twitter.AuthorizationAPP{
		ConsumerKey:    "app consumer key",
		ConsumerSecret: "app consumer secery",
		CallbackURL:    "https://localhost/",
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
