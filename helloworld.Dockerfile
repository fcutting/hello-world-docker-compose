FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY ./cmd ./
RUN go build main.go
CMD ["./main"]