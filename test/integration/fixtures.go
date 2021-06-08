package integration

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	gt "go-twitter/twitter"
)

type BCSuite struct {
	suite.Suite
	Tw *gt.Client
}

func (bc *BCSuite) SetupSuite() {
	bc.Tw = gt.NewBearerClient("")
}

func (bc *BCSuite) SetupTest() {
	httpmock.ActivateNonDefault(bc.Tw.Cli.GetClient())
}

func (bc *BCSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}
