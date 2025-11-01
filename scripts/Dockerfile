# Multi-stage Dockerfile for cross-compiling Yoga library
# Supports building for multiple architectures while keeping build environment clean

# Stage 1: Build environment
FROM ubuntu:22.04 AS builder

# Install build dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    python3 \
    python3-pip \
    ninja-build \
    gcc-aarch64-linux-gnu \
    g++-aarch64-linux-gnu \
    gcc-x86-64-linux-gnu \
    g++-x86-64-linux-gnu \
    gcc-mingw-w64-x86-64 \
    g++-mingw-w64-x86-64 \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY scripts/requirements.txt /tmp/
RUN pip3 install -r /tmp/requirements.txt || true

# Set working directory
WORKDIR /workspace

# Copy build scripts
COPY scripts/ /workspace/scripts/
RUN chmod +x /workspace/scripts/build_yoga.sh

# Default command
CMD ["/bin/bash"]