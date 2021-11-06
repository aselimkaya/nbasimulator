FROM golang:1.15
WORKDIR /app
COPY . .
RUN go get github.com/labstack/echo/v4
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go mod tidy
RUN go build -o main src/main.go
CMD ["./main"]