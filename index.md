## Indices

* [Conversion](#conversion)

  * [Add Conversion](#1-add-conversion)
  * [Get Conversion](#2-get-conversion)

* [Currency](#currency)

  * [Add Currency](#1-add-currency)
  * [Get Currency](#2-get-currency)


--------


## Conversion



### 1. Add Conversion



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7070/v1/conversion/add
```



**Request**
***Body:***

```js        
{
    "currencyIDFrom": 1,
    "currencyIDTo": 2,
    "rate": 50
}
```

**Response**
***Body:***

```js
{
    "code": "success",
    "message": "success",
    "data": {}
}
```


### 2. Get Conversion



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:7070/v1/conversion
```


**Request**
***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| currencyIDFrom | 1 |  |
| currencyIDTo | 2 |  |
| amount | 5000 |  |


**Response**
***Body:***

```js
{
    "code": "success",
    "message": "success",
    "data": {
        "result": 100
    }
}
```


## Currency



### 1. Add Currency



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7070/v1/currency/add
```


**Request**
***Body:***

```js        
{
    "name": "currency-2"
}
```

**Response**
***Body:***

```js
{
    "code": "success",
    "message": "success",
    "data": {}
}
```



### 2. Get Currency



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:7070/v1/currency
```

**Response**
***Body:***

```js
{
    "code": "success",
    "message": "success",
    "data": [
        {
            "id": 1,
            "name": "currency-1"
        },
        {
            "id": 2,
            "name": "currency-2"
        },
        {
            "id": 3,
            "name": "currency-2"
        }
    ]
}
```
