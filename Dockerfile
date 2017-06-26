FROM alpine:3.5

MAINTAINER Helder Farias <helderfarias@gmail.com>

COPY dynamock_alpine /usr/bin/dynamock
RUN chmod +x /usr/bin/dynamock

COPY templates /etc/config

VOLUME /config

EXPOSE 3010

ENTRYPOINT ["/usr/bin/dynamock"]

CMD ["-c", "/etc/config/sample.json"]
