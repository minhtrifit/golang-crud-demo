FROM golang:1.23.3-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]