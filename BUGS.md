# Баг-репорт

## BUG1. Негативное тестирование. Успешная отправка некоректных данных 

### Описание:
Запрос прошел успешно при отсутствиии заполнения необходимых полей
данных

Используйте следующую информацию для воспроизведения бага:

| sellerID | name | price | contacts | likes | views|
|-------------|-------------|-------------|-------------|-------------|-------------|
|          | смартфон Vivo    |  500  |  800 | 555 | 3535 
|  880055  |                  |  500  |  800 | 555 | 3535 
|  880055  | смартфон Vivo    |       |  800 | 555 | 3535 
|  880055  | смартфон Vivo    |  500  | 

### Приоритет: 
medium

### Шаги для воспроизведения
Отправить POST-запрос на 
[https://qa-internship.avito.com/api/1/item](`https://qa-internship.avito.com/api/1/item`) с невалидными данными объявления
### Ожидаемый результат
Ожидалось получение статуса 400 Bad Request и сообщение об ошибке

### Фактический результат

получение статуса 200 ок

### Подробная информация:
```
Error:          Not equal:
                                expected: 400
                                actual  : 200

Test:           TestAdsAPISuiteV1/Test2CreateAdMissingRequiredFields/Потерянный_statistics
Messages:       Должен вернуть 400 при отсутствии statistics

Test:           TestAdsAPISuiteV1/Test2CreateAdMissingRequiredFields/Потерянный_statistics
Messages:       Должен вернуть 400 при отсутствии statistics

Test:           TestAdsAPISuiteV1/Test2CreateAdMissingRequiredFields/Потерянный_name
Messages:       Должен вернуть 400 при отсутствии name

Messages:       Должен вернуть 400 при отрицательном sellerID
```
## BUG2. Негативное тестирование. Успешная отправка данных при несоблюдении валидных значений

### Описание:
При отправке запроса допускаются отрицательные значения числовых данных, что нелогично

Используйте следующую информацию для воспроизведения бага:

| sellerID | name | price | contacts | likes | views|
|-------------|-------------|-------------|-------------|-------------|-------------|
|  -1  | смартфон Vivo    |  500  |  300 | 555 | 3535 
|  880055  | смартфон Vivo    |  -100  |  300 | 555 | 3535 

### Приоритет: 
high

### Шаги для воспроизведения
Отправить POST-запрос на 
[https://qa-internship.avito.com/api/1/item](`https://qa-internship.avito.com/api/1/item`) с невалидными данными объявления
### Ожидаемый результат
Ожидалось получение статуса 400 Bad Request и сообщение об ошибке

### Фактический результат

получение статуса 200 ОК

### Подробная информация:
```
 Error:          Not equal:
                                expected: 400
                                actual  : 200

Test:           TestAdsAPISuiteV1/Test3CreateAdNegativeValues/Отрицательный_sellerID
Messages:       Должен вернуть 400 при отрицательном sellerID

Test:           TestAdsAPISuiteV1/Test3CreateAdNegativeValues/Отрицательный_price
Messages:       Должен вернуть 400 при отрицательном price
```
## BUG3. Позитивное тестирование. Получение некачественной информации при отправке GET-запроса

### Описание:
при отправке GET запроса пришли данные, несоответствующие желаемым(правильным)

Используйте следующую информацию для воспроизведения бага:
| id |
|-------------|
| 83b2a2f0-d92d-4280-ba5d-79346ec27107    |

Корректная информация:
```
{
    "id": "83b2a2f0-d92d-4280-ba5d-79346ec27107",
    "sellerID": 880055,
    "name": "mashroom",
    "price": 530,
    "statistics": {
      "contacts": 54,
      "likes": 145,
      "viewCount": 2989
    }
}
```
### Приоритет: 
high

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/item/:id](`https://qa-internship.avito.com/api/1/item/:id`) с валидными данными объявления
Сравнить полученные данные с ожидаемыми

### Ожидаемый результат
1. получение статуса 200 ОК
2. Данные корректны

### Фактический результат

1. получение статуса 200 ОК
2. Данные некорректны


### Подробная информация:
```
Error:          Not equal:
                expected: "mashroom"
                actual  : "mashroom 'а вот тут ошибка=)'"

Test:           TestAdsAPISuiteV1/Test5GetAdByIDSuccess
Messages:       name объявления должно совпадать

Test:           TestAdsAPISuiteV1/Test5GetAdByIDSuccess
Messages:       price объявления должно совпадать
```

## BUG4. Негативное тестирование. Получение информации отправке GET-запроса с некорректными данными

### Описание:
при отправке GET запроса должны прийти данные только в том случае, если продавец существует. " Хоть речь идет об объявлениях, а не о продавцах, в данном тесте нет различий между продавцом без объявлений и отсутсвием продавца, возможно это не баг".

Используйте следующую информацию для воспроизведения бага:
| sellerID |
|-------------|
| 0    |

### Приоритет: 
low

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/:sellerID/item](`https://qa-internship.avito.com/api/1/:sellerID/item`) с валидными данными объявления
### Ожидаемый результат
получение статуса 404 Not Founded

