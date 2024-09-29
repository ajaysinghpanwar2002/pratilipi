# Define service names for docker-compose
COMPOSE_FILE = docker-compose.yml

# Create .env file if it doesn't exist
create-env:
	@if [ ! -f .env ]; then \
		echo "Creating .env file from .env.example..."; \
		cp .env.example .env; \
	else \
		echo ".env file already exists."; \
	fi

# Start all the services (Postgres, Redis, RabbitMQ) after ensuring .env file exists
start: create-env
	@echo "Starting Docker containers..."
	docker-compose -f $(COMPOSE_FILE) up -d

# Stop all the services
stop:
	@echo "Stopping Docker containers..."
	docker-compose -f $(COMPOSE_FILE) down

# Rebuild the services if Dockerfile changes
build:
	@echo "Building Docker containers..."
	docker-compose -f $(COMPOSE_FILE) up --build -d

# Clean up and remove containers, volumes, and networks
clean:
	@echo "Cleaning up Docker environment..."
	docker-compose -f $(COMPOSE_FILE) down -v

# Flush Redis cache
cache-flush:
	@echo "Flushing Redis cache..."
	docker exec -it `docker ps -qf "name=redis"` redis-cli FLUSHALL

# Check PostgreSQL status
db-status:
	@echo "Checking PostgreSQL status..."
	docker exec -it `docker ps -qf "name=postgres"` psql -U your_user -d your_db -c '\l'

# Running User Service (replace 'cmd/user_service/main.go' with actual path)
run-user-service:
	@echo "Running User Service..."
	go run ./cmd/user_service/main.go

# Running Product Service (replace 'cmd/product_service/main.go' with actual path)
run-product-service:
	@echo "Running Product Service..."
	go run ./cmd/product_service/main.go

# Running Order Service (replace 'cmd/order_service/main.go' with actual path)
run-order-service:
	@echo "Running Order Service..."
	go run ./cmd/order_service/main.go
