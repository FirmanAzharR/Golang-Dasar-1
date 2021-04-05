FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN get https://github.com/FirmanAzharR/Golang
RUN cd /build && git clone https://github.com/FirmanAzharR/Golang.git

RUN cd /build/Golang/ && go build

EXPOSE 8080

ENTRYPOINT ["/build/Golang"]