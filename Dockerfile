FROM golang:1.11.1-stretch

# Install Resources
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	build-essential \
	nano \
	wget \
	git \
	&& rm -rf /var/lib/apt/lists/*

ENV FLOTILLA_DIR=$GOPATH/src/github.com/ximidar/Flotilla/
RUN mkdir -p $FLOTILLA_DIR
COPY . $FLOTILLA_DIR

RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/setupGo.sh

# not enough tests and this command doesn't work.
#RUN go test $FLOTILLA_DIR/BuildResources/Test/...

RUN bash $FLOTILLA_DIR/BuildResources/Build/scripts/buildFlotilla.sh