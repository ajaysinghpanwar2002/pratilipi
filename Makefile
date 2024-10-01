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

# Rebuild and restart a specific service
rebuild-service:
	@read -p "Enter service name (user_service, product_service, order_service): " SERVICE; \
	echo "Rebuilding and restarting $$SERVICE..."; \
	docker-compose -f $(COMPOSE_FILE) up -d --no-deps --build $$SERVICE

# View logs of a specific service
logs-service:
	@read -p "Enter service name (user_service, product_service, order_service): " SERVICE; \
	echo "Viewing logs of $$SERVICE..."; \
	docker-compose -f $(COMPOSE_FILE) logs -f $$SERVICE