# Name of output
BINARY=FlotillaWebAPI

# Prefixes
ARM_PREFIX=arm_
ARM64_PREFIX=arm64_
AMD64_PREFIX=amd64_

# Tags for Version, Author, and Date
VERSION=0.0.1
DATE=`date '+%d %b %y at %H:%M:%S %p'`
AUTHOR=Matt Pedler
COMMIT_HASH=`git rev-parse HEAD`

VERSION_PATH=github.com/Ximidar/Flotilla/CommonTools/versioning

# Sources 
SOURCE_DIR=./
OUT_DIR=bin
ENTRY_POINT=${SOURCE_DIR}/FlotillaWebAPI.go
TASK_PID = /tmp/flotillawebapi.pid
FILES=$(shell find $(SOURCE_DIR) -type f \( -iname "*.go" ! -iname "*_test.go" \))

# Web Sources
WEB_ENTRY_POINT=./frontend
WEB_SOURCE=./frontend/src
WEB_FILES=$(shell find $(WEB_SOURCE) -type f \( -iname "*.vue" -or -iname "*.js" -or -iname "*.css" -or -iname "*.html" \))
WEB_DIST=./frontend/dist

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X '${VERSION_PATH}.Version=${VERSION}' -X '${VERSION_PATH}.CompiledBy=${AUTHOR}' -X '${VERSION_PATH}.CompiledDate=${DATE}' -X '${VERSION_PATH}.CommitHash=${COMMIT_HASH}'"

# Setup build vars
ENVARM=env GOOS=linux GOARCH=arm
ENVARM64=env GOOS=linux GOARCH=arm64
ENVAMD64=env GOOS=linux GOARCH=amd64
GO_BUILD=go build

# Builds the project
build_all:
	@echo Building ${BINARY} Binary for this machine
	${GO_BUILD} ${LDFLAGS} -o ${OUT_DIR}/${BINARY} ${ENTRY_POINT}

	@echo
	@echo Building ${ARM_PREFIX}${BINARY} Binary for ARM32
	${ENVARM} ${GO_BUILD} ${LDFLAGS} -o ${OUT_DIR}/${ARM_PREFIX}${BINARY} ${ENTRY_POINT}

	@echo
	@echo Building ${ARM64_PREFIX}${BINARY} Binary for ARM64
	${ENVARM64} ${GO_BUILD} ${LDFLAGS} -o ${OUT_DIR}/${ARM64_PREFIX}${BINARY} ${ENTRY_POINT}

	@echo
	@echo Building ${AMD64_PREFIX}${BINARY} Binary for AMD64
	${ENVAMD64} ${GO_BUILD} ${LDFLAGS} -o ${OUT_DIR}/${AMD64_PREFIX}${BINARY} ${ENTRY_POINT}

	@echo
	@echo All Binaries are built to the ${OUT_DIR} Folder

build:
	@echo Building ${BINARY} Binary for this machine
	${GO_BUILD} ${LDFLAGS} -o ${OUT_DIR}/${BINARY} ${ENTRY_POINT}
	@echo
	@echo $(BINARY) built in ${OUT_DIR} Folder

	@echo Building Web Distribution
	cd ${WEB_ENTRY_POINT}; \
		npm run develop;

	if [ -d "$(OUT_DIR)/dist" ]; then \
		echo "Deleting previous run"; \
		rm -rf $(OUT_DIR)/dist; \
	fi
	cp -r ${WEB_DIST} ${OUT_DIR}


clean:
	@echo
	@echo Cleaning the $(OUT_DIR) Folder
	rm -rf $(OUT_DIR)


start: build
	${OUT_DIR}/${BINARY} & echo $$! > $(TASK_PID)

stop: 
	-kill `pstree -p \`cat $(TASK_PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

# Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
before:
	@echo "STOPED $(BINARY)" && printf '%*s\n' "40" '' | tr ' ' -

# Restart task will execute stop, before and start tasks in strict order and prints message. 
restart: stop before start
	@echo "STARTED $(BINARY)" && printf '%*s\n' "40" '' | tr ' ' -

# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
serve: start
	fswatch -or --event=Updated $(FILES) $(WEB_FILES) | \
	xargs -n1 -I {} make restart

.PHONY: start before stop restart serve build build_all clean