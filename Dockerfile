FROM golang:1.18.4-alpine3.16

WORKDIR /src

COPY . .

RUN apk add git

RUN go mod download

CMD ["go", "run", "main.go"]