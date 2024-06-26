FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY internal ./internal
COPY frontend/build ./frontend/build

RUN go build -o magic-kingdom

EXPOSE 8080

CMD ["./magic-kingdom"]
