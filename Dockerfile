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

ENV ENVIRONMENT=production

# Build the Go app
RUN go build -ldflags="-s -w" -v -o /usr/local/bin/app ./

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates curl

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the Casbin configuration file
COPY --from=builder /app/pkg/authorization/rbac_model.conf /app/pkg/authorization/rbac_model.conf

COPY --from=builder /app/pkg/authorization/policy.csv /app/pkg/authorization/policy.csv

COPY --from=builder /app/Invite_email_template.html /app/Invite_email_template.html

ENV ENVIRONMENT=production

EXPOSE 8080

# Command to run the executable
CMD ["/usr/local/bin/app"]