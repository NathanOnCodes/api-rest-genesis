FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-rest-genesis
RUN chmod +x api-rest-genesis

EXPOSE 8000

CMD ["./api-rest-genesis"]