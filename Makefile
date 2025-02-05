# Makefile in accordance with the docs on git management (to use in combination with meta)
.PHONY: build start clean test

BUILD_DIR=bin/
BINARY_NAME=SERVICE_NAME

lint:
	@echo "Lint check..."
	@golangci-lint run

build: #lint
	@echo "building ${BINARY_NAME}"
	@cd src/ && go build -o "../$(BUILD_DIR)${BINARY_NAME}" ${buildargs}

#
# You can specify run arguments and build arguments using runargs and buildargs, like this:
# make start runargs="-debug"
# make start runargs="-debug" buildargs="-verbose"
# make build buildargs="-verbose"
#
start: build
	@echo "starting ${BINARY_NAME}"
	./${BUILD_DIR}${BINARY_NAME} ${runargs}

clean:
	@echo "Cleaning all targets for ${BINARY_NAME}"
	rm -rf $(BUILD_DIR)

test: lint
	go test ./src -v -count=1 -timeout 0
