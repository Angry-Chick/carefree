FROM golang:latest

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/carefree

ADD . .

RUN go build ./project/door/frontend

EXPOSE 8000

ENTRYPOINT [ "./frontend" ]
