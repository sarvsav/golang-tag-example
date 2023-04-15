FROM golang:1.20

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -v -o /go/bin/app ./...

CMD ["/go/bin/app"]
