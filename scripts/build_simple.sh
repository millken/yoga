#!/bin/bash
# 简化版的Yoga库构建脚本
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

# 简单的构建函数
build_yogacore() {
    log_info "Building Yoga core library..."

    # 创建构建目录
    BUILD_DIR="$YOGA_DIR/build"
    mkdir -p "$BUILD_DIR"
    cd "$BUILD_DIR"

    # 检查操作系统
    case "$(uname -s)" in
        Linux*)
            log_info "Detected Linux system"

            # 如果cmake不可用，尝试简单的make构建
            if command -v cmake >/dev/null 2>&1; then
                log_info "Using CMake build system"
                cmake .. -DCMAKE_BUILD_TYPE=Release -DYOGA_BUILD_TESTS=OFF
                make yogacore -j$(nproc 2>/dev/null || echo 4)
            else
                log_error "CMake is required for building Yoga library"
                exit 1
            fi
            ;;

        Darwin*)
            log_info "Detected macOS system"

            if command -v cmake >/dev/null 2>&1; then
                log_info "Using CMake build system"
                cmake .. -DCMAKE_BUILD_TYPE=Release -DYOGA_BUILD_TESTS=OFF
                make yogacore -j$(sysctl -n hw.ncpu 2>/dev/null || echo 4)
            else
                log_error "CMake is required for building Yoga library"
                exit 1
            fi
            ;;

        CYGWIN*|MINGW*|MSYS*)
            log_info "Detected Windows system"

            if command -v cmake >/dev/null 2>&1; then
                cmake .. -G "Visual Studio 16 2019" -A x64 -DCMAKE_BUILD_TYPE=Release -DYOGA_BUILD_TESTS=OFF
                cmake --build . --target yogacore --config Release
            else
                log_error "CMake is required for building Yoga library"
                exit 1
            fi
            ;;

        *)
            log_error "Unsupported operating system: $(uname -s)"
            exit 1
            ;;
    esac
}

# 安装头文件
install_headers() {
    log_info "Installing header files..."

    if [ ! -d "$YOGA_DIR/yoga" ]; then
        log_error "Yoga source directory not found"
        exit 1
    fi

    # 删除旧的头文件
    rm -rf "$INCLUDE_DIR/yoga"

    # 复制新的头文件
    cp -r "$YOGA_DIR/yoga" "$INCLUDE_DIR/"
    log_success "Headers installed to $INCLUDE_DIR/yoga/"
}

# 安装库文件
install_library() {
    log_info "Installing library file..."

    # 查找构建的库文件
    LIB_FILE=$(find "$YOGA_DIR/build" -name "libyogacore.a" -type f 2>/dev/null | head -1)

    if [ -z "$LIB_FILE" ]; then
        log_error "Built library file not found"
        exit 1
    fi

    # 确定平台和架构
    PLATFORM=""
    ARCH=""

    case "$(uname -s)" in
        Linux*)  PLATFORM="linux";;
        Darwin*) PLATFORM="darwin";;
        CYGWIN*|MINGW*|MSYS*) PLATFORM="windows";;
    esac

    case "$(uname -m)" in
        x86_64|amd64) ARCH="x86_64";;
        aarch64|arm64) ARCH="arm64";;
        *) ARCH="unknown";;
    esac

    if [ "$PLATFORM" = "unknown" ] || [ "$ARCH" = "unknown" ]; then
        log_warning "Could not determine platform/architecture, using default"
        PLATFORM="linux"
        ARCH="x86_64"
    fi

    # 创建目标目录
    TARGET_DIR="$LIBS_DIR/$PLATFORM/$ARCH"
    mkdir -p "$TARGET_DIR"

    # 复制库文件
    cp "$LIB_FILE" "$TARGET_DIR/"
    log_success "Library installed to $TARGET_DIR/"
}

# 检查依赖
check_dependencies() {
    log_info "Checking dependencies..."

    if ! command -v git >/dev/null 2>&1; then
        log_error "Git is required"
        exit 1
    fi

    if ! command -v cmake >/dev/null 2>&1; then
        log_error "CMake is required"
        exit 1
    fi

    case "$(uname -s)" in
        Linux*)
            if ! command -v make >/dev/null 2>&1; then
                log_error "Make is required on Linux"
                exit 1
            fi
            ;;
    esac

    log_success "Dependencies check passed"
}

# 更新子模块
update_submodule() {
    log_info "Updating yoga submodule..."
    cd "$PROJECT_ROOT"
    git submodule update --init --recursive
    log_success "Submodule updated"
}

# 显示帮助
show_help() {
    cat << EOF
Simple Yoga Library Builder

Usage: $0 [OPTIONS]

OPTIONS:
    --help              Show this help message
    --check-deps        Check dependencies only
    --headers-only      Install headers only
    --skip-deps         Skip dependency checking
    --skip-submodule    Skip submodule update

EXAMPLES:
    $0                  # Build and install everything
    $0 --check-deps     # Check dependencies only
    $0 --headers-only   # Install headers only

EOF
}

# 主函数
main() {
    local check_deps_only=false
    local headers_only=false
    local skip_deps=false
    local skip_submodule=false

    # 解析参数
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
            --headers-only)
                headers_only=true
                shift
                ;;
            --skip-deps)
                skip_deps=true
                shift
                ;;
            --skip-submodule)
                skip_submodule=true
                shift
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done

    echo "Simple Yoga Library Builder"
    echo "==========================="

    # 检查依赖
    if [ "$skip_deps" != true ]; then
        check_dependencies
    fi

    if [ "$check_deps_only" = true ]; then
        log_success "Dependencies check completed"
        exit 0
    fi

    # 更新子模块
    if [ "$skip_submodule" != true ]; then
        update_submodule
    fi

    # 安装头文件
    install_headers

    # 如果只需要头文件，则退出
    if [ "$headers_only" = true ]; then
        log_success "Headers installation completed"
        exit 0
    fi

    # 构建库
    build_yogacore
    install_library

    log_success "Yoga library build completed successfully!"
}

# 错误处理
trap 'log_error "Build interrupted"; exit 1' INT TERM

# 运行主函数
main "$@"