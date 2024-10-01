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
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   ├── sqlx_setup.go
│   │   │   └── migrate.go
│   │   ├── Dockerfile
│   │   └── main.go
│   ├── order_service/
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   ├── sqlx_setup.go
│   │   │   └── migrate.go
│   │   ├── Dockerfile
│   │   └── main.go
├── docker-compose.yml
├── init/
├── Makefile
├── go.mod
├── go.sum
├── wait-for-it.sh
└── .env
```