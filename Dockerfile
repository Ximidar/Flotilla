FROM golang:latest

# Install basic Resources
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	build-essential \
	nano \
	wget \
	git \
	unzip \
	openssl \
	fswatch \
	&& rm -rf /var/lib/apt/lists/*


# Add Home variable for later use
ENV HOME=/home/flotilla

# Setup scripts and variables for setup
RUN mkdir -p $HOME
WORKDIR $HOME

# Define, make, and populate the Flotilla Directory
ENV FLOTILLA_DIR=$HOME/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Build
WORKDIR $FLOTILLA_DIR
RUN go mod download


