```bash
pratilipi/
├── cmd/
│   ├── user_service/
|   |   ├── handlers/
|   |   │   └── user_handler.go
|   |   ├── repositories/
|   |   │   └── user_repository.go
|   |   ├── services/
|   |   │   └── user_service.go
|   |   ├── middlewares/
|   |   │   └── jwt_middleware.go
|   |   ├── models/
|   |   │   └── user.go
|   |   ├── db/
|   |   │   ├── migrations/
|   |   │   │   ├── 001_create_users_table.up.sql
|   |   │   │   └── 001_create_users_table.down.sql
|   |   │   └── sqlx_setup.go
|   |   ├── main.go
│   ├── product_service/
│   │   ├── handlers/
│   │   │   └── product_handler.go
│   │   ├── repositories/
│   │   │   └── product_repository.go
│   │   ├── services/
│   │   │   └── product_service.go
│   │   ├── models/
│   │   │   └── product.go
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   ├── 001_create_product_table.up.sql
│   │   │   │   └── 001_create_product_table.down.sql
│   │   ├── main.go
│   │   ├── Dockerfile
│   ├── order_service/
│   │   ├── handlers/
│   │   │   └── order_handler.go
│   │   ├── repositories/
│   │   │   └── order_repository.go
│   │   ├── services/
│   │   │   └── order_service.go
│   │   ├── models/
│   │   │   └── order.go
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   ├── 001_create_orders_table.up.sql
│   │   │   │   └── 001_create_orders_table.down.sql
│   │   ├── main.go
│   │   ├── Dockerfile
├── docker-compose.yml
├── init/
├── Makefile
├── go.mod
├── go.sum
├── wait-for-it.sh
└── .env
```