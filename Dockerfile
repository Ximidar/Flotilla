FROM golang:stretch

# Install Resources
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	build-essential \
	nano \
	wget \
	git \
	unzip \
	openssl \
	&& rm -rf /var/lib/apt/lists/*

# Add Home variable for later use
ENV HOME=/home/flotilla

# Setup scripts and variables for setup
ENV NATS_LOC=/nats
RUN mkdir -p $HOME/scripts
RUN mkdir -p $NATS_LOC
COPY ./BuildResources/Build/scripts/* $HOME/scripts/

# Download NATS
# RUN bash $HOME/scripts/setupNats.sh

# Download Go packages
RUN bash $HOME/scripts/setupGo.sh

# Define, make, and populate the Flotilla Directory
ENV FLOTILLA_DIR=$GOPATH/src/github.com/Ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Test
# RUN bash $FLOTILLA_DIR/BuildResources/Test/scripts/test.sh

# Build
WORKDIR $HOME/
RUN exit

