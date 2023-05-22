# Build
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.20-alpine AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY . /app
RUN go mod tidy && GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app main.go

# Release
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
