### API Endpoints

---

#### 1. **Register User**

- **Method**: `POST`
- **URL**: `http://localhost:8081/register`

- **Request Body**:

```json
{
  "username": "ajaysinghpanwar2002",
  "email": "ajaysinghpanwar2002@gmail.com",
  "password": "ajaysinghpanwar2002pass"
}
```

- **Expected Response**:

```json
{
  "message": "User registered successfully",
  "user_id": "b091999c-319d-4e88-b9d7-d98505064a56"
}
```

---

#### 2. **Login**

- **Method**: `POST`
- **URL**: `http://localhost:8081/login`

- **Request Body**:

```json
{
  "username": "ajaysinghpanwar2002",
  "password": "ajaysinghpanwar2002pass"
}
```

- **Expected Response**:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgwNTQ1ODcsInVzZXJfaWQiOiJiMDkxOTk5Yy0zMTlkLTRlODgtYjlkNy1kOTg1MDUwNjRhNTYifQ.dUhsKGkJ_S3460PxpAceNcnSMSVUsA5Z3VLqgLgeO1w"
}
```

---

#### 3. **Update Profile**

- **Method**: `PUT`
- **URL**: `http://localhost:8081/profile`

- **Headers**:
  - `Authorization: <token>`

- **Request Body**:

```json
{
  "email": "heyajaysingh@gmail.com"
}
```

- **Expected Response**:

```json
{
  "message": "Profile updated successfully"
}
```

---

#### 4. **Create Product**

- **Method**: `POST`
- **URL**: `http://localhost:8082/products`

- **Request Body**:

```json
{
  "name": "yjtyj classic 350",
  "description": "bhrthtrike",
  "price": 13200,
  "stock": 150
}
```

- **Expected Response**:

```json
{
  "message": "Product created successfully",
  "product_id": "795a75e9-6d06-47c6-b868-4e2ac104adb6"
}
```

---

#### 5. **Get Product**

- **Method**: `GET`
- **URL**: `http://localhost:8082/products/{product_id}`

- **Example Response**:

```json
{
  "ID": "39d55f47-7620-4505-8f39-0e5714d73d87",
  "Name": "classic 350",
  "Description": "bike",
  "Price": 1403200,
  "Stock": 10350,
  "CreatedAt": "2024-10-03T15:11:09.408812Z",
  "UpdatedAt": "2024-10-03T15:11:09.408812Z"
}
```

---

#### 6. **Update Product**

- **Method**: `PUT`
- **URL**: `http://localhost:8082/products/{product_id}`

- **Request Body**:

```json
{
  "stock": 400
}
```

- **Expected Response**:

```
updated product
```

---

#### 7. **Delete Product**

- **Method**: `DELETE`
- **URL**: `http://localhost:8082/products/{product_id}`

- **Expected Response**:

```
Product deleted
```

---

#### 8. **Place Order**

- **Method**: `POST`
- **URL**: `http://localhost:8083/orders`

- **Request Body**:

```json
{
  "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
  "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
  "quantity": 10
}
```

- **Expected Response**:

```json
{
  "id": "98ff0676-2b67-407f-9d85-0b3fa791f121",
  "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
  "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
  "quantity": 10,
  "status": "Placed",
  "total_price": 14032000,
  "created_at": "2024-10-03T15:16:18.725453178Z",
  "updated_at": "2024-10-03T15:16:18.725453228Z"
}
```

---

#### 9. **Get All Orders**

- **Method**: `GET`
- **URL**: `http://localhost:8083/orders`

- **Example Response**:

```json
[
  {
    "id": "98ff0676-2b67-407f-9d85-0b3fa791f121",
    "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
    "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
    "quantity": 10,
    "status": "Placed",
    "total_price": 14032000,
    "created_at": "2024-10-03T15:16:18.725453178Z",
    "updated_at": "2024-10-03T15:16:18.725453228Z"
  },
  {
    "id": "6a3e6219-8bbf-41d6-98ff-a5ab0c1dd1b8",
    "user_id": "e2c5d846-ef19-4014-9a6f-5a6f71af6861",
    "product_id": "d9bbfbc4-5c1e-4764-941d-52f61cb53093",
    "quantity": 5,
    "status": "Shipped",
    "total_price": 250000,
    "created_at": "2024-10-02T10:22:45.725453178Z",
    "updated_at": "2024-10-02T10:22:45.725453228Z"
  }
]
```

---