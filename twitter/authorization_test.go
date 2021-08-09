package twitter

import (
	"github.com/dghubble/oauth1"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/url"
	"testing"
)

type AuthSuite struct {
	suite.Suite
	app *AuthorizationAPP
}

func (auth *AuthSuite) SetupSuite() {
	auth.app = NewAuthorizationAPP(AuthorizationAPP{
		ConsumerKey:    "consumer key",
		ConsumerSecret: "consumer secret",
		CallbackURL:    "https://localhost/",
	})
}

func (auth *AuthSuite) SetupTest() {
	httpmock.ActivateNonDefault(http.DefaultClient)
}

func (auth *AuthSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (auth AuthSuite) TestGetAuthorizationWithWrongResponse() {
	httpmock.RegisterResponder(
		HttpPost, "https://api.twitter.com/oauth/request_token",
		httpmock.NewStringResponder(400, ``),
	)

	_, _, err := auth.app.GetAuthorizationURL()
	auth.Contains(err.Error(), "oauth1")
}

func (auth *AuthSuite) TestGetAuthorizationSuccess() {
	data := url.Values{
		"oauth_token":              []string{"oauth-token"},
		"oauth_token_secret":       []string{"oauth token secret"},
		"oauth_callback_confirmed": []string{"true"},
	}

	httpmock.RegisterResponder(
		HttpPost, "https://api.twitter.com/oauth/request_token",
		httpmock.NewBytesResponder(200, []byte(data.Encode())))

	authUrl, secret, _ := auth.app.GetAuthorizationURL()
	auth.Equal(secret, "oauth token secret")
	auth.Equal(authUrl, "https://api.twitter.com/oauth/authorize?oauth_token=oauth-token")

	auth.app.Config.Endpoint.AuthorizeURL = "your#$%^&*(proper$#$%%^(password"
	_, _, err := auth.app.GetAuthorizationURL()
	auth.Contains(err.Error(), "invalid")
}

func (auth *AuthSuite) TestGenerateAccessTokenWrongResponse() {
	_, err := auth.app.GenerateAccessToken("your#$%^&*(proper$#$%%^(password")
	auth.Contains(err.Error(), "invalid")

	_, err = auth.app.GenerateAccessToken("https://example.com?123")
	auth.Contains(err.Error(), "oauth1")

	data := url.Values{
		"oauth_token":        []string{"oauth-token"},
		"oauth_token_secret": []string{"oauth token secret"},
	}

	httpmock.RegisterResponder(
		HttpPost, "https://api.twitter.com/oauth/access_token",
		httpmock.NewBytesResponder(200, []byte(data.Encode())),
	)

	token, _ := auth.app.GenerateAccessToken("https://localhost/?oauth_token=oauth_token&oauth_verifier=oauth_verifier")
	auth.IsType(&oauth1.Token{}, token)

	cli := auth.app.GetUserClient()
	auth.IsType(&Client{}, cli)

	auth.app.Config.Endpoint.AccessTokenURL = ""
	_, err = auth.app.GenerateAccessToken("https://localhost/?oauth_token=oauth_token&oauth_verifier=oauth_verifier")
	auth.Contains(err.Error(), "responder")

}
