# Choose whatever you want, version >= 1.16
FROM golang:1.19.5-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy && go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml", "server"]