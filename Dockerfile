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

# Add the flotilla user
RUN adduser --disabled-password --gecos '' --home /home/flotilla -u 10000 flotilla
RUN chown flotilla $GOPATH
USER flotilla 

# Define and make the Flotilla Directory
ENV FLOTILLA_DIR=$GOPATH/src/github.com/ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Download NATS
ENV NATS_LOC=$HOME/nats
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/setupNats.sh

# Download Go packages
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/setupGo.sh

# Test
RUN bash $FLOTILLA_DIR/BuildResources/Test/scripts/test.sh

# Build
WORKDIR $HOME/
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

CMD bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

