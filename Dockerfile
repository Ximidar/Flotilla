FROM ubuntu:bionic

COPY . /root/

# Install Resources
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	golang \
	build-essential \
	nano \
	wget \
	git \
	# && update-ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

ENV GOPATH=/root/go
ENV GOBIN=$GOPATH/bin
RUN mkdir $GOPATH 

RUN bash /root/BuildResources/Build/scripts/setupGo.sh