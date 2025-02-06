# Start from the official Golang base image
FROM golang:1.23.4-alpine AS builder

#RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk --no-cache add ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download && go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Run tests before building the Go app
# RUN go test -tags=unit ./internal/test/unit/

# Run integration tests
# RUN go test -tags=integration ./internal/test/integration/

ENV CGO_ENABLED=0 

# Build the Go app
RUN go build -ldflags="-s -w" -v -o /usr/local/bin/app ./

# Empty image
FROM scratch

WORKDIR /app

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

# Command to run the executable
CMD ["/usr/local/bin/app"]