Go Vidflex Test
===============

## 🛠️ Installation Steps

1. Clone the repository

```Bash
git clone https://github.com/xadrijo/vidflex-test.git
```

2. Open a terminal in vidfelx-test directory
   
3. Run docker-compose command (Docker should ON)

```Bash
docker-compose up --build
```

4. After the app is running:
- Create some products
```
curl --location --request POST 'http://localhost:8000/api/product' \
--header 'Content-Type: application/json' \
--data-raw '{
    "label": "product-1",
    "type": "type-1",
    "url": "",
    "weight": 54.50
}'
```
```
curl --location --request POST 'http://localhost:8000/api/product' \
--header 'Content-Type: application/json' \
--data-raw '{
    "label": "product-2",
    "type": "type-2",
    "url": "http://someweb.com",
    "weight": 0.00
}'
```

- You can check if the product is on db with this endpoint

```
curl --location --request GET 'http://localhost:8000/api/product/1' \
--header 'Content-Type: application/json' \
```

- Create a shopping cart
```
curl --location --request POST 'http://localhost:8000/api/cart' \
--header 'Content-Type: application/json' \
--data-raw '{

}'
```

- Add some product to the shopping cart
 ``` 
  curl --location --request POST 'http://localhost:8000/api/cart/products/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "cart_id": 1,
    "quantity": 5
  }'
```

 ``` 
  curl --location --request POST 'http://localhost:8000/api/cart/products/2' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "cart_id": 1,
    "quantity": 3
  }'
```

- Get list of product in the shopping cart
```
curl --location --request GET 'http://localhost:8000/api/cart/1' \
--data-raw ''
```

You should see something like this:
```json
[
    {
        "ID": 1,
        "Label": "product-1",
        "Type": "type-1",
        "Url": "",
        "Weight": 54.50,
        "CreatedAt": "2021-11-01T15:56:47.815906Z",
        "UpdatedAt": "2021-11-01T15:56:47.815906Z"
    },
    {
        "ID": 2,
        "Label": "product-2",
        "Type": "type-2",
        "Url": "http://someweb.com",
        "Weight": 0.00,
        "CreatedAt": "2021-11-01T17:34:58.935002Z",
        "UpdatedAt": "2021-11-01T17:34:58.935002Z"
    }
]
```

- Create an order with a cart_id
 ``` 
    curl --location --request POST 'http://localhost:8000/api/order' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "cart_id": 1
    }'
 ``` 
Result:
```json
{
    "id": 67,
    "cart_id": 1,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
}
```

- Get list of product from an order
```
curl --location --request GET 'http://localhost:8000/api/order/67'
 ``` 

You should see something similar to a list of products:
```json
[
    {
        "ID": 1,
        "Label": "product-1",
        "Type": "type-1",
        "Url": "",
        "Weight": 54.5,
        "CreatedAt": "2021-11-09T02:27:59.880887Z",
        "UpdatedAt": "2021-11-09T02:27:59.880887Z"
    },
    {
        "ID": 2,
        "Label": "product-2",
        "Type": "type-2",
        "Url": "http://someweb.com",
        "Weight": 0,
        "CreatedAt": "2021-11-09T02:28:23.451749Z",
        "UpdatedAt": "2021-11-09T02:28:23.451749Z"
    }
]
```
