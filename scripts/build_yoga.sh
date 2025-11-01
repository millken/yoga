#!/bin/bash
# Cross-platform Yoga library builder script (simplified version)

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Project directories
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
YOGA_DIR="$PROJECT_ROOT/yoga"
BUILD_DIR="$YOGA_DIR/build"
INCLUDE_DIR="$PROJECT_ROOT/include"
LIBS_DIR="$PROJECT_ROOT/_libs"

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

check_dependencies() {
    log_info "Checking dependencies..."

    if ! command -v cmake &> /dev/null; then
        log_error "CMake is required but not installed"
        exit 1
    fi

    if ! command -v git &> /dev/null; then
        log_error "Git is required but not installed"
        exit 1
    fi

    # Platform-specific checks
    case "$(uname -s)" in
        Linux*)
            if ! command -v make &> /dev/null; then
                log_error "Make is required on Linux"
                exit 1
            fi
            ;;
        Darwin*)
            if ! command -v xcodebuild &> /dev/null; then
                log_warning "Xcode command line tools may not be installed"
            fi
            ;;
        CYGWIN*|MINGW*|MSYS*)
            log_warning "Make sure you're running from Visual Studio Developer Command Prompt"
            ;;
    esac

    log_success "Dependencies check passed"
}

update_submodule() {
    log_info "Updating yoga submodule..."
    cd "$PROJECT_ROOT"
    git submodule update --init --recursive
    log_success "Submodule updated"
}

clean_build() {
    log_info "Cleaning build directories..."
    rm -rf "$BUILD_DIR"
    log_success "Build directories cleaned"
}

build_yoga() {
    local platform="$1"
    local arch="$2"
    local build_subdir="$BUILD_DIR/${platform}-${arch}"

    log_info "Building Yoga library for ${platform}-${arch}..."

    mkdir -p "$build_subdir"
    cd "$build_subdir"

  # Configure based on platform
    case "$(uname -s)" in
        Linux*)
            cmake "$YOGA_DIR" \
                -G "Unix Makefiles" \
                -DCMAKE_BUILD_TYPE=Release \
                -DCMAKE_C_COMPILER=gcc \
                -DCMAKE_CXX_COMPILER=g++
            cmake --build . --target yogacore --config Release --parallel
            ;;
        Darwin*)
            if [[ "$arch" == "arm64" ]]; then
                cmake "$YOGA_DIR" \
                    -G Xcode \
                    -DCMAKE_BUILD_TYPE=Release \
                    -DCMAKE_OSX_ARCHITECTURES=arm64
                cmake --build . --target yogacore --config Release
            else
                cmake "$YOGA_DIR" \
                    -G Xcode \
                    -DCMAKE_BUILD_TYPE=Release \
                    -DCMAKE_OSX_ARCHITECTURES=x86_64
                cmake --build . --target yogacore --config Release
            fi
            ;;
        CYGWIN*|MINGW*|MSYS*)
            cmake "$YOGA_DIR" \
                -G "Visual Studio 16 2019" \
                -A x64 \
                -DCMAKE_BUILD_TYPE=Release
            cmake --build . --target yogacore --config Release
            ;;
    esac

    # Find and copy library
    local lib_name="libyogacore.a"
    case "$(uname -s)" in
        CYGWIN*|MINGW*|MSYS*)
            lib_name="yogacore.lib"
            ;;
    esac

    local lib_file=$(find . -name "$lib_name" -type f | head -1)
    if [[ -z "$lib_file" ]]; then
        log_error "Library file not found: $lib_name"
        exit 1
    fi

    local target_dir="$LIBS_DIR/$platform/$arch"
    mkdir -p "$target_dir"
    cp "$lib_file" "$target_dir/"
    log_success "Library installed to $target_dir/"

    # Copy headers (only .h files) from yoga source directory
    local yoga_source_dir="$YOGA_DIR/yoga"
    local target_include_dir="$INCLUDE_DIR/yoga"

    if [[ -d "$yoga_source_dir" ]]; then
        # Remove old headers
        rm -rf "$target_include_dir"
        mkdir -p "$INCLUDE_DIR"

        # Copy only header files
        find "$yoga_source_dir" -name "*.h" -type f | while read -r header_file; do
            # Create relative directory structure
            rel_path="${header_file#$yoga_source_dir/}"
            target_file="$target_include_dir/$rel_path"
            target_dir_path=$(dirname "$target_file")
            mkdir -p "$target_dir_path"
            cp "$header_file" "$target_file"
        done

        log_success "Headers installed to $target_include_dir/"

        # Count header files for verification
        header_count=$(find "$target_include_dir" -name "*.h" -type f | wc -l)
        log_info "Installed $header_count header files"
    else
        log_error "Yoga source directory not found: $yoga_source_dir"
        exit 1
    fi
}

