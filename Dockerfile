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

# Add Home variable for later use
ENV HOME=/home/flotilla

# Setup scripts and variables for setup
ENV NATS_LOC=$HOME/nats/
RUN mkdir -p $HOME/scripts && mkdir -p $NATS_LOC
COPY ./BuildResources/Build/scripts/* $HOME/scripts/

# Download NATS
RUN bash $HOME/scripts/setupNats.sh

# Download Go packages
RUN bash $HOME/scripts/setupGo.sh

# Define, make, and populate the Flotilla Directory
ENV FLOTILLA_DIR=$GOPATH/src/github.com/ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Test
RUN bash $FLOTILLA_DIR/BuildResources/Test/scripts/test.sh

# Build
WORKDIR $HOME/
RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

CMD bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh

