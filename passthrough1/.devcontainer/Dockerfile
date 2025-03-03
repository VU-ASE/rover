# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Set environment variables to prevent interactive prompts during package installation
ARG DEBIAN_FRONTEND=noninteractive

# Update package list and install necessary tools and dependencies
RUN apt-get update && \
    apt-get install -y \
    git \
    curl \
    gnupg \
    make \
    build-essential \
    pkg-config \
    libzmq3-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Go 1.21 based on the system architecture
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
    GO_ARCH="go1.21.0.linux-amd64.tar.gz"; \
    elif [ "$ARCH" = "arm64" ]; then \
    GO_ARCH="go1.21.0.linux-arm64.tar.gz"; \
    else \
    echo "Unsupported architecture"; exit 1; \
    fi && \
    curl -fsSL "https://go.dev/dl/$GO_ARCH" -o "$GO_ARCH" && \
    tar -C /usr/local -xzf "$GO_ARCH" && \
    rm "$GO_ARCH" && \
    # Install Go language server
    /usr/local/go/bin/go install golang.org/x/tools/gopls@latest

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.54.2

# Set up Go environment variables
ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

# Ensure CGO is enabled for the zmq4 package
ENV CGO_ENABLED=1

# Set the CGO CFLAGS and LDFLAGS for the zmq4 library
ENV CGO_CFLAGS_ALLOW="-std=gnu99"
ENV PKG_CONFIG_PATH="/usr/local/lib/pkgconfig"


