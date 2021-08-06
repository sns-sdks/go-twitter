package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestCountsRecent() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/counts/recent",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.CountsRecent(TweetCountsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/counts/recent",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"end":"2021-06-02T23:00:00.000Z","start":"2021-06-02T22:00:00.000Z","tweet_count":3506},{"end":"2021-06-03T00:00:00.000Z","start":"2021-06-02T23:00:00.000Z","tweet_count":2544}],"meta":{"total_tweet_count":744364}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.CountsRecent(TweetCountsOpts{Query: "lakers"})
	bc.Equal(len(resp.Data), 2)
	bc.Equal(*resp.Data[0].TweetCount, 3506)
	bc.Equal(*resp.Meta.TotalTweetCount, 744364)
}

func (bc *BCSuite) TestCountsAll() {

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/counts/all",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.CountsALL(TweetCountsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/counts/all",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"end":"2021-01-02T00:00:00.000Z","start":"2021-01-01T00:00:00.000Z","tweet_count":18392},{"end":"2021-01-03T00:00:00.000Z","start":"2021-01-02T00:00:00.000Z","tweet_count":46914},{"end":"2021-01-04T00:00:00.000Z","start":"2021-01-03T00:00:00.000Z","tweet_count":20441},{"end":"2021-01-05T00:00:00.000Z","start":"2021-01-04T00:00:00.000Z","tweet_count":33500},{"end":"2021-01-06T00:00:00.000Z","start":"2021-01-05T00:00:00.000Z","tweet_count":21046},{"end":"2021-01-07T00:00:00.000Z","start":"2021-01-06T00:00:00.000Z","tweet_count":48095},{"end":"2021-01-08T00:00:00.000Z","start":"2021-01-07T00:00:00.000Z","tweet_count":26190},{"end":"2021-01-09T00:00:00.000Z","start":"2021-01-08T00:00:00.000Z","tweet_count":48977},{"end":"2021-01-10T00:00:00.000Z","start":"2021-01-09T00:00:00.000Z","tweet_count":49339},{"end":"2021-01-11T00:00:00.000Z","start":"2021-01-10T00:00:00.000Z","tweet_count":13753},{"end":"2021-01-12T00:00:00.000Z","start":"2021-01-11T00:00:00.000Z","tweet_count":44699},{"end":"2021-01-13T00:00:00.000Z","start":"2021-01-12T00:00:00.000Z","tweet_count":19848},{"end":"2021-01-14T00:00:00.000Z","start":"2021-01-13T00:00:00.000Z","tweet_count":161169},{"end":"2021-01-15T00:00:00.000Z","start":"2021-01-14T00:00:00.000Z","tweet_count":85194}],"meta":{"total_tweet_count":637557}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.CountsALL(TweetCountsOpts{
		Query: "lakers", Granularity: "day",
		StartTime: "2020-01-01T00%3A00%3A00Z", EndTime: "2020-01-15T00%3A00%3A00Z",
	})
	bc.Equal(len(resp.Data), 14)
	bc.Equal(*resp.Data[0].TweetCount, 18392)
	bc.Equal(*resp.Meta.TotalTweetCount, 637557)
}
