# Aplikasi E-commerce Synapsys

## Rendy Eka F

## Panduan Installasi

1. Download source code
2. Buka terminal dan jalankan perintah: `go run main.go`
3. Buka file postman yang ada di aplikasi Postman
4. Jalankan login untuk mendapatkan JWT token
5. Setelah itu, masukkan JWT token di authorization pada API selain register dan login

## Sistem yang Dibutuhkan

1. Golang 1.19
2. PostgreSQL
3. Postman

## Authentication

1. **"/register"**
    - **Method**: POST
    - **Request**: 
      ```json
      { 
        "name": "string", 
        "email": "string", 
        "password": "string" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Register successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 409: 
         ```json
         {"status": 409, "data" : nil,  "message": "Email already registered"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to register"}
         ```

2. **"/login"**
    - **Method**: POST
    - **Request**:
      ```json
      { 
        "name": "string", 
        "email": "string", 
        "password": "string" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Login successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to login"}
         ```

## Products

1. **"/products/category"**
    - **Method**: POST
    - **Request**:
      ```json
      { 
        "category": "string" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Get Product successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to get products list"}
         ```

## Shopping Cart

1. **"/cart/add"**
    - **Method**: POST
    - **Request**:
      ```json
      { 
        "user_id": "int", 
        "product_id": "int", 
        "quantity": "int minimal 8" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Add cart successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to add cart list"}
         ```

2. **"/cart/get-cart"**
    - **Method**: POST
    - **Request**:
      ```json
      { 
        "user_id": "int" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Get Cart Successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to get cart list"}
         ```

3. **"/cart/delete"**
    - **Method**: PUT
    - **Request**:
      ```json
      { 
        "user_id": "int", 
        "cart_id": "int" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Delete Cart Successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to delete cart list"}
         ```

## Checkout

1. **"/order/new-order"**
    - **Method**: POST
    - **Request**:
      ```json
      { 
        "user_id": "int" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Add order Successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "No order found"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Invalid Input"}
         ```
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to add order"}
         ```

1. **"/order/checkout"**
    - **Method**: PUT
    - **Request**:
      ```json
      { 
        "method": "string",
        "order_Id": "int" 
      }
      ```
    - **Response**:
       - 200: 
         ```json
         {"status": 200, "data" : "name",  "message": "Payment order Successful"}
         ```
       - 400: 
         ```json
         {"status": 400, "data" : nil,  "message": "Failed to payment order"}
       - 401: 
         ```json
         {"status": 401, "data" : nil,  "message": "Invalid credentials"}
         ```
       - 500: 
         ```json
         {"status": 500, "data" : nil,  "message": "Failed to add order"}
         ```