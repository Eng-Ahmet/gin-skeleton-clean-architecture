# Use official Golang image as base
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install git for go mod tidy
RUN apk add --no-cache git

# Copy Go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy || cat /root/.cache/go-build/log.txt

# Copy source code
COPY . .

# Build the Go app
RUN go build -o fiber-clean-skelton .

# Use a smaller image to run the app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app ./

EXPOSE 8080

CMD ["./fiber-clean-skelton"]
