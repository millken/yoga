#!/usr/bin/env python3
"""
Cross-platform Yoga library builder and updater.
This script compiles the Yoga library and replaces static libraries and header files.
Supports Windows, macOS, and Linux with multiple architectures.
"""

import os
import sys
import platform
import subprocess
import shutil
import argparse
from pathlib import Path
from typing import List, Tuple, Optional

# Project structure
PROJECT_ROOT = Path(__file__).parent.parent
YOGA_SUBMODULE_DIR = PROJECT_ROOT / "yoga"
INCLUDE_DIR = PROJECT_ROOT / "include"
LIBS_DIR = PROJECT_ROOT / "_libs"
SCRIPTS_DIR = PROJECT_ROOT / "scripts"

# Platform-specific configurations
PLATFORMS = {
    "darwin": {
        "architectures": ["arm64", "x86_64"],
        "cmake_generator": "Xcode",
        "lib_extension": ".a",
        "lib_name": "libyogacore.a"
    },
    "linux": {
        "architectures": ["amd64"],
        "cmake_generator": "Unix Makefiles",
        "lib_extension": ".a",
        "lib_name": "libyogacore.a"
    },
    "windows": {
        "architectures": ["x64"],
        "cmake_generator": "Ninja",
        "lib_extension": ".lib",
        "lib_name": "yogacore.lib"
    }
}

class Colors:
    """ANSI color codes for terminal output"""
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

def log_info(message: str):
    """Print info message"""
    print(f"{Colors.OKCYAN}[INFO]{Colors.ENDC} {message}")

def log_success(message: str):
    """Print success message"""
    print(f"{Colors.OKGREEN}[SUCCESS]{Colors.ENDC} {message}")

def log_warning(message: str):
    """Print warning message"""
    print(f"{Colors.WARNING}[WARNING]{Colors.ENDC} {message}")

def log_error(message: str):
    """Print error message"""
    print(f"{Colors.FAIL}[ERROR]{Colors.ENDC} {message}")

def run_command(cmd: List[str], cwd: Optional[Path] = None, check: bool = True, capture_output: bool = True) -> subprocess.CompletedProcess:
    """Run a command with proper error handling"""
    cmd_str = ' '.join(cmd)
    log_info(f"Running: {cmd_str}")

    try:
        result = subprocess.run(
            cmd,
            cwd=cwd,
            check=check,
            capture_output=capture_output,
            text=True
        )
        return result
    except subprocess.CalledProcessError as e:
        log_error(f"Command failed: {cmd_str}")
        log_error(f"Return code: {e.returncode}")
        if e.stdout:
            log_error(f"STDOUT:\n{e.stdout}")
        if e.stderr:
            log_error(f"STDERR:\n{e.stderr}")
        raise

def check_dependencies():
    """Check if required dependencies are available"""
    log_info("Checking dependencies...")

    # Check cmake
    try:
        result = run_command(["cmake", "--version"])
        version = result.stdout.split()[2]
        log_info(f"Found CMake version: {version}")
    except (subprocess.CalledProcessError, FileNotFoundError):
        log_error("CMake is required but not found. Please install CMake.")
        sys.exit(1)

    # Check git
    try:
        run_command(["git", "--version"])
    except (subprocess.CalledProcessError, FileNotFoundError):
        log_error("Git is required but not found. Please install Git.")
        sys.exit(1)

    # Platform-specific checks
    current_platform = platform.system().lower()
    if current_platform == "linux":
        # Check for build-essential packages
        try:
            run_command(["gcc", "--version"])
            run_command(["make", "--version"])
        except (subprocess.CalledProcessError, FileNotFoundError):
            log_error("GCC and Make are required on Linux. Please install build-essential.")
            sys.exit(1)
    elif current_platform == "windows":
        # Check for Visual Studio
        try:
            run_command(["cl", "/?"], check=False)
        except FileNotFoundError:
            log_warning("Visual Studio C++ compiler not found in PATH. Make sure you're using Developer Command Prompt.")

def update_submodule():
    """Update the yoga submodule to latest version"""
    log_info("Updating yoga submodule...")

    # Initialize and update submodule
    run_command(["git", "submodule", "update", "--init", "--recursive"], cwd=PROJECT_ROOT)
    log_success("Yoga submodule updated")

