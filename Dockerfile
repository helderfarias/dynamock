FROM golang

MAINTAINER Helder Farias <helderfarias@gmail.com>

RUN curl -L https://github.com/Masterminds/glide/releases/download/0.10.2/glide-0.10.2-linux-amd64.tar.gz > glide.tar.gz \
    && tar xvzf glide.tar.gz \
    && cp linux-amd64/glide /usr/bin/glide \
    && rm -rf linux-amd64

RUN mkdir -p cd $GOPATH/src/github.com/helderfarias/dynamock
COPY cli $GOPATH/src/github.com/helderfarias/dynamock/cli
COPY main.go $GOPATH/src/github.com/helderfarias/dynamock/
COPY glide.yaml $GOPATH/src/github.com/helderfarias/dynamock/glide.yaml

RUN cd src/github.com/helderfarias/dynamock \
    && glide install \
    && go build

RUN cd src/github.com/helderfarias/dynamock \
    && cp dynamock /usr/bin/dynamock \
    && chmod +x /usr/bin/dynamock \
    && rm -rf /go

COPY entrypoint.sh /entrypoint.sh
COPY templates /templates
RUN chmod +x /entrypoint.sh

VOLUME /templates

EXPOSE 3010

CMD "/entrypoint.sh"

