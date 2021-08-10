package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

// BCSuite For the tests with app context
type BCSuite struct {
	suite.Suite
	Tw *Client
}

func (bc *BCSuite) SetupSuite() {
	bc.Tw = NewBearerClient("")
}

func (bc *BCSuite) SetupTest() {
	httpmock.ActivateNonDefault(bc.Tw.Cli.GetClient())
}

func (bc *BCSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

// UCSuite For the tests with user context
type UCSuite struct {
	suite.Suite
	Tw *Client
}

func (uc *UCSuite) SetupSuite() {
	app := NewAuthorizationAPP(AuthorizationAPP{
		ConsumerKey:       "",
		ConsumerSecret:    "",
		AccessTokenKey:    "",
		AccessTokenSecret: "",
	})
	uc.Tw = app.GetUserClient()
}

func (uc *UCSuite) SetupTest() {
	httpmock.ActivateNonDefault(uc.Tw.Cli.GetClient())
}

func (uc *UCSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}
