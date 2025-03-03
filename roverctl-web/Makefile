# Makefile in accordance with the docs on git management (to use in combination with meta)
.PHONY: build start clean test

build-open-api:
	# Check if the API spec file exists
	@if [ ! -f ../spec/api/schema.yaml ]; then \
		echo "../spec/api/schema.yaml file not found."; \
		exit 1; \
	fi
	@echo "generating openapi client"
	@openapi-generator-cli generate -i ../spec/api/schema.yaml -g typescript-axios -o src/lib/openapi 

build: #build-open-api
	@ echo "Not implemented yet"
	# @if [ -z ${VERSION} ]; then \
	# 	echo "VERSION is not set. Please set it to the semantic version you want to build."; \
	# 	exit 1; \
	# fi
	# @echo "building ${BINARY_NAME}"
	# @cd src/ && go build -buildvcs=false -ldflags="$(LDFLAGS)" -o "../$(BUILD_DIR)${BINARY_NAME}" ${buildargs}


#
# You can specify run arguments and build arguments using runargs and buildargs, like this:
# make start runargs="-debug"
# make start runargs="-debug" buildargs="-verbose"
# make build buildargs="-verbose"
#
start: build
	@ echo "Not implemented yet"
	# @echo "starting ${BINARY_NAME}"
	# ./${BUILD_DIR}${BINARY_NAME} ${runargs}

clean:
	@ echo "Not implemented yet"
	# @echo "Cleaning all targets for ${BINARY_NAME}"
	# rm -rf $(BUILD_DIR)
	# rm -rf src/openapi
	# go mod tidy

test:
	@ echo "Not implemented yet"
	# go test ./src -v -count=1 -timeout 0
