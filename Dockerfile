FROM golang:1.11.1-stretch

# Install Resources
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	build-essential \
	nano \
	wget \
	git \
	unzip \
	&& rm -rf /var/lib/apt/lists/*

ENV FLOTILLA_DIR=$GOPATH/src/github.com/ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/setupGo.sh

# not enough tests and this command doesn't work.
WORKDIR $FLOTILLA_DIR/BuildResources/Test/FlotillaFileManager/
RUN go test ./...

RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

ADD https://github.com/nats-io/gnatsd/releases/download/v1.3.0/gnatsd-v1.3.0-linux-amd64.zip $HOME/nats.zip
RUN unzip $HOME/nats.zip /usr/local/nats/