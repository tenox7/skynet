FROM alpine
ARG TARGETARCH
RUN apk add --no-cache busybox-extras
RUN passwd -d root
RUN rm -f /etc/securetty
COPY skynet-${TARGETARCH}-linux /skynet
ENTRYPOINT ["telnetd", "-F", "-p", "4000", "-l", "/skynet"]