def clean_build_dirs():
    """Clean build directories"""
    log_info("Cleaning build directories...")

    yoga_build_dir = YOGA_SUBMODULE_DIR / "build"
    if yoga_build_dir.exists():
        shutil.rmtree(yoga_build_dir)
        log_info(f"Removed {yoga_build_dir}")

def manual_build_windows_yoga(build_dir: Path, c_compiler: str, cxx_compiler: str):
    """Manually build Yoga library using MinGW when CMake Makefile fails"""
    log_info("Attempting manual MinGW build of Yoga library...")

    yoga_source_dir = YOGA_SUBMODULE_DIR / "yoga"

    # Find all source files
    source_files = []
    source_files.extend(yoga_source_dir.glob("**/*.cpp"))
    source_files.extend(yoga_source_dir.glob("**/*.c"))

    if not source_files:
        raise RuntimeError("No source files found in yoga directory")

    log_info(f"Found {len(source_files)} source files")

    # Create object files directory
    obj_dir = build_dir / "manual_objects"
    obj_dir.mkdir(exist_ok=True)

    # Compile each source file
    object_files = []
    for source_file in source_files:
        obj_file = obj_dir / (source_file.stem + ".o")
        object_files.append(obj_file)

        compile_cmd = [c_compiler, "-c", str(source_file), "-o", str(obj_file),
                      "-I", str(yoga_source_dir),
                      "-O2", "-DNDEBUG", "-DYOGA_EXPORT=",
                      "-static", "-static-libgcc", "-static-libstdc++"]

        log_info(f"Compiling {source_file.name}...")
        run_command(compile_cmd, capture_output=False)

    # Create static library
    lib_name = "yogacore.lib"
    lib_path = build_dir / lib_name

    # Use ar (mingw32-ar) to create the library
    ar_cmd = ["x86_64-w64-mingw32-ar", "rcs", str(lib_path)] + [str(f) for f in object_files]
    log_info(f"Creating static library {lib_name}...")
    run_command(ar_cmd, capture_output=False)

    log_success(f"Manual build completed: {lib_path}")
    return lib_path

