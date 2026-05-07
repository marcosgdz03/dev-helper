FROM golang:1.24-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o dev-helper

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
RUN addgroup -S devhelper && adduser -S devhelper -G devhelper
COPY --from=builder /build/dev-helper .
USER devhelper
EXPOSE 8080
ENTRYPOINT ["./dev-helper"]
