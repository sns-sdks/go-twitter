package twitter

import (
	"errors"
	"github.com/dghubble/oauth1"
	"net/http"
	"net/url"
)

const (
	oauthTokenParam    = "oauth_token"
	oauthVerifierParam = "oauth_verifier"
)

var AuthorizeEndpoint = oauth1.Endpoint{
	RequestTokenURL: "https://api.twitter.com/oauth/request_token",
	AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
	AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
}

// AuthorizationAPP Twitter app config
type AuthorizationAPP struct {
	ConsumerKey       string         `json:"consumer_key"`
	ConsumerSecret    string         `json:"consumer_secret"`
	CallbackURL       string         `json:"callback_url,omitempty"`
	AccessTokenKey    string         `json:"access_token_key,omitempty"`
	AccessTokenSecret string         `json:"access_token_secret,omitempty"`
	RequestSecret     string         `json:"request_secret,omitempty"`
	Config            *oauth1.Config `json:"config,omitempty"`
}

func (app AuthorizationAPP) String() string {
	return Stringify(app)
}

// GetAuthorizationURL Get authorization url for user
func (app *AuthorizationAPP) GetAuthorizationURL() (string, error) {
	app.Config = &oauth1.Config{
		ConsumerKey:    app.ConsumerKey,
		ConsumerSecret: app.ConsumerSecret,
		CallbackURL:    app.CallbackURL,
		Endpoint:       AuthorizeEndpoint,
	}

	reqToken, reqSecret, err := app.Config.RequestToken()
	if err != nil {
		return "", err
	}

	app.RequestSecret = reqSecret

	authorizationURL, err := app.Config.AuthorizationURL(reqToken)
	if err != nil {
		return "", err
	}
	return authorizationURL.String(), err
}

// GenerateAccessToken Generate user access token for the app
func (app *AuthorizationAPP) GenerateAccessToken(response string) (*oauth1.Token, error) {
	qUrl, err := url.Parse(response)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(qUrl.RawQuery)
	if err != nil {
		return nil, err
	}
	requestToken := values.Get(oauthTokenParam)
	verifier := values.Get(oauthVerifierParam)
	if requestToken == "" || verifier == "" {
		return nil, errors.New("oauth1: Request missing oauth_token or oauth_verifier")
	}

	accessToken, accessSecret, err := app.Config.AccessToken(requestToken, app.RequestSecret, verifier)
	if err != nil {
		return nil, err
	}
	app.AccessTokenKey = accessToken
	app.AccessTokenSecret = accessSecret

	return oauth1.NewToken(accessToken, accessSecret), nil
}

// GetAuthorizedHttpClient Get user authorized http client
func (app *AuthorizationAPP) GetAuthorizedHttpClient() *http.Client {
	config := oauth1.NewConfig(app.ConsumerKey, app.ConsumerSecret)
	token := oauth1.NewToken(app.AccessTokenKey, app.AccessTokenSecret)
	hc := config.Client(oauth1.NoContext, token)
	return hc
}

// GetUserClient get library client with user authorization
func (app *AuthorizationAPP) GetUserClient() *Client {
	hc := app.GetAuthorizedHttpClient()
	return NewUserClint(hc)
}
