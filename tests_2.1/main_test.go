package tests

import (
	u "avito_tests/tests_2.1/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdsAPITestSuiteV1 struct {
	suite.Suite
	sellerID    int
	client  *u.APIClient
}

func (suite *AdsAPITestSuiteV1) SetupTest() {
	// Генерируем уникальный sellerID для тестов
	suite.sellerID = BaseSellerID
	//111111 + int(time.Now().Unix() % (888889))	
	suite.client = u.NewAPIClient(BaseURL)
}

func TestAdsAPISuiteV1(t *testing.T) {
	suite.Run(t, new(AdsAPITestSuiteV1))
}