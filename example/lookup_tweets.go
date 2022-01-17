package main

import (
	"fmt"
	"github.com/sns-sdks/go-twitter/twitter"
)

const BearerToken = "" // Your access token

func main() {
	tw := twitter.NewBearerClient(BearerToken)

	tweetsResp, err := tw.Tweets.LookupByIDs("1460324089726320643,1460323737035677698", twitter.TweetOpts{TweetFields: "id,text,created_at"})
	if err != nil {
		fmt.Println("Not get tweets info")
		return
	}
	fmt.Printf("Tweets count: %d", len(tweetsResp.Data))
	tweet := tweetsResp.Data[0]
	fmt.Printf("Tweet ID: %v\n", *tweet.ID)
	fmt.Printf("Tweet text: %v\n", *tweet.Text)
	fmt.Printf("Tweet Created At: %v\n", *tweet.CreatedAt)
	fmt.Printf("Now ratelimit: %v", tw.RateLimit.Get())
}
