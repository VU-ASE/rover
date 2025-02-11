.PHONY: build start clean test
SHELL := /bin/bash

# This Makefile entrypoint is used by CI

# Ensures that the latest version from the changelog file is in the rover/Cargo.toml
# Run this target before building anything
ifeq ($(VERSION),)
$(error Missing VERSION variable please pass it in as an argument)
endif


# The docker container has to be setup such that it is identical to the filesystem of
# the rovers. This script creates the necessary files to pretend it is a rover.
dev-setup:
	@echo ">>> "
	@echo ">>> STEP -1. Make sure that devcontainer has proper files in place"
	@echo ">>> "
	
	sudo ./scripts/setup_dev_files.sh

# This sets the version of roverd's Cargo.toml so that it is baked into the binary
# throug an additional build time Rust script.
version-update:
	@echo ">>> "
	@echo ">>> STEP 0. Set roverd version to $(VERSION)"
	@echo ">>> "

	./scripts/update_cargo_version.sh $(VERSION) ./roverd/Cargo.toml

lint:
	@cargo fmt
	@cargo clippy

test:
	@cargo test

gen-open-api:
	@echo ">>> "
	@echo ">>> STEP 1. Generting OpenAPI Rust code and running lints on it"
	@echo ">>> "

	@openapi-generator-cli generate -i ../spec/api/schema.yaml -g rust-axum -o openapi/ > /dev/null
	@cd openapi ; cargo fmt ; cargo clippy

	@echo ">>> "
	@echo ">>> STEP 2. Fixing OpenAPI Rust code"
	@echo ">>> "

	@./scripts/fix_open_api_gen.sh

build: lint dev-setup version-update gen-open-api
	@echo ">>> "
	@echo ">>> STEP 3. Build roverd"
	@echo ">>> "

	@cargo build --release

setup:
	sudo ./scripts/setup_rover_files.sh

dev: build
	sudo ./target/release/roverd

build-arm: lint version-update gen-open-api
	@echo ">>> "
	@echo ">>> STEP 3. Build roverd"
	@echo ">>> "
	cargo build --target=aarch64-unknown-linux-gnu --release

TEST_ROVER=rover12

dev-arm: build-arm
	scp ./target/aarch64-unknown-linux-gnu/release/roverd $(TEST_ROVER):/tmp/
	ssh $(TEST_ROVER) "echo debix | sudo -S mv /tmp/roverd /usr/local/bin/ ; echo debix | sudo -S chown root:root /usr/local/bin/roverd"

clean:
	@cargo clean

loc:
	@echo roverd:
	@cd roverd/src && find . -name '*.rs' | xargs wc -l
	@echo 

	@echo openapi:
	@cd openapi/src && find . -name '*.rs' | xargs wc -l
	@echo 

	@echo rovervalidate
	@cd rovervalidate/src && find . -name '*.rs' | xargs wc -l
	@echo 

