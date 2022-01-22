FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v . && \
    go build -v .

CMD ["./schablone-server -v"]
