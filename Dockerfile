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
	npm \
	&& rm -rf /var/lib/apt/lists/*


# Add Home variable for later use
ENV HOME=/home/flotilla

# Setup scripts and variables for setup
RUN mkdir -p $HOME
WORKDIR $HOME

# Define, make, and populate the Flotilla Directory
ENV FLOTILLA_DIR=$HOME
ENV GIT_DISCOVERY_ACROSS_FILESYSTEM=1
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

# Build
WORKDIR $FLOTILLA_DIR
RUN go mod tidy
RUN make build

#TODO start a second lightweight image here that lifts the executable out of the first image

ENTRYPOINT [ "/home/flotilla/bin/flot" ]


