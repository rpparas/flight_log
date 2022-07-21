FROM golang:1.18-alpine

RUN apk add --no-cache git

WORKDIR /go/src/app
ENV CGO_ENABLED=0

COPY . .

RUN go get -d -v
RUN go install -v .
RUN go mod tidy

RUN go build -o server main.go

ENTRYPOINT [ "./server" ]
