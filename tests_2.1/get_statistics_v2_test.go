package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAdByIDSuccess проверяет успешное получение объявления по ID
func (suite *AdsAPITestSuiteV1) TestGetStatisticsV2ByIDSuccess() {

	ads, err := loadTestAds("./utils/TestData.json")
	assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), len(ads), 1, "Должно быть хотя бы одно тестовое объявление в файле")

	testAd := ads[0] 

	
	respGet, err := suite.client.Get(fmt.Sprintf("/api/2/statistic/%s", testAd.ID))
	assert.NoError(suite.T(), err)
	defer respGet.Body.Close()

	assert.Equal(suite.T(), 200, respGet.StatusCode, "Должен вернуть 200 для существующего ID")

	var response []Statistics
	err = json.NewDecoder(respGet.Body).Decode(&response)
	assert.NoError(suite.T(), err)
    responseAd := response[0]
	
	assert.Equal(suite.T(), testAd.Statistics.Contacts, responseAd.Contacts)
	assert.Equal(suite.T(), testAd.Statistics.Likes, responseAd.Likes)
	assert.Equal(suite.T(), testAd.Statistics.ViewCount, responseAd.ViewCount)
}
// TestGetAdByIDNotFound проверяет обработку несуществующего ID
func (suite *AdsAPITestSuiteV1) TestGetAStatisticsV2ByIDNotFound() {
	deletedAds, err := loadDelTestAds("./utils/TestDataDeleted.json")
    assert.NoError(suite.T(), err)
    assert.True(suite.T(), len(deletedAds) > 0, "Должен быть хотя бы один тестовый id в файле")

    //testAd := ads[1]
    testDelAd := deletedAds[1]
    testCases := []struct {
        name string
        id   string
    }{
        {"Не существующий UUID", "00000000-0000-0000-0000-000000000000"},
        {"Недавно удаленный ID", testDelAd.ID},
    }

    for _, tc := range testCases {
        suite.T().Run(tc.name, func(t *testing.T) {
			resp, err := suite.client.Get(fmt.Sprintf("/api/2/statistic/%s", tc.id))
            assert.NoError(t, err)
            defer resp.Body.Close()

            assert.Equal(t, 404, resp.StatusCode, "Должен вернуть 404 для несуществующего ID")
 
        })
    }
}

// TestGetAdByIDInvalidFormat проверяет обработку некорректных форматов ID
func (suite *AdsAPITestSuiteV1) TestGetStatisticsV2ByIDInvalidFormat() {
    testCases := []struct {
        name string
        id   string
    }{
        {"Пустой ID", ""},
        {"Короткий ID", "123"},
        {"Non-UUID формат", "item_12345"},
    }

    for _, tc := range testCases {
        suite.T().Run(tc.name, func(t *testing.T) {
            resp, err := suite.client.Get(fmt.Sprintf("/api/2/statistic/%s", tc.id))
            assert.NoError(t, err)
            defer resp.Body.Close()

            assert.Equal(t, 400, resp.StatusCode, "Должен вернуть 400 при некорректном формате ID")

 
        })
    }
}

