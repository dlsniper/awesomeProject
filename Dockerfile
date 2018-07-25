FROM jetbrains/teamcity-minimal-agent:latest

MAINTAINER Florin Patan <florin@jetbrains.com>

LABEL dockerImage.teamcity.version="latest" \
      dockerImage.teamcity.buildNumber="latest"

ENV GIT_SSH_VARIANT=ssh

RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository -y ppa:openjdk-r/ppa && add-apt-repository -y ppa:git-core/ppa && apt-get update && \
    apt-get install -y git mercurial apt-transport-https ca-certificates && \
    \
    apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 9DC858229FC7DD38854AE2D88D81803C0EBFCD88 && \
    echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable" > /etc/apt/sources.list.d/docker.list && \
    \
    apt-cache policy docker-ce && \
    apt-get update && \
    apt-get install -y docker-ce=18.03.1~ce-0~ubuntu && \
    systemctl disable docker && \
    curl -SL https://github.com/docker/compose/releases/download/1.21.2/docker-compose-Linux-x86_64 -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose && \
    \
#    apt-get install -y --no-install-recommends \
#            libc6 \
#            libcurl3 \
#            libgcc1 \
#            libgssapi-krb5-2 \
#            libicu55 \
#            liblttng-ust0 \
#            libssl1.0.0 \
#            libstdc++6 \
#            libunwind8 \
#            libuuid1 \
#            zlib1g \
#    && \
    rm -rf /var/lib/apt/lists/* && \
    \
    usermod -aG docker buildagent

COPY run-docker.sh /services/run-docker.sh