### Фактический результат
получение статуса 200 ОК

## BUG5. Позитивное тестирование. Получение некачественной информации по объявлениям продавца при отправке GET-запроса

### Описание:
при отправке GET запроса пришли данные, несоответствующие желаемым(правильным)

Используйте следующую информацию для воспроизведения бага:
| sellerID |
|-------------|
| 880055    |

### Приоритет: 
high

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/:sellerID/item](`https://qa-internship.avito.com/api/1/:sellerID/item`) с валидными данными объявления
### Ожидаемый результат
1. получение статуса 200 ОК
2. корректные данные
### Фактический результат
1. получение статуса 200 ОК
2. некорректные данные

### Подробная информация:
```
Error:      	Not equal: 
        	            	expected: 530
        	            	actual  : 880055

Test:       	TestAdsAPISuiteV1/Test8GetAdsBySellerSuccess

Error:      	Not equal: 
        	            	expected: 880055
        	            	actual  : 1060
Test:       	TestAdsAPISuiteV1/Test8GetAdsBySellerSuccess
```
## BUG6. Негативное тестирование. Получение информации при отправке GET-запроса с некорректным id

### Описание:
при отправке GET запроса пришли данные, несоответствующие желаемым(правильным)

Используйте следующую информацию для воспроизведения бага:
| sellerID |
|-------------|
| -400    |

### Приоритет: 
low

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/:sellerID/item](`https://qa-internship.avito.com/api/1/:sellerID/item`) с валидными данными объявления
### Ожидаемый результат
получение статуса 400 Bad Request

### Фактический результат
получение статуса 200 ОК

### Подробная информация:
```
Error:      	Not equal: 
        	            	expected: 400
        	            	actual  : 200
Test:       	TestAdsAPISuiteV1/Test9GetAdsBySellerBadRequest/Отрицательный_ID
Messages:   	Должен вернуть 400 для некорректного ID -400

```

## BUG7. Позитивное тестирование. Получение некачественной информации по статистике при отправке GET-запроса

### Описание:
При отправке GET запроса пришли  некорректные данные.

Используйте следующую информацию для воспроизведения бага:

| id |
|-------------|
| 83b2a2f0-d92d-4280-ba5d-79346ec27107   |

### Приоритет: 
high

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/statistic/:id](`https://qa-internship.avito.com/api/1/statistic/:id`) с валидными данными объявления
### Ожидаемый результат
1. получение статуса 200 ОК
2. Данные корректны
### Фактический результат

1. получение статуса 200 ОК
2. Данные некорректны

### Подробная информация:
```
Error:          Not equal:
                                expected: 145
                                actual  : 1
Test:           TestAdsAPISuiteV1/Test10GetStatisticsByIDSuccess

Error:          Not equal:
                                expected: 2989
                                actual  : 1
Test:           TestAdsAPISuiteV1/Test10GetStatisticsByIDSuccess
```

