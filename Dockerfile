FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./src .

RUN go mod download

RUN go build -o main .

CMD ["./main"]