.PHONY: lint build-open-api build run test clean
SHELL := /bin/bash


lint:
	@cargo fmt
	@cargo clippy

build-open-api:
	@openapi-generator-cli generate -i ../../spec/api/schema.yaml -g rust-axum -o openapi/
	@cd openapi ; cargo fmt ; cargo clippy

build: build-open-api lint
	@cargo build --release

run: build
	@sudo ./target/release/roverd

test:
	@cargo test -- --show-output --test-threads 1

clean:
	rm -rf openapi/
	cargo clean
