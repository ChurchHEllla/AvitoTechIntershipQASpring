package tests

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test8GetAdsBySellerSuccess проверяет успешное получение объявлений по SellerID
func (suite *AdsAPITestSuiteV1) Test8GetAdsBySellerSuccess() {
	resp, err := suite.client.Get(fmt.Sprintf("/api/1/%s/item",strconv.Itoa(suite.sellerID)))
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), 200, resp.StatusCode, fmt.Sprintf("Должен вернуть 200 для существующего ID %s", strconv.Itoa(suite.sellerID)))

	var response []AdResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), len(response), 1)

	ads, err := loadTestAds("./utils/TestData.json")
	assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), len(ads), 1, "Должно быть хотя бы одно тестовое объявление в файле")

	var filAds []Ad
    for _, ad := range ads {
        if ad.SellerID == suite.sellerID {
            filAds = append(filAds, ad)
        }
    }
	
	for _, ad := range response {
	 if filAds[0].ID == ad.ID {
		assert.Equal(suite.T(), filAds[0].Name, ad.Name)
		assert.Equal(suite.T(), filAds[0].Price, ad.Price)
		assert.Equal(suite.T(), filAds[0].SellerID, ad.SellerID)
		assert.Equal(suite.T(), filAds[0].Statistics.Contacts, ad.Statistics.Contacts)
		assert.Equal(suite.T(), filAds[0].Statistics.Likes, ad.Statistics.Likes)
		assert.Equal(suite.T(), filAds[0].Statistics.ViewCount, ad.Statistics.ViewCount)
	 }
	}	
}
//Test9GetAdsBySellerBadRequest проверяет обработку некорректных форматов SellerID
func (suite *AdsAPITestSuiteV1) Test9GetAdsBySellerBadRequest() {
    testCases := []struct {
        name string
        id   string
    }{
        {"Отрицательный ID", "-400"},
		{"Слишком большой ID", "999999009"},
		{"Некорректный ID", "Charmander"},
    }

    for _, tc := range testCases {
        suite.T().Run(tc.name, func(t *testing.T) {
			resp, err := suite.client.Get(fmt.Sprintf("/api/1/%s/item", tc.id))
			assert.NoError(suite.T(), err)
			defer resp.Body.Close()

            assert.Equal(t, 400, resp.StatusCode, fmt.Sprintf("Должен вернуть 400 для некорректного ID %s", tc.id))

    
        })
    }
}