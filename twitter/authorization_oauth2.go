package twitter

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

// OAuth2Endpoint endpoint for twitter oauth2
var OAuth2Endpoint = oauth2.Endpoint{
	AuthURL:  "https://twitter.com/i/oauth2/authorize",
	TokenURL: "https://api.twitter.com/2/oauth2/token",
}

// OAuth2AuthorizationAPP Twitter OAuth2 app config
type OAuth2AuthorizationAPP struct {
	ClientID    string         `json:"client_id"`
	CallbackURL string         `json:"callback_url,omitempty"`
	Scopes      []string       `json:"scopes,omitempty"`
	Token       *oauth2.Token  `json:"access_token,omitempty"`
	Config      *oauth2.Config `json:"config,omitempty"`
}

func (app OAuth2AuthorizationAPP) String() string {
	return Stringify(app)
}

// NewOAuth2AuthorizationAPP Return app for oauth2 authorization
func NewOAuth2AuthorizationAPP(app OAuth2AuthorizationAPP) *OAuth2AuthorizationAPP {
	app.Config = &oauth2.Config{
		ClientID:    app.ClientID,
		RedirectURL: app.CallbackURL,
		Scopes:      app.Scopes,
		Endpoint:    OAuth2Endpoint,
	}
	return &app
}

// GetOAuth2AuthorizationURL Return authorization url and code verifier for user
func (app *OAuth2AuthorizationAPP) GetOAuth2AuthorizationURL() (string, string) {
	state := GenerateNonce()
	verifier := GenerateCodeVerifier(128)

	challengeOpt := oauth2.SetAuthURLParam("code_challenge", PkCEChallengeWithSHA256(verifier))
	challengeMethodOpt := oauth2.SetAuthURLParam("code_challenge_method", "s256")

	return app.Config.AuthCodeURL(state, challengeOpt, challengeMethodOpt), verifier
}

// GenerateAccessToken Generate user access token for the app
func (app *OAuth2AuthorizationAPP) GenerateAccessToken(code, verifier string) (*oauth2.Token, error) {
	ctx := context.Background()
	verifierOpt := oauth2.SetAuthURLParam("code_verifier", verifier)
	token, err := app.Config.Exchange(ctx, code, verifierOpt)
	if err != nil {
		return nil, err
	}
	app.Token = token
	return token, err
}

// GetAuthorizedHttpClient Get user authorized http client
func (app *OAuth2AuthorizationAPP) GetAuthorizedHttpClient() *http.Client {
	hc := app.Config.Client(context.TODO(), app.Token)
	return hc
}

// GetUserClient get library client with user authorization
func (app *OAuth2AuthorizationAPP) GetUserClient() *Client {
	hc := app.GetAuthorizedHttpClient()
	return NewUserClint(hc)
}

// PkCEChallengeWithSHA256 base64-URL-encoded SHA256 hash of verifier, per rfc 7636
func PkCEChallengeWithSHA256(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	challenge := b64.RawURLEncoding.EncodeToString(sum[:])
	return challenge
}

// GenerateCodeVerifier Generate code verifier (length 43~128) for PKCE.
func GenerateCodeVerifier(length int) string {
	if length > 128 {
		length = 128
	}
	if length < 43 {
		length = 43
	}
	return randStringBytes(length)
}

// GenerateNonce Generate random nonce.
func GenerateNonce() string {
	return b64.RawURLEncoding.EncodeToString([]byte(randStringBytes(8)))
}

// randStringBytes Return random string by length
func randStringBytes(n int) string {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		return ""
	}
	return b64.RawURLEncoding.EncodeToString(b[:])
}
