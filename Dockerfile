# ================================
# PRODUCTION DOCKERFILE
# Multi-stage build for security & performance
# ================================

# --------------------------------
# Build Stage
# --------------------------------
FROM golang:1.22-alpine AS builder

# Security: Install ca-certificates and create non-root user
RUN apk add --no-cache \
    ca-certificates \
    git \
    tzdata \
    && update-ca-certificates

# Create appuser for security
RUN adduser -D -g '' -u 10001 appuser

# Set working directory
WORKDIR /build

# Copy dependency files first (better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build optimized static binary
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -tags netgo \
    -o app \
    main.go

# --------------------------------  
# Production Stage (Minimal)
# --------------------------------
FROM scratch

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy CA certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy user and group files
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /build/app /app

# Use non-root user for security
USER appuser

# Expose port
EXPOSE 3000

# Add metadata labels
LABEL maintainer="Your Name <your.email@example.com>"
LABEL version="1.0"
LABEL description="Dinsos API - Production Ready"

# Health check
HEALTHCHECK --interval=30s \
    --timeout=10s \
    --start-period=5s \
    --retries=3 \
    CMD ["/app", "--health-check"] || exit 1

# Run the binary
ENTRYPOINT ["/app"]