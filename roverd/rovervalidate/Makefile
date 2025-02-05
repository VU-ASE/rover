.PHONY: lint test build clean
SHELL := /bin/bash

lint:
	@cargo fmt
	@cargo clippy

test: lint
	@cargo test

build: lint
	@cargo build --release

clean:
	@cargo clean
