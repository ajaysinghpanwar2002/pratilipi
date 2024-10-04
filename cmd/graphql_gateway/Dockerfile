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
RUN go build -o graphql_gateway ./cmd/graphql_gateway

# Stage 2: Run the Go binary
FROM alpine:latest

# Set up timezone if needed (optional)
RUN apk add --no-cache bash tzdata

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/graphql_gateway /app/graphql_gateway

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./graphql_gateway"]
