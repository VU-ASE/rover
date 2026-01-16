.PHONY: build start clean test

build-open-api:
	# Check if the API spec file exists
	@if [ ! -f ../spec/api/schema.yaml ]; then \
		echo "../spec/api/schema.yaml file not found."; \
		exit 1; \
	fi
	@echo "generating openapi client"
	@openapi-generator-cli generate -i ../spec/api/schema.yaml -g typescript-axios -o src/lib/openapi 

install:
	@echo "Installing dependencies..."
	@bun install

start: install
	@echo "Starting the application..."
	@echo ""
	@echo "\033[0;33m⚠️  WARNING: Starting in DEV mode\033[0m"
	@echo "\033[0;33m⚠️  You need a running 'roverctl passthrough' for the passthrough server to work\033[0m"
	@echo ""
	@echo "Please enter configuration values (press Enter for defaults):"
	@echo ""
	@read -p "PUBLIC_ROVERD_HOST [rover01.local]: " roverd_host; \
	roverd_host=$${roverd_host:-rover01.local}; \
	read -p "PUBLIC_ROVERD_PORT [80]: " roverd_port; \
	roverd_port=$${roverd_port:-80}; \
	read -p "PUBLIC_ROVERD_USERNAME [debix]: " roverd_username; \
	roverd_username=$${roverd_username:-debix}; \
	read -p "PUBLIC_ROVERD_PASSWORD [debix]: " roverd_password; \
	roverd_password=$${roverd_password:-debix}; \
	read -p "PUBLIC_PASSTHROUGH_HOST [empty]: " passthrough_host; \
	read -p "PUBLIC_PASSTHROUGH_PORT [7500]: " passthrough_port; \
	passthrough_port=$${passthrough_port:-7500}; \
	echo ""; \
	echo "Starting with configuration:"; \
	echo "  ROVERD_HOST: $$roverd_host"; \
	echo "  ROVERD_PORT: $$roverd_port"; \
	echo "  ROVERD_USERNAME: $$roverd_username"; \
	echo "  PASSTHROUGH_HOST: $$passthrough_host"; \
	echo "  PASSTHROUGH_PORT: $$passthrough_port"; \
	echo ""; \
	PUBLIC_ROVERD_HOST=$$roverd_host \
	PUBLIC_ROVERD_PORT=$$roverd_port \
	PUBLIC_ROVERD_USERNAME=$$roverd_username \
	PUBLIC_ROVERD_PASSWORD=$$roverd_password \
	PUBLIC_PASSTHROUGH_HOST=$$passthrough_host \
	PUBLIC_PASSTHROUGH_PORT=$$passthrough_port \
	bun run dev --host 0.0.0.0