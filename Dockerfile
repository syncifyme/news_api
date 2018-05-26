FROM golang:1.9.2

WORKDIR /goworks/src/github.com/syncifyme/news_api

# Setting up environment variables
ENV GO_ENV docker
ENV GOPATH /goworks

# RUN go get github.com/tools/godep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

EXPOSE 3000

ENTRYPOINT ["go", "run", "main.go"]