FROM golang:1.22-alpine

# Install Docker client and build dependencies
RUN apk add --no-cache docker-cli build-base

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o app

ENTRYPOINT ["/app/app"] 