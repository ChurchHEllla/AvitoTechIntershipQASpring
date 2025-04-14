package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
var (
	BaseURL = "https://qa-internship.avito.com"
	BaseSellerID = 880055
)

// Берет на входе отформатированный ответ запроса и возвращает только id
func ExtractIDFromStatusResponse(input Ans) (string, error) {

	parts := strings.Fields(input.Status)
    if len(parts) == 0 {
        return "", fmt.Errorf("пустая строка")
    }

    // Берем последнюю часть
    lastPart := parts[len(parts)-1]

    // Проверяем, что это валидный UUID (упрощенная проверка)
    if len(lastPart) != 36 ||
        strings.Count(lastPart, "-") != 4 {
        return "", fmt.Errorf("неверный формат UUID")
    }

    return lastPart, nil
}

//Объявления структур для работы с данными
type Ad struct {
	ID         string     `json:"id"`
	SellerID   int        `json:"sellerID"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Statistics Statistics `json:"statistics"`
}
type AdResp struct {
	CreatedAt  string	  `json:"createdAt"`
    ID 		   string	  `json:"id"`	
	SellerID   int        `json:"sellerID"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Statistics Statistics `json:"statistics"`
}
type AdDeleted struct {	
    ID 		   string	  `json:"id"`	
}
type Statistics struct {
	Contacts  int `json:"contacts"`
	Likes     int `json:"likes"`
	ViewCount int `json:"viewCount"`
}
type Ans struct{
	Status string `json:"status"`
}

//Создает тестовые данные в нужном формате
func CreateAdData(sellerID int, name string, price int, contacts int, likes int, viewCount int) map[string]interface{} {
    return map[string]interface{}{
        "sellerID": sellerID,
        "name":     name,
        "price":    price,
        "statistics": map[string]interface{}{
            "contacts":  contacts,
            "likes":    likes,
            "viewCount": viewCount,
        },
    }
}

//Загружают тестовые данные из json
func loadTestAds(filename string) ([]Ad, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ads []Ad
	err = json.NewDecoder(file).Decode(&ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func loadDelTestAds(filename string) ([]AdDeleted, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ads []AdDeleted
	err = json.NewDecoder(file).Decode(&ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}