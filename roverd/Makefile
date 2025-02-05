.PHONY: build start clean test
SHELL := /bin/bash

# This Makefile entrypoint is used by CI

# Ensures that the latest version from the changelog file is in the rover/Cargo.toml
# Run this target before building anything
version-update:
	./scripts/update_cargo_version.sh `cat CHANGELOG.md | ./scripts/latest_changelog.sh` ./roverd/Cargo.toml 

lint:
	@cargo fmt
	@cargo clippy

test:
	@cargo test

build: version-update
	@cargo build --release

setup:
	sudo ./scripts/setup_rover_files.sh

dev: build
	sudo ./target/release/roverd

build-arm: version-update
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
	@cd roverd/openapi/src && find . -name '*.rs' | xargs wc -l
	@echo 

	@echo rovervalidate
	@cd rovervalidate/src && find . -name '*.rs' | xargs wc -l
	@echo 


