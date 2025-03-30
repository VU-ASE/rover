.PHONY: build start clean test

build-open-api:
	# Check if the API spec file exists
	@if [ ! -f ../spec/api/schema.yaml ]; then \
		echo "../spec/api/schema.yaml file not found."; \
		exit 1; \
	fi
	@echo "generating openapi client"
	@openapi-generator-cli generate -i ../spec/api/schema.yaml -g typescript-axios -o src/lib/openapi 

