RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list && \
    apt update && \
    apt-get install -y --no-install-recommends apt-file curl tcpdump bind9-dnsutils telnet file net-tools iproute2 inetutils-ping  && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
COPY kyanite /kyanite
ENTRYPOINT ["/kyanite"]


