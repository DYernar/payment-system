FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o app ./cmd/app/*.go

EXPOSE 8111

ENV HTTP_PORT=8111

CMD ["./app"]