def build_yoga_library(platform_name: str, arch: str) -> Path:
    """Build Yoga library for specific platform and architecture"""
    log_info(f"Building Yoga library for {platform_name}-{arch}...")

    platform_config = PLATFORMS[platform_name]
    build_dir = YOGA_SUBMODULE_DIR / "build" / f"{platform_name}-{arch}"
    build_dir.mkdir(parents=True, exist_ok=True)

    # Configure CMake - build only the core library
    # Select appropriate generator
    generator = platform_config["cmake_generator"]

    # Check if the preferred generator is available
    if platform_name == "windows":
        # For Windows cross-compilation, let CMake auto-select the best generator
        # This avoids forcing Unix Makefiles which might have issues
        log_info("Letting CMake auto-select generator for Windows cross-compilation")
        generator = None  # Remove generator specification
    elif platform_name == "linux" and generator == "Unix Makefiles":
        if shutil.which("ninja"):
            generator = "Ninja"
            log_info("Using Ninja generator for faster builds")
    elif platform_name == "darwin" and generator == "Xcode":
        # Xcode is usually available on macOS
        pass

    cmake_config_cmd = [
        "cmake",
        str(YOGA_SUBMODULE_DIR),
        "-DCMAKE_BUILD_TYPE=Release",
        "-DYOGA_BUILD_TESTS=OFF",
        "-DYOGA_BUILD_EXPORT=ON",
        "-DBUILD_SHARED_LIBS=OFF",  # Build static library only
        "-DYOGA_BUILD_SAMPLES=OFF"   # Don't build samples
    ]

    # Add generator only if specified
    if generator:
        cmake_config_cmd.extend(["-G", generator])

    # Add architecture-specific flags and toolchain files
    if platform_name == "darwin":
        if arch == "arm64":
            cmake_config_cmd.extend(["-DCMAKE_OSX_ARCHITECTURES=arm64"])
        else:
            cmake_config_cmd.extend(["-DCMAKE_OSX_ARCHITECTURES=x86_64"])
    elif platform_name == "linux":
        if arch == "arm64":
            # Check for cross-compilation environment
            if os.environ.get("CMAKE_TOOLCHAIN_FILE"):
                cmake_config_cmd.extend(["-DCMAKE_TOOLCHAIN_FILE=" + os.environ["CMAKE_TOOLCHAIN_FILE"]])
            else:
                # Try to detect cross-compilation setup
                cross_compilers = ["aarch64-linux-gnu-gcc", "aarch64-none-linux-gnu-gcc"]
                c_compiler = None
                for compiler in cross_compilers:
                    if shutil.which(compiler):
                        c_compiler = compiler
                        break

                if c_compiler:
                    cxx_compiler = c_compiler.replace("gcc", "g++")
                    cmake_config_cmd.extend([
                        "-DCMAKE_C_COMPILER=" + c_compiler,
                        "-DCMAKE_CXX_COMPILER=" + cxx_compiler
                    ])
                    log_info(f"Using cross-compiler: {c_compiler}")
                else:
                    log_warning("ARM64 cross-compiler not found, falling back to native compiler")
                    cmake_config_cmd.extend(["-DCMAKE_C_COMPILER=gcc", "-DCMAKE_CXX_COMPILER=g++"])
        else:
            cmake_config_cmd.extend(["-DCMAKE_C_COMPILER=gcc", "-DCMAKE_CXX_COMPILER=g++"])
    elif platform_name == "windows":
        # Check for MinGW cross-compiler
        if os.environ.get("CMAKE_TOOLCHAIN_FILE"):
            cmake_config_cmd.extend(["-DCMAKE_TOOLCHAIN_FILE=" + os.environ["CMAKE_TOOLCHAIN_FILE"]])
        else:
            mingw_compilers = ["x86_64-w64-mingw32-gcc", "x86_64-pc-mingw32-gcc"]
            c_compiler = None
            for compiler in mingw_compilers:
                if shutil.which(compiler):
                    c_compiler = compiler
                    break

            if c_compiler:
                cxx_compiler = c_compiler.replace("gcc", "g++")
                cmake_config_cmd.extend([
                    "-DCMAKE_C_COMPILER=" + c_compiler,
                    "-DCMAKE_CXX_COMPILER=" + cxx_compiler
                ])
                log_info(f"Using MinGW cross-compiler: {c_compiler}")
            else:
                log_warning("MinGW cross-compiler not found, falling back to native compiler")
                cmake_config_cmd.extend(["-DCMAKE_C_COMPILER=gcc", "-DCMAKE_CXX_COMPILER=g++"])

    run_command(cmake_config_cmd, cwd=build_dir)

    # Check what files were created after CMake configuration
    log_info("Checking build directory contents after CMake configuration:")
    build_files = list(build_dir.glob("*"))
    for file in sorted(build_files):
        log_info(f"  - {file.name}")

    # Build only the yogacore target (no tests, no examples)
    if platform_name == "windows":
        # For Windows cross-compilation, try different approaches based on available files
        makefiles = list(build_dir.glob("Makefile*"))
        ninja_files = list(build_dir.glob("build.ninja"))

        if makefiles:
            log_info("Found Makefile, using make command...")
            # First try to see what targets are available
            try:
                log_info("Available make targets:")
                run_command(["make", "help"], cwd=build_dir, capture_output=False)
            except:
                pass

            # Try verbose make to see what's failing
            build_cmd = ["make", "yogacore", "VERBOSE=1"]
            log_info("Building Windows library with make (verbose)...")
            try:
                run_command(build_cmd, cwd=build_dir, capture_output=False)
            except subprocess.CalledProcessError as e:
                log_error("Make failed. Let's try basic compilation test...")
                # Try to compile a simple file to identify the issue
                try:
                    simple_build_cmd = ["make", "VERBOSE=1"]
                    log_info("Trying basic make to see detailed error...")
                    run_command(simple_build_cmd, cwd=build_dir, capture_output=False)
                except:
                    log_error("Even basic make failed. Trying to diagnose MinGW setup...")
                    # Check if MinGW is working
                    try:
                        run_command(["x86_64-w64-mingw32-gcc", "--version"], capture_output=False)
                        run_command(["x86_64-w64-mingw32-g++", "--version"], capture_output=False)
                        log_info("MinGW compilers are available")

                        # Try to test direct compilation without CMake
                        log_info("Testing direct MinGW compilation...")
                        test_file = build_dir / "test.cpp"
                        test_file.write_text("int main() { return 0; }")

                        test_compile = ["x86_64-w64-mingw32-g++", "-c", str(test_file), "-o", str(build_dir / "test.o")]
                        run_command(test_compile, capture_output=False)
                        log_info("Direct MinGW compilation works")

                        # Check CMake files for issues
                        makefile = build_dir / "Makefile"
                        if makefile.exists():
                            log_info("Checking first few lines of Makefile...")
                            makefile_content = makefile.read_text().split('\n')[:20]
                            for i, line in enumerate(makefile_content):
                                log_info(f"  {i+1}: {line}")

                        # Try to run make with different targets
                        log_info("Trying different make targets...")
                        try:
                            log_info("Trying 'make all'...")
                            run_command(["make", "all"], cwd=build_dir, capture_output=False)
                        except:
                            try:
                                log_info("Trying 'make' without target...")
                                run_command(["make"], cwd=build_dir, capture_output=False)
                            except:
                                log_error("All make targets failed")
                                # Try building manually with MinGW since CMake Makefile is broken
                                log_info("CMake Makefile is broken. Trying manual MinGW build...")
                                manual_build_windows_yoga(build_dir, c_compiler, cxx_compiler)

                    except Exception as e:
                        log_error(f"MinGW test failed: {e}")
                raise
        elif ninja_files:
            log_info("Found build.ninja, using ninja command...")
            build_cmd = ["ninja", "yogacore"]
            log_info("Building Windows library with ninja...")
            run_command(build_cmd, cwd=build_dir, capture_output=False)
        else:
            # Fallback to cmake build
            log_info("No specific build files found, using cmake build...")
            build_cmd = ["cmake", "--build", ".", "--target", "yogacore", "--config", "Release"]
            run_command(build_cmd, cwd=build_dir, capture_output=False)
    else:
        # Use cmake build for other platforms
        build_cmd = ["cmake", "--build", ".", "--target", "yogacore", "--config", "Release", "--parallel"]
        if platform_name == "linux" and generator == "Ninja":
            build_cmd.extend(["-j", str(os.cpu_count())])

        run_command(build_cmd, cwd=build_dir)

    log_success(f"Yoga library built for {platform_name}-{arch}")
    return build_dir

