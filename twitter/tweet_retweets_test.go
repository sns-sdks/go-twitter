package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestRetweetedBy() {
	tid := "1354143047324299264"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/retweeted_by",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1065249714214457345","name":"Spaces","username":"TwitterSpaces"},{"id":"783214","name":"Twitter","username":"Twitter"},{"id":"1526228120","name":"Twitter Data","username":"TwitterData"},{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"},{"id":"6253282","name":"Twitter API","username":"TwitterAPI"}],"meta":{"result_count":5}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetRetweetedBy(tid, UserOpts{})
	bc.Equal(len(resp.Data), 5)
}