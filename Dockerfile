FROM golang:1.20-alpine

WORKDIR /app
COPY . /app
RUN go mod tidy && go build -o app main.go

EXPOSE 8080
ENTRYPOINT ["./app"]
