FROM golang
WORKDIR /go/src/github.com/jonathan/Go-naive-user-service
COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN cd /go/src/github.com/jonathan/Go-naive-user-service && go build
EXPOSE 8080

CMD ["./Go-naive-user-service"]

