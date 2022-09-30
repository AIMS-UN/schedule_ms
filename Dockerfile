FROM golang:alpine as builder

WORKDIR /app .

COPY ./go.mod go.sum /

RUN go mod download && go mod verify

COPY . .

RUN go build -o main main.go

CMD ["./main"]