def install_headers():
    """Install Yoga header files"""
    log_info("Installing header files...")

    yoga_header_dir = YOGA_SUBMODULE_DIR / "yoga"
    target_header_dir = INCLUDE_DIR / "yoga"

    # Remove old headers
    if target_header_dir.exists():
        shutil.rmtree(target_header_dir)

    # Copy new headers
    shutil.copytree(yoga_header_dir, target_header_dir)
    log_success(f"Headers installed to {target_header_dir}")

def install_library(platform_name: str, arch: str, build_dir: Path):
    """Install built library and headers by copying directly from build directory"""
    log_info(f"Installing library and headers for {platform_name}-{arch}...")

    platform_config = PLATFORMS[platform_name]

    # Find the library file in build directory (not install directory)
    lib_files = list(build_dir.rglob(platform_config["lib_name"]))
    if not lib_files:
        # Try different patterns for different platforms
        if platform_name == "windows":
            # Windows libraries might have different naming
            lib_files = list(build_dir.rglob("*.lib"))
            # Filter for yoga libraries (including variations)
            lib_files = [f for f in lib_files if "yoga" in f.name.lower() or "yogacore" in f.name.lower()]

    if not lib_files:
        log_error(f"Library not found in build directory: {build_dir}")
        log_error(f"Looking for: {platform_config['lib_name']}")
        # Debug: list what files are actually there
        all_libs = list(build_dir.rglob("*.a")) + list(build_dir.rglob("*.lib"))
        log_info(f"Available libraries: {[f.name for f in all_libs[:10]]}")
        sys.exit(1)

    lib_path = lib_files[0]
    log_info(f"Found library: {lib_path}")

    # Create target directory
    target_dir = LIBS_DIR / platform_name / arch
    target_dir.mkdir(parents=True, exist_ok=True)

    target_lib_path = target_dir / platform_config["lib_name"]

    # Remove old library
    if target_lib_path.exists():
        target_lib_path.unlink()

    # Copy new library
    shutil.copy2(lib_path, target_lib_path)
    log_success(f"Library installed to {target_lib_path}")

    # Copy headers (only .h files) from yoga source directory
    yoga_source_dir = YOGA_SUBMODULE_DIR / "yoga"
    target_include_dir = INCLUDE_DIR / "yoga"

    if yoga_source_dir.exists():
        # Remove old headers
        if target_include_dir.exists():
            shutil.rmtree(target_include_dir)

        # Copy only header files
        header_files = list(yoga_source_dir.rglob("*.h"))

        for header_file in header_files:
            # Create relative directory structure
            rel_path = header_file.relative_to(yoga_source_dir)
            target_file = target_include_dir / rel_path
            target_file.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(header_file, target_file)

        log_success(f"Headers installed to {target_include_dir}")
        log_info(f"Installed {len(header_files)} header files")
    else:
        log_warning(f"Yoga source directory not found: {yoga_source_dir}")

