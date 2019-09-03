FROM golang:1.12

ADD ./ /web/src/hello-world
WORKDIR /web

RUN mkdir bin
ENV GOPATH /web
ENV PATH ${GOPATH}/bin:${PATH}

RUN curl -sL -o install_glide.sh https://glide.sh/get
RUN sh install_glide.sh

WORKDIR src/hello-world
RUN glide install

RUN go build -o app main.go
CMD ./app