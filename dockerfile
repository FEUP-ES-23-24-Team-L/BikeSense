# Use alpine base for minimal size
FROM golang:alpine AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Install dependencies (replace with your package manager if needed)
RUN go mod download

# Build the application (multi-stage for smaller final image)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/bikesense-web.go

FROM alpine:3.19.1

COPY --from=builder /app/app /app

# Expose port
EXPOSE 8080

# Start the application
CMD [ "/app" ]
