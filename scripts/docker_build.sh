#!/bin/bash
# Docker-based cross-compilation script for Yoga library
# Keeps the host system clean while building for multiple platforms

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Project directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
DOCKER_IMAGE="yoga-builder"
DOCKER_TAG="latest"

# Available platforms for cross-compilation
PLATFORMS=(
    "linux-amd64"
    "linux-arm64"
    "darwin-amd64"
    "darwin-arm64"
    "windows-amd64"
)

# Container runtime detection and setup
check_container_runtime() {
    local runtime=""
    local runtime_cmd=""

    # Try Docker first
    if command -v docker &> /dev/null && docker info &> /dev/null; then
        runtime="docker"
        runtime_cmd="docker"
        log_success "Using Docker runtime"
    # Try Podman
    elif command -v podman &> /dev/null; then
        runtime="podman"
        runtime_cmd="podman"
        log_success "Using Podman runtime"
    else
        log_error "Neither Docker nor Podman is available or running"
        log_error "Please install Docker or Podman"
        exit 1
    fi

    # Export for other functions to use
    export CONTAINER_RUNTIME=$runtime
    export CONTAINER_CMD=$runtime_cmd
}

# Build container image
build_container_image() {
    log_info "Building container image: ${DOCKER_IMAGE}:${DOCKER_TAG}"

    # Podman and Docker have slightly different syntax for some operations
    if [[ "$CONTAINER_RUNTIME" == "podman" ]]; then
        # Podman-specific flags
        $CONTAINER_CMD build \
            --format docker \
            -t "${DOCKER_IMAGE}:${DOCKER_TAG}" \
            -f "${SCRIPT_DIR}/Dockerfile" \
            "${PROJECT_ROOT}"
    else
        # Docker syntax
        $CONTAINER_CMD build \
            -t "${DOCKER_IMAGE}:${DOCKER_TAG}" \
            -f "${SCRIPT_DIR}/Dockerfile" \
            "${PROJECT_ROOT}"
    fi

    log_success "Container image built successfully"
}

# Build for specific platform using container runtime
build_platform() {
    local platform="$1"
    local IFS='-'
    read -ra PLATFORM_PARTS <<< "$platform"
    local os="${PLATFORM_PARTS[0]}"
    local arch="${PLATFORM_PARTS[1]}"

    log_info "Building Yoga library for ${os}-${arch} using ${CONTAINER_RUNTIME}..."

    # Determine platform-specific container settings
    local container_platform=""
    local cmake_toolchain=""
    local c_compiler=""
    local cxx_compiler=""

    case "${os}-${arch}" in
        "linux-amd64")
            container_platform="linux/amd64"
            c_compiler="gcc"
            cxx_compiler="g++"
            ;;
        "linux-arm64")
            container_platform="linux/arm64"
            c_compiler="aarch64-linux-gnu-gcc"
            cxx_compiler="aarch64-linux-gnu-g++"
            ;;
        "darwin-amd64")
            # macOS cross-compilation (using osxcross or similar would be needed)
            log_warning "macOS cross-compilation requires additional setup (osxcross)"
            log_info "Skipping ${platform} - use native macOS build instead"
            return 0
            ;;
        "darwin-arm64")
            # macOS cross-compilation (using osxcross or similar would be needed)
            log_warning "macOS cross-compilation requires additional setup (osxcross)"
            log_info "Skipping ${platform} - use native macOS build instead"
            return 0
            ;;
        "windows-amd64")
            container_platform="linux/amd64"
            c_compiler="x86_64-w64-mingw32-gcc"
            cxx_compiler="x86_64-w64-mingw32-g++"
            ;;
        *)
            log_error "Unsupported platform: ${platform}"
            return 1
            ;;
    esac

    # Create build command
    local build_cmd="cd /workspace && "

    # Set up cross-compilation environment
    if [[ -n "$c_compiler" && -n "$cxx_compiler" ]]; then
        build_cmd+="export CC=${c_compiler} CXX=${cxx_compiler} && "
    fi

    # Map architecture names correctly for the build script
    local build_arch="$arch"
    if [[ "$os" == "windows" && "$arch" == "amd64" ]]; then
        build_arch="x64"  # Windows uses x64, not amd64
    fi

    # Run the build script
    build_cmd+="./scripts/build_yoga.py --platform ${os} ${build_arch}"

    # Run in container with appropriate platform
    local container_args=(
        "--rm"
        "-v" "${PROJECT_ROOT}:/workspace"
        "-w" "/workspace"
    )

    # Add platform-specific args for different runtimes
    if [[ -n "$container_platform" ]]; then
        if [[ "$CONTAINER_RUNTIME" == "podman" ]]; then
            # Podman platform syntax
            container_args+=("--arch" "$(echo $container_platform | cut -d'/' -f2)")
        else
            # Docker platform syntax
            container_args+=("--platform" "$container_platform")
        fi
    fi

    # Add runtime-specific flags
    if [[ "$CONTAINER_RUNTIME" == "podman" ]]; then
        # Podman might need additional flags for some systems
        container_args+=("--userns" "keep-id")
        # Add pull policy to prefer local images
        container_args+=("--pull" "never")
    else
        # Docker pull policy
        container_args+=("--pull" "never")
    fi

    # Use the image (prefer local, no remote pull)
    local image_name="${DOCKER_IMAGE}"
    if [[ -n "$DOCKER_TAG" ]]; then
        image_name="${DOCKER_IMAGE}:${DOCKER_TAG}"
    fi

    $CONTAINER_CMD run "${container_args[@]}" "$image_name" /bin/bash -c "$build_cmd"

    log_success "Build completed for ${platform}"
}