## BUG8. Негативное тестирование. Получение  информации по удаленной или несуществующей статистике при отправке GET-запроса

### Описание:
При отправке GET запроса пришли удаленные данные.

Используйте следующую информацию для воспроизведения бага:

| id |
|-------------|
|  fe6f3034-b637-439e-a901-efeed3f9517a   |

### Приоритет: 
high

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/1/statistic/:id](`https://qa-internship.avito.com/api/1/statistic/:id`) с валидными данными объявления
### Ожидаемый результат
1. получение статуса 404 Not Found

### Фактический результат

1. получение статуса 200 ОК


### Подробная информация:
```
Error:      	Not equal: 
        	            	expected: 404
        	            	actual  : 200

Test:       	TestAdsAPISuiteV1/Test11GetAStatisticsByIDNotFound/Недавно_удаленный_ID
Messages:   	Должен вернуть 404 для несуществующего ID
```
## Дополнения к версии 2

### BUG9. Получение информации по статистике (V2)
### Описание:
При выполнении Тест кейсов 10, 11, 12 с незначительным редактированием для 2-й версии (например адрес запроса) тесты показали те же результаты. Также был замечен новый баг: 
Были получены данные при некорректном формате id объявления

Используйте следующую информацию для воспроизведения бага:

| id |
|-------------|
|  fe6f3034-b637-439e-a901-efeed3f9517a   |

### Приоритет: 
medium

### Шаги для воспроизведения
Отправить GET-запрос на [https://qa-internship.avito.com/api/2/statistic/:id](`https://qa-internship.avito.com/api/2/statistic/:id`) с невалидными данными объявления
### Ожидаемый результат
1. получение статуса 400 Bad Request

### Фактический результат

1. получение статуса 200 ОК


### Подробная информация:
```
Error:      	Not equal: 
        	            	expected: 400
        	            	actual  : 200
Test:       	TestAdsAPISuiteV2/Test12V2GetStatisticsByIDInvalidFormat/Non-UUID_формат
Messages:   	Должен вернуть 400 при некорректном формате ID

Test:       	TestAdsAPISuiteV2/Test12V2GetStatisticsByIDInvalidFormat/Non-UUID_формат
Messages:   	Должен вернуть 400 при некорректном формате ID
```


### BUG10. Несоответствие данных в разных версиях запросов:

При отправке запросов на API с разными версиями, получаем разные ответы.

Используйте следующую информацию для воспроизведения бага:

| id |
|-------------|
|  83b2a2f0-d92d-4280-ba5d-79346ec27107   |

### Приоритет: 
low

### Шаги для воспроизведения
1. Отправить GET-запрос на [https://qa-internship.avito.com/api/1/statistic/:id](`https://qa-internship.avito.com/api/1/statistic/:id`) с валидными данными объявления
2. Отправить GET-запрос на [https://qa-internship.avito.com/api/2/statistic/:id](`https://qa-internship.avito.com/api/2/statistic/:id`) с валидными данными объявления
3. Сравнить ответы двух запросов

### Ожидаемый результат
1. получение статуса 200 ОК
2. Данные сходятся

### Фактический результат

1. получение статуса 200 ОК
2. Данные не сходятся

### Подробная информация:
```
Error:      	Not equal: 
        	            	expected: 999999999
        	            	actual  : 54
Test:       	TestAdsAPISuiteV2/Test16CompabilityStatisticsByIDSuccess
Messages:   	Должно быть одинаковые contacts

Error:      	Not equal: 
                expected: 999999999
                actual  : 1
Test:       	TestAdsAPISuiteV2/Test16CompabilityStatisticsByIDSuccess
Messages:   	Должно быть одинаковые likes

Error:      	Not equal: 
                expected: 999999999
                actual  : 1
Test:       	TestAdsAPISuiteV2/Test16CompabilityStatisticsByIDSuccess
Messages:   	Должно быть одинаковые viewCount
```

