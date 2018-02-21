FROM golang:1.8.5-jessie as builder
LABEL maintainer "Helder Farias <helderfarias@gmail.com>"

RUN apt-get update && apt-get install -y xz-utils && rm -rf /var/lib/apt/lists/*
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

WORKDIR /root
ADD templates templates
ADD release release
RUN strip --strip-unneeded /root/release/alpine/dynamock
RUN upx /root/release/alpine/dynamock


FROM alpine:3.5
LABEL maintainer "Helder Farias <helderfarias@gmail.com>"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /root/release/alpine/dynamock /bin/dynamock
COPY --from=builder /root/templates /etc/dynamock
RUN chmod +x /bin/dynamock
VOLUME /etc/dynamock
EXPOSE 3010

ENTRYPOINT ["/bin/dynamock"]

CMD ["-c", "/etc/dynamock/config.json"]
