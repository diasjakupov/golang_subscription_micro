# Start from the official Golang image.
FROM golang:1.23-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy migration files

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application, disable CGO to create a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app cmd/app/main.go

# Use a smaller image to run the app
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN ls -la

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/demo-app .

COPY .env .env

EXPOSE 5050

# Command to run the executable
CMD ["./demo-app"]