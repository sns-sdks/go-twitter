package main

import (
	"fmt"
	"github.com/sns-sdks/go-twitter/twitter"
)

var (
	ClientID          = "Your app client ID"
	OAuth2CallbackURL = "https://localhost/" // Your redirect uri
	// ClientSecret      = "Your app client secret"
)

func main() {
	app := twitter.OAuth2AuthorizationAPP{
		ClientID:    ClientID,
		CallbackURL: OAuth2CallbackURL,
		Scopes:      []string{"tweet.read", "users.read"},
	}
	// If your app is `confidential client`, you can initial as follows
	/*
		app := twitter.OAuth2AuthorizationAPP{
			ClientID:     ClientID,
			ClientSecret: ClientSecret,
			CallbackURL:  OAuth2CallbackURL,
			Scopes:       []string{"tweet.read", "users.read"},
		}
	*/

	authUrl, verifier, _ := app.GetOAuth2AuthorizationURL()

	fmt.Println("Click the authorization url: " + authUrl)
	fmt.Println("Enter redirect response: ")

	var code string
	// resp := "https://localhost/?code=code" -> Code
	fmt.Scanln(&code)
	token, err := app.GenerateAccessToken(code, verifier)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get user token: %v", token)

	cli := app.GetUserClient()
	user, err := cli.Users.LookupMe(twitter.UserOpts{})
	fmt.Println(user, err)
}
