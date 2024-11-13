# Build
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.22-alpine3.19 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY . /app
RUN go mod tidy && GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app main.go

# Release
FROM alpine:3.19

RUN apk add --update words-en && \
    ln -s /usr/share/dict/american-english /usr/share/dict/words
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
