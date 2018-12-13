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

# Download NATS
RUN wget https://github.com/nats-io/gnatsd/releases/download/v1.3.0/gnatsd-v1.3.0-linux-amd64.zip -O $HOME/nats.zip
RUN unzip $HOME/nats.zip -d /usr/local/ && mv /usr/local/gnatsd-v1.3.0-linux-amd64/ /usr/local/nats
ENV PATH=$PATH:/usr/local/nats/

# Get Flotilla Source
ENV FLOTILLA_DIR=$GOPATH/src/github.com/ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Download Go packages
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/setupGo.sh

# Test
RUN bash $FLOTILLA_DIR/BuildResources/Test/scripts/test.sh

# Build
WORKDIR $HOME/
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

CMD bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

