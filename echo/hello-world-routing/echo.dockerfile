FROM golang:1.12

ADD ./ /web/src/hello-world-routing
WORKDIR /web

RUN mkdir bin
ENV GOPATH /web
ENV PATH ${GOPATH}/bin:${PATH}

RUN curl -sL -o install_glide.sh https://glide.sh/get
RUN sh install_glide.sh

WORKDIR src/hello-world-routing
RUN glide install
WORKDIR ..
RUN go install hello-world-routing
CMD hello-world-routing