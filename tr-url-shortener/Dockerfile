FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
# COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

# Stage 2: Create a minimal runtime image
FROM alpine:latest AS runtime
COPY --from=builder /myapp /myapp

RUN apk update --no-cache && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
# For scratch: COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8000

CMD ["./myapp"]