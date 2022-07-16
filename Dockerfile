FROM golang:1.18.4-alpine3.16

COPY . /src

WORKDIR /src

RUN go mod download

RUN apk add git

RUN go build -o main .

EXPOSE ${APP_PORT}

CMD ./main