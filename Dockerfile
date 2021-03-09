FROM golang:1.16.2

WORKDIR /go/src/app
COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000