install_headers() {
    log_info "This function is deprecated. Headers are now installed during build process."
    log_warning "Use build functions to install headers automatically."
}

build_current_platform() {
    local platform=""
    local arch=""

    case "$(uname -s)" in
        Linux*)  platform="linux";;
        Darwin*) platform="darwin";;
        CYGWIN*|MINGW*|MSYS*) platform="windows";;
        *)       log_error "Unsupported platform: $(uname -s)"; exit 1;;
    esac

    case "$(uname -m)" in
        x86_64|amd64) arch="amd64";;
        aarch64|arm64) arch="arm64";;
        *) log_error "Unsupported architecture: $(uname -m)"; exit 1;;
    esac

    build_yoga "$platform" "$arch"
}

build_all_platforms() {
    log_info "Building for all supported platforms..."

    # Build for Linux amd64
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        build_yoga "linux" "amd64"
    fi

    # Build for macOS (if on Mac)
    if [[ "$OSTYPE" == "darwin"* ]]; then
        build_yoga "darwin" "arm64"
        build_yoga "darwin" "x86_64"
    fi

    log_warning "Cross-compilation for other platforms requires proper toolchains"
    log_info "Consider using the Python script for full cross-platform builds"
}

show_help() {
    cat << EOF
Yoga Library Builder

Usage: $0 [OPTIONS]

OPTIONS:
    --help              Show this help message
    --check-deps        Check dependencies only
    --clean             Clean build directories only
    --current-platform  Build for current platform only (default)
    --all-platforms     Build for all supported platforms
    --skip-deps         Skip dependency checking

EXAMPLES:
    $0                           # Build for current platform
    $0 --check-deps              # Check dependencies only
    $0 --clean                   # Clean build directories
    $0 --all-platforms           # Build for all platforms (limited on current system)

For full cross-platform builds, use:
    python3 scripts/build_yoga.py --all-platforms
EOF
}

main() {
    local check_deps_only=false
    local clean_only=false
    local build_current=true
    local build_all=false
    local skip_deps=false

    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            --help)
                show_help
                exit 0
                ;;
            --check-deps)
                check_deps_only=true
                shift
                ;;
            --clean)
                clean_only=true
                shift
                ;;
            --current-platform)
                build_current=true
                build_all=false
                shift
                ;;
            --all-platforms)
                build_current=false
                build_all=true
                shift
                ;;
            --skip-deps)
                skip_deps=true
                shift
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done

    echo "Yoga Library Builder"
    echo "===================="

    # Check dependencies
    if [[ "$skip_deps" != true ]]; then
        check_dependencies
    fi

    if [[ "$check_deps_only" == true ]]; then
        log_success "Dependencies check completed"
        exit 0
    fi

    if [[ "$clean_only" == true ]]; then
        clean_build
        log_success "Clean completed"
        exit 0
    fi

    # Update submodule and build
    update_submodule
    clean_build

    if [[ "$build_all" == true ]]; then
        build_all_platforms
    else
        build_current_platform
    fi

    log_success "Yoga library build completed successfully!"
}

# Handle signals gracefully
trap 'log_error "Build interrupted"; exit 1' INT TERM

main "$@"