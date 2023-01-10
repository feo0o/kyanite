FROM docker.io/alpine:3.16.3
RUN apk add --update --no-cache tzdata curl busybox busybox-extras bind-tools tcpdump strace ltrace tree iperf iperf3 && \
    rm -rf /var/cache/apk/* && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
COPY release/kyanite_linux_amd64 /kyanite
ENTRYPOINT ["/kyanite"]


