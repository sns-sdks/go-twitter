package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

type Auth2Suite struct {
	suite.Suite
	app *OAuth2AuthorizationAPP
}

func (auth *Auth2Suite) SetupSuite() {
	auth.app = NewOAuth2AuthorizationAPP(OAuth2AuthorizationAPP{
		ClientID:    "client id",
		CallbackURL: "https://localhost/",
		Scopes:      []string{"users.read", "tweet.read"},
	})
}

func (auth *Auth2Suite) SetupTest() {
	httpmock.ActivateNonDefault(http.DefaultClient)
}

func (auth *Auth2Suite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func TestAuth2Suite(t *testing.T) {
	suite.Run(t, new(Auth2Suite))
}

func (auth *Auth2Suite) TestGetAuthorizationURL() {
	authUrl, verifier := auth.app.GetOAuth2AuthorizationURL()
	auth.NotNil(authUrl)
	auth.NotNil(verifier)
}

func (auth *Auth2Suite) TestGenerateAccessToken() {
	httpmock.RegisterResponder(
		HttpPost, OAuth2Endpoint.TokenURL,
		httpmock.NewStringResponder(400, `{"error":"invalid_request","error_description":"Value passed for the authorization code was invalid."}`),
	)
	_, err := auth.app.GenerateAccessToken("code", "verifier")
	auth.Contains(err.Error(), "invalid")

	httpmock.RegisterResponder(
		HttpPost, OAuth2Endpoint.TokenURL,
		httpmock.NewStringResponder(
			200,
			`{"access_token":"access_token","expires_in":7200,"scope":"users.read tweet.read","token_type":"bearer"}`,
		),
	)

	token, _ := auth.app.GenerateAccessToken("code", "verifier")
	auth.IsType(&oauth2.Token{}, token)

	cli := auth.app.GetUserClient()
	auth.IsType(&Client{}, cli)
}

func (auth *Auth2Suite) TestGenerateCodeVerifier() {
	v1 := GenerateCodeVerifier(150)
	auth.NotNil(v1)
	v2 := GenerateCodeVerifier(20)
	auth.NotNil(v2)
}
