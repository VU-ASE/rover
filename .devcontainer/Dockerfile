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

# Create a non-root user with a home directory and bash as default shell
ARG USERNAME=devuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m -s /bin/bash $USERNAME \
    && apt-get update && apt-get install -y sudo \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Copy skeleton files to set up the user's shell environment
RUN cp /etc/skel/.bashrc /home/$USERNAME/.bashrc \
    && cp /etc/skel/.profile /home/$USERNAME/.profile \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME

# Switch to the new user
USER $USERNAME

# Set the working directory
WORKDIR /home/$USERNAME
# Set up Go environment variables
ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

# Ensure CGO is enabled for the zmq4 package
ENV CGO_ENABLED=1

# Set the CGO CFLAGS and LDFLAGS for the zmq4 library
ENV CGO_CFLAGS_ALLOW="-std=gnu99"
ENV PKG_CONFIG_PATH="/usr/local/lib/pkgconfig"

