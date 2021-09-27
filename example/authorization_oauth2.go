package main

import (
	"fmt"
	"github.com/sns-sdks/go-twitter/twitter"
)

func main() {
	app := twitter.OAuth2AuthorizationAPP{
		ClientID:    "client id",
		CallbackURL: "https://localhost/",
	}
	authUrl, verifier := app.GetOAuth2AuthorizationURL()

	fmt.Println("Click the authorization url: " + authUrl)
	fmt.Println("Enter redirect response: ")

	var code string
	//resp := "https://localhost/?code=code" -> Code
	fmt.Scanln(&code)
	token, err := app.GenerateAccessToken(code, verifier)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get user token: %v", token)

	cli := app.GetUserClient()
	followers, err := cli.Users.GetFollowers("Your id", twitter.FollowsOpts{})
	fmt.Println(followers, err)
}
