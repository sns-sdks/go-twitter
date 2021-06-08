package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

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
