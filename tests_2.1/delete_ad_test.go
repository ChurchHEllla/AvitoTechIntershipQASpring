package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAdByIDSuccess проверяет успешное удаление объявления по ID
func (suite *AdsAPITestSuiteV1) TestDeleteAdValid() {
	ad := CreateAdData(
		suite.sellerID,
		"smartphone Vivo",
		530,
		54,
		145,
		2989,
	)
	jsonData, err := json.Marshal(ad)
	assert.NoError(suite.T(), err)

	resp, err := suite.client.Post("/api/1/item", jsonData)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), 200, resp.StatusCode)

	
	var response Ans
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	adId, err := ExtractIDFromStatusResponse(response)
	assert.NoError(suite.T(), err)
	respDel, err := suite.client.Delete(fmt.Sprintf("/api/2/item/%s", adId))
	assert.NoError(suite.T(), err)
	defer respDel.Body.Close()

	assert.Equal(suite.T(), 200, respDel.StatusCode, "Должен вернуть 200 для удаленного ID %s", adId)
}

// TestGetAdByIDNotFound проверяет обработку несуществующего ID
func (suite *AdsAPITestSuiteV1) TestDeleteAdNotFound() {
	
	deletedAds, err := loadDelTestAds("./utils/TestDataDeleted.json")
    assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), len(deletedAds), 1, "Должно быть хотя бы одно тестовое объявление в файле")

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
            respDel, err := suite.client.Delete(fmt.Sprintf("/api/2/item/%s", tc.id))
			assert.NoError(suite.T(), err)
			defer respDel.Body.Close()

            assert.Equal(t, 404, respDel.StatusCode, "Должен вернуть 404 для несуществующего ID %s", tc.id)

        })
    }
}

// TestGetAdByIDInvalidFormat проверяет обработку некорректных форматов ID
func (suite *AdsAPITestSuiteV1) TestDeleteAdInvalidFormat() {
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
            respDel, err := suite.client.Delete(fmt.Sprintf("/api/2/item/%s", tc.id))
			assert.NoError(suite.T(), err)
			defer respDel.Body.Close()

            assert.Equal(t, 400, respDel.StatusCode, "Должен вернуть 400 при некорректном формате ID %s", tc.id)

        })
    }
}
