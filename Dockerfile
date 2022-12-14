FROM golang:1.19.1-alpine

RUN apk update && apk add git
RUN apk add curl
#よくビルドに必要なるものをまとめたパック（alpineは軽量化のために色々省いているので、そのままだとgo testが動かない）
RUN apk add build-base

#mysql install （これしないと mysqlコマンドが go app側から使えない）
RUN apk add mysql-client

RUN go install -v golang.org/x/tools/gopls@latest \
	&& go install -v github.com/ramya-rao-a/go-outline@latest \
	&& go install -v golang.org/x/tools/cmd/goimports@latest \
	&& go install -v github.com/mdempsky/gocode@latest

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY go.mod ./
RUN go mod download && go mod verify
