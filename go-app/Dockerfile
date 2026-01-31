FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -v -o server

# Use a minimal base image for the final container
FROM gcr.io/distroless/base-debian11
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080 (Cloud Run default)
EXPOSE 8080

# Run the Go app
CMD ["/app/server"]
