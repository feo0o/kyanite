FROM debian:12

RUN apt-get update && \
    apt-get install -y --no-install-recommends ethtool netcat-traditional procps apt-file curl tcpdump bind9-dnsutils telnet file net-tools iproute2 inetutils-ping  && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
COPY release/kyanite /
ENTRYPOINT ["/kyanite"]