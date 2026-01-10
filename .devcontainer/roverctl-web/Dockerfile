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
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Install nodejs (20) and npm
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Install bun, pinned to a specific version
RUN curl -fsSL https://bun.com/install | bash -s "bun-v1.3.5"

# Install Java based on architecture
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
    JAVA_ARCH="openjdk-17-jdk"; \
    elif [ "$ARCH" = "arm64" ]; then \
    JAVA_ARCH="openjdk-17-jdk"; \
    else \
    echo "Unsupported architecture"; exit 1; \
    fi && \
    apt-get update && \
    apt-get install -y $JAVA_ARCH && \
    rm -rf /var/lib/apt/lists/*

# Verify Java installation
RUN java -version

# Verify Node.js and NPM installation
RUN node -v && npm -v

# Install openapi-generator-cli globally
RUN npm install -g @openapitools/openapi-generator-cli

# Verify openapi-generator-cli installation
RUN openapi-generator-cli version




