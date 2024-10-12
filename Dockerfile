FROM golang:1.22.0-alpine

WORKDIR /build

COPY . .


RUN go build -o main cmd/dolg_list/main.go
CMD ["./main"]