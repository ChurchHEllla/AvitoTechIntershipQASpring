package tests

import (
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// TestCompabilityStatisticsByIDSucces проверяет консистентность разных версий ручек общие поля

func (suite *AdsAPITestSuiteV2) TestCompabilityStatisticsByIDSuccess() {

	ads, err := loadTestAds("./utils/TestData.json")
	assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), len(ads), 1, "Должно быть хотя бы одно тестовое объявление в файле")

	testAd := ads[0] 

	respGet, err := suite.client.Get(fmt.Sprintf("/api/1/statistic/%s", testAd.ID))
	assert.NoError(suite.T(), err)
	defer respGet.Body.Close()

	assert.Equal(suite.T(), 200, respGet.StatusCode, "Должен вернуть 200 для существующего ID")

	var response []Statistics
	err = json.NewDecoder(respGet.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	responseAd := response[0]

	respGet1, err := suite.client.Get(fmt.Sprintf("/api/2/statistic/%s", testAd.ID))
	assert.NoError(suite.T(), err)
	defer respGet1.Body.Close()

	assert.Equal(suite.T(), 200, respGet1.StatusCode, "Должен вернуть 200 для существующего ID")

	var response2 []Statistics
	err = json.NewDecoder(respGet1.Body).Decode(&response2)
	assert.NoError(suite.T(), err)
    responseAd2 := response2[0]

	assert.Equal(suite.T(), responseAd2.Contacts, responseAd.Contacts, "Должно быть одинаковые contacts")
	assert.Equal(suite.T(), responseAd2.Likes, responseAd.Likes, "Должно быть одинаковые likes")
	assert.Equal(suite.T(), responseAd2.ViewCount, responseAd.ViewCount, "Должно быть одинаковые viewCount")
}