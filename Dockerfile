# ---------- Builder Stage ----------
FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# ---------- Runtime Stage ----------
FROM alfg/nginx-rtmp:latest 

# Install any additional tools you need
#RUN apt-get update && apt-get install -y ffmpeg curl ca-certificates && rm -rf /var/lib/apt/lists/*
RUN apk add --no-cache ffmpeg curl ca-certificates


# Set working directory
WORKDIR /app

# Copy Go app and site assets
COPY --from=builder /app/main /app/main
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

# Replace nginx config
COPY nginx.conf /etc/nginx/nginx.conf
RUN mkdir -p /tmp/hls

# Start script
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Expose ports
EXPOSE 1935 8080 9090

# Healthcheck for Fly
HEALTHCHECK --interval=10s --timeout=3s CMD curl -fs http://localhost:8080/health || exit 1

# Start services
CMD ["/app/start.sh"]