def build_current_platform():
    """Build for current platform only"""
    current_platform = platform.system().lower()
    current_arch = platform.machine().lower()

    # Normalize architecture names to Go standards
    if current_arch in ["x86_64", "amd64"]:
        if current_platform == "windows":
            current_arch = "x64"  # Windows uses x64, not amd64
        else:
            current_arch = "amd64"
    elif current_arch in ["aarch64", "arm64"]:
        current_arch = "arm64"

    if current_platform not in PLATFORMS:
        log_error(f"Unsupported platform: {current_platform}")
        return

    if current_arch not in PLATFORMS[current_platform]["architectures"]:
        log_error(f"Unsupported architecture: {current_arch} for platform {current_platform}")
        return

    build_dir = build_yoga_library(current_platform, current_arch)
    install_library(current_platform, current_arch, build_dir)

def build_all_platforms():
    """Build for all supported platforms"""
    log_info("Building for all platforms...")

    for platform_name in PLATFORMS:
        for arch in PLATFORMS[platform_name]["architectures"]:
            try:
                build_dir = build_yoga_library(platform_name, arch)
                install_library(platform_name, arch, build_dir)
            except Exception as e:
                log_error(f"Failed to build for {platform_name}-{arch}: {e}")
                continue

def main():
    parser = argparse.ArgumentParser(
        description="Cross-platform Yoga library builder and updater",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  %(prog)s --platform-only          # Build for current platform only
  %(prog)s --all-platforms          # Build for all supported platforms
  %(prog)s --clean                  # Clean build directories only
  %(prog)s --platform linux x86_64  # Build for specific platform/arch
        """
    )

    parser.add_argument("--platform-only", action="store_true",
                       help="Build for current platform only")
    parser.add_argument("--all-platforms", action="store_true",
                       help="Build for all supported platforms")
    parser.add_argument("--platform", nargs=2, metavar=("PLATFORM", "ARCH"),
                       help="Build for specific platform and architecture")
    parser.add_argument("--clean", action="store_true",
                       help="Clean build directories only")
    parser.add_argument("--skip-deps-check", action="store_true",
                       help="Skip dependency checking")

    args = parser.parse_args()

    print(f"{Colors.HEADER}Yoga Library Cross-Platform Builder{Colors.ENDC}")
    print("=" * 50)

    if not args.skip_deps_check:
        check_dependencies()

    if args.clean:
        clean_build_dirs()
        log_success("Build directories cleaned")
        return

    try:
        update_submodule()
        clean_build_dirs()

        if args.platform_only:
            build_current_platform()
        elif args.all_platforms:
            build_all_platforms()
        elif args.platform:
            platform_name, arch = args.platform
            if platform_name not in PLATFORMS:
                log_error(f"Unsupported platform: {platform_name}")
                sys.exit(1)
            if arch not in PLATFORMS[platform_name]["architectures"]:
                log_error(f"Unsupported architecture: {arch} for platform {platform_name}")
                sys.exit(1)

            build_dir = build_yoga_library(platform_name, arch)
            install_library(platform_name, arch, build_dir)
        else:
            log_warning("No build target specified. Use --platform-only, --all-platforms, or --platform PLATFORM ARCH")
            parser.print_help()
            return

        log_success("Yoga library build completed successfully!")

    except Exception as e:
        log_error(f"Build failed: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()