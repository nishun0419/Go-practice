FROM golang:latest

COPY . /go/src

WORKDIR /go/src/

ENV GO111MODULE=on
RUN go mod download

CMD ["go", "run",  "main.go"]