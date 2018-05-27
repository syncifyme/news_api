FROM golang:1.9.2


WORKDIR /go/src/github.com/syncifyme/news_api

# Setting up environment variables
ENV GO_ENV docker
ENV GOPATH /go

# RUN go get github.com/tools/godep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

ADD . /go/src/github.com/syncifyme/news_api

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]