# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app binary
RUN go build -o order_service ./cmd/order_service

# Stage 2: Run the Go binary
FROM alpine:latest

# Set up timezone if needed (optional)
RUN apk add --no-cache tzdata

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/order_service /app/order_service

# Copy the wait-for-it script
COPY wait-for-it.sh /app/wait-for-it.sh

# Make the wait-for-it script executable
RUN chmod +x /app/wait-for-it.sh

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./wait-for-it.sh", "postgres:5432", "--", "./order_service"]