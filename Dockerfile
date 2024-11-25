FROM golang:1.22.8-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o library-category-service

RUN chmod +x library-category-service

EXPOSE 9091

EXPOSE 6001

CMD ["./library-category-service"]
