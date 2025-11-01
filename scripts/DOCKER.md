# Docker 跨平台构建指南

使用 Docker 进行 Yoga 库的跨平台编译，保持宿主系统干净。

## 前置要求

- Docker Engine 20.10+
- Docker Compose (可选)
- 支持 multi-platform builds (buildx)

## 快速开始

### 1. 使用 Docker 脚本构建

```bash
# 构建所有平台
./scripts/docker_build.sh --all-platforms

# 构建特定平台
./scripts/docker_build.sh --platform linux-amd64
./scripts/docker_build.sh --platform linux-arm64
./scripts/docker_build.sh --platform windows-amd64

# 列出支持的平台
./scripts/docker_build.sh --list-platforms

# 清理 Docker 镜像
./scripts/docker_build.sh --clean
```

### 2. 使用 Make 目标

```bash
# 构建 Docker 镜像
make docker-build-image

# 构建特定平台
make docker-build-linux-amd64
make docker-build-linux-arm64
make docker-build-windows-amd64

# 构建所有平台
make docker-build-all

# 清理
make docker-clean
```

### 3. 使用 Docker Compose

```bash
# 进入 scripts 目录
cd scripts

# 构建 Linux AMD64
docker-compose --profile linux up linux-amd64-builder

# 构建 Linux ARM64
docker-compose --profile linux up linux-arm64-builder

# 构建 Windows AMD64
docker-compose --profile windows up windows-amd64-builder

# 进入调试 shell
docker-compose --profile debug run shell
```

## 支持的平台

| 平台 | 架构 | 状态 | 备注 |
|------|------|------|------|
| Linux | amd64 | ✅ | 完全支持 |
| Linux | arm64 | ✅ | 需要交叉编译器 |
| Windows | amd64 | ✅ | 使用 MinGW |
| macOS | amd64 | ⚠️ | 需要 osxcross |
| macOS | arm64 | ⚠️ | 需要 osxcross |

## Docker 镜像结构

### 基础镜像 (Dockerfile)
- Ubuntu 22.04
- CMake, Git, Python3
- 交叉编译工具链

### 高级镜像 (Dockerfile.cross)
- 多阶段构建
- CMake toolchain 文件
- 预配置的构建脚本

## 环境变量

在 Docker 容器中可以使用以下环境变量：

- `CC`: C 编译器
- `CXX`: C++ 编译器
- `CMAKE_TOOLCHAIN_FILE`: CMake 工具链文件路径

## 故障排除

### 1. Docker multi-platform 支持问题

```bash
# 检查 buildx 是否可用
docker buildx version

# 创建 buildx builder
docker buildx create --name multiarch --use
docker buildx inspect --bootstrap
```

### 2. 交叉编译器问题

```bash
# 进入 Docker 容器检查
docker run -it yoga-builder:latest bash

# 检查交叉编译器
which aarch64-linux-gnu-gcc
which x86_64-w64-mingw32-gcc
```

### 3. 权限问题

```bash
# 确保 Docker 有权限访问项目文件
sudo chown -R $USER:$USER _libs/ include/
```

## 自定义配置

### 添加新平台

1. 在 `docker_build.sh` 中添加平台定义
2. 在 `Dockerfile.cross` 中添加相应的工具链
3. 更新 `build_yoga.py` 中的平台检测逻辑

### 使用自定义工具链

```bash
# 设置环境变量
export CC=your-custom-gcc
export CXX=your-custom-g++
export CMAKE_TOOLCHAIN_FILE=/path/to/toolchain.cmake

# 运行构建
./scripts/docker_build.sh --platform your-platform
```

## 性能优化

### 并行构建

Docker 构建脚本自动利用多核 CPU：

```bash
# 查看CPU核心数
nproc

# CMake 会自动使用可用的核心
```

### 缓存优化

```bash
# 使用 Docker build cache
docker build --cache-from yoga-builder:latest -t yoga-builder:latest .
```

## 安全注意事项

- Docker 镜像只包含必要的构建工具
- 构建过程在容器中隔离运行
- 不会影响宿主系统的工具链

## 维护

### 更新工具链

```bash
# 重新构建镜像
./scripts/docker_build.sh --clean
./scripts/docker_build.sh --build-image
```

### 清理资源

```bash
# 清理 Docker 资源
docker system prune -f
docker volume prune -f
```

## 相关文件

- `scripts/Dockerfile` - 基础 Docker 镜像
- `scripts/Dockerfile.cross` - 高级多阶段构建镜像
- `scripts/docker_build.sh` - Docker 构建脚本
- `scripts/docker-compose.yml` - Docker Compose 配置
- `scripts/DOCKER.md` - 本文档