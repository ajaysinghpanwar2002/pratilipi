```bash
pratilipi/
├── cmd/
    |──user-service
        ├── db/
          ├── migrations/
              ├── 001_create_users_table.up.sql
                 └── 001_create_users_table.down.sql
        |──internal/
                ├── handlers/
                │   └── user_handler.go
                ├── repositories/
                │   └── user_repository.go
                ├── services/
                │   └── user_service.go
                ├── middlewares/
                │   └── jwt_middleware.go
                ├── models/
                │   └── user.go
        ├── main.go
        ├── Dockerfile
    |──product-service
        ├── db/
          ├── migrations/
              ├── 001_create_product_table.up.sql
                 └── 001_create_product_table.down.sql
        |──internal/
                ├── handlers/
                │   └── product_handler.go
                ├── repositories/
                │   └── product_repository.go
                ├── services/
                │   └── product_service.go
                ├── models/
                │   └── product.go
        ├── main.go
        ├── Dockerfile
    |──order-service
        ├── db/
          ├── migrations/
              ├── 001_create_order_table.up.sql
                 └── 001_create_order_table.down.sql
        |──internal/
                ├── handlers/
                │   └── order_handler.go
                ├── repositories/
                │   └── order_repository.go
                ├── services/
                │   └── order_service.go
                ├── models/
                │   └── order.go
        ├── main.go
        ├── Dockerfile
├── pkg
    ├── db
        ├── init_db
        ├── migrate.go
        ├── sqlx_setup.go
    ├── rabbitmq
        ├── connection.go
        ├── events.go
├── docker-compose.yml
├── init/
├── Makefile
├── go.mod
├── go.sum
├── wait-for-it.sh
└── .env
```