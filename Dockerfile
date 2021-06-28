FROM golang:latest

LABEL maintainer="Tiago Martins <tiago.martins180@gmail.com>"

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod . 
COPY go.sum . 
RUN go mod download

COPY . . 

ENV PORT 8000

RUN go build

CMD ["./WebChatBackend"]