# Build for all platforms
build_all_platforms() {
    log_info "Building for all platforms using Docker..."

    for platform in "${PLATFORMS[@]}"; do
        build_platform "$platform" || log_warning "Failed to build for ${platform}"
    done
}

# Show help
show_help() {
    cat << EOF
Container-based Yoga Library Cross-Compilation Builder (Docker/Podman)

Usage: $0 [OPTIONS]

OPTIONS:
    --help              Show this help message
    --build-image       Build container image only
    --all-platforms     Build for all supported platforms
    --platform PLATFORM Build for specific platform (e.g., linux-amd64, linux-arm64, windows-amd64)
    --list-platforms    List available platforms
    --clean             Remove container image

SUPPORTED PLATFORMS:
EOF
    for platform in "${PLATFORMS[@]}"; do
        echo "    $platform"
    done
    echo ""
    cat << EOF
EXAMPLES:
    $0 --build-image                    # Build Docker image
    $0 --platform linux-amd64           # Build for Linux AMD64
    $0 --platform linux-arm64           # Build for Linux ARM64
    $0 --platform windows-amd64         # Build for Windows AMD64
    $0 --all-platforms                  # Build for all platforms
    $0 --clean                          # Remove Docker image

REQUIREMENTS:
    - Docker or Podman installed and running
    - Container runtime supports multi-platform builds

NOTE:
    - macOS cross-compilation requires osxcross setup
    - Some platforms may require additional toolchain configuration
    - Podman users may need additional configuration for multi-arch builds
EOF
}

# List available platforms
list_platforms() {
    log_info "Available platforms:"
    for platform in "${PLATFORMS[@]}"; do
        echo "  - $platform"
    done
}

# Clean container image
clean_container_image() {
    log_info "Removing container image: ${DOCKER_IMAGE}"

    # Try to remove tagged version first
    if $CONTAINER_CMD image inspect "${DOCKER_IMAGE}:${DOCKER_TAG}" &> /dev/null; then
        $CONTAINER_CMD rmi "${DOCKER_IMAGE}:${DOCKER_TAG}" 2>/dev/null || true
        log_success "Container image removed: ${DOCKER_IMAGE}:${DOCKER_TAG}"
    fi

    # Also try to remove untagged version
    if $CONTAINER_CMD image inspect "${DOCKER_IMAGE}" &> /dev/null; then
        $CONTAINER_CMD rmi "${DOCKER_IMAGE}" 2>/dev/null || true
        log_success "Container image removed: ${DOCKER_IMAGE}"
    fi
}

# Main function
main() {
    log_info "Container-based Yoga Library Cross-Compilation Builder"
    echo "============================================================="

    check_container_runtime

    local build_image_only=false
    local target_platform=""
    local build_all=false

    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            --help)
                show_help
                exit 0
                ;;
            --build-image)
                build_image_only=true
                shift
                ;;
            --all-platforms)
                build_all=true
                shift
                ;;
            --platform)
                target_platform="$2"
                shift 2
                ;;
            --list-platforms)
                list_platforms
                exit 0
                ;;
            --clean)
                clean_container_image
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done

    # Build container image if needed
    if [[ "$build_image_only" == true ]]; then
        build_container_image
        exit 0
    fi

    # Check if container image exists, build if not
    # Check both tagged and untagged versions
    local image_found=false

    if $CONTAINER_CMD image inspect "${DOCKER_IMAGE}:${DOCKER_TAG}" &> /dev/null; then
        log_success "Container image found: ${DOCKER_IMAGE}:${DOCKER_TAG}"
        image_found=true
    elif $CONTAINER_CMD image inspect "${DOCKER_IMAGE}" &> /dev/null; then
        log_success "Container image found: ${DOCKER_IMAGE} (using latest tag)"
        image_found=true
        DOCKER_TAG=""  # Use untagged version
    else
        log_info "Container image not found, building..."
        build_container_image
        image_found=true
    fi

    if [[ "$image_found" == false ]]; then
        log_error "Failed to build or find container image"
        exit 1
    fi

    # Execute build
    if [[ "$build_all" == true ]]; then
        build_all_platforms
    elif [[ -n "$target_platform" ]]; then
        if [[ ! " ${PLATFORMS[@]} " =~ " ${target_platform} " ]]; then
            log_error "Unsupported platform: $target_platform"
            list_platforms
            exit 1
        fi
        build_platform "$target_platform"
    else
        log_warning "No build target specified."
        show_help
        exit 1
    fi

    log_success "Cross-compilation build completed!"
}

# Handle signals gracefully
trap 'log_error "Build interrupted"; exit 1' INT TERM

main "$@"