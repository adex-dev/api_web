FROM golang:1.21.6-alpine
WORKDIR /app

COPY . .

RUN go build -o apiweb

EXPOSE 8088

CMD ./apiweb