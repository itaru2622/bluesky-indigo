# Stage 1: Build the Go binary
FROM golang:1.24-alpine3.22 AS build-env

# Install SSL ca certificates.
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# Create a directory for the application
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
ENV GOEXPERIMENT=loopvar
RUN CGO_ENABLED=0 GOOS=linux go build -o /sonar ./cmd/sonar

# Stage 2: Build a minimal Docker image
FROM scratch

# Import the SSL certificates from the first stage.
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the binary from the first stage.
COPY --from=build-env /sonar /sonar

# Set the startup command to run the binary
CMD ["/sonar"]
