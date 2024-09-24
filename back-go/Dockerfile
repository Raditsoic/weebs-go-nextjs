FROM golang:1.22.5-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Download Dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o weebs-server .

# Create a new image from scratch
FROM alpine:latest

# Copy the binary from the builder image
WORKDIR /app
COPY --from=builder /app/weebs-server .

EXPOSE 8080

CMD ["./weebs-server"]

