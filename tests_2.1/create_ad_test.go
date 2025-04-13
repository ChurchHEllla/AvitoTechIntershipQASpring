package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateAdValid проверяет создание объявления с валидными данными
func (suite *AdsAPITestSuiteV1) TestCreateAdValid() {
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

}

// TestCreateAdMissingRequiredFields проверяет обработку отсутствия обязательных полей
func (suite *AdsAPITestSuiteV1) TestCreateAdMissingRequiredFields() {
	testCases := []struct {
		name        string
		adData      map[string]interface{}
		description string
	}{
		{
			name: "Потерянный sellerID",
			adData: map[string]interface{}{
				"price":    500,
				"name": "смартфон Vivo",
				"statistics": map[string]interface{}{
					"contacts":  800,
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при отсутствии sellerID",
		},
		{
			name: "Потерянный name",
			adData: map[string]interface{}{
				"sellerID": suite.sellerID,
				"price":    500,
				"statistics": map[string]interface{}{
					"contacts":  800,
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при отсутствии name",
		},
		{
			name: "Потерянный price",
			adData: map[string]interface{}{
				"sellerID": suite.sellerID,
				"name":     "Телефон без цены",
				"statistics": map[string]interface{}{
					"contacts":  800,
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при отсутствии price",
		},
		{
			name: "Потерянный statistics",
			adData: map[string]interface{}{
				"sellerID": suite.sellerID,
				"name":     "Телефон без статистики",
				"price":    500,
			},
			description: "Должен вернуть 400 при отсутствии statistics",
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.adData)
			assert.NoError(t, err)

			resp, err := suite.client.Post("/api/1/item", jsonData)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, 400, resp.StatusCode, tc.description)
		})
	}
}

// TestCreateAdInvalidDataTypes проверяет обработку некорректных типов данных
func (suite *AdsAPITestSuiteV1) TestCreateAdInvalidDataTypes() {
	testCases := []struct {
		name        string
		adData      map[string]interface{}
		description string
	}{
		{
			name: "Неверный тип sellerID",
			adData: map[string]interface{}{
				"sellerID": "не число",
				"name":     "Телефон",
				"price":    500,
				"statistics": map[string]interface{}{
					"contacts":  800,
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при строковом sellerID вместо числа",
		},
		{
			name: "Неверный тип price",
			adData: map[string]interface{}{
				"sellerID": suite.sellerID,
				"name":     "Телефон",
				"price":    "пятьсот",
				"statistics": map[string]interface{}{
					"contacts":  "восемьсот",
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при строковом price вместо числа",
		},
		{
			name: "Неверный тип contacts в statistics",
			adData: map[string]interface{}{
				"sellerID": suite.sellerID,
				"name":     "Телефон",
				"price":    500,
				"statistics": map[string]interface{}{
					"contacts":  "восемьсот",
					"likes":     555,
					"viewCount": 3535,
				},
			},
			description: "Должен вернуть 400 при строковом contacts в statistics",
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.adData)
			assert.NoError(t, err)

			resp, err := suite.client.Post("/api/1/item", jsonData)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, 400, resp.StatusCode, tc.description)
		})
	}
}

// TestCreateAdNegativeValues проверяет обработку отрицательных значений в полях
func (suite *AdsAPITestSuiteV1) TestCreateAdNegativeValues() {
	testCases := []struct {
		name        string
		adData      map[string]interface{}
		description string
	}{
		{
			name: "Отрицательный sellerID",
			adData: CreateAdData(
				-1,
				"Телефон с отрицательным sellerID",
				500,
				300,
				555,
				3535,
			),
			description: "Должен вернуть 400 при отрицательном sellerID",
		},
		{
			name: "Отрицательный price",
			adData: CreateAdData(
				suite.sellerID,
				"Телефон с отрицательной ценой",
				-100,
				300,
				555,
				3535,
			),
			description: "Должен вернуть 400 при отрицательном price",
		},
		
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.adData)
			assert.NoError(t, err)

			resp, err := suite.client.Post("/api/1/item", jsonData)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, 400, resp.StatusCode, tc.description)
		})
	}
}