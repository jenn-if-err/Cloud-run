FROM golang:1.24.5-alpine AS builder


WORKDIR /app


# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build a statically linked Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server .

# Use a minimal base image for the final container
FROM gcr.io/distroless/static-debian11
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080 (Cloud Run default)
EXPOSE 8080

# Run the Go app
CMD ["/app/server"]
