#include <stdint.h>
#include <stdarg.h>
#include <stdio.h>
#include <string.h>
#include <yoga/Yoga.h>

// Forward declaration of Go functions
extern int goBridgeLogger(YGConfigConstRef config, YGNodeConstRef node, YGLogLevel level, char* message);
extern YGSize goMeasureInvoke(YGNodeRef node, float width, YGMeasureMode widthMode, float height, YGMeasureMode heightMode);

// C bridge function that calls our Go logger
int c_bridge_yg_logger(YGConfigConstRef config, YGNodeConstRef node, YGLogLevel level, const char* format, va_list args) {
    // Simple buffer to hold formatted message
    char buffer[1024];
    vsnprintf(buffer, sizeof(buffer), format, args);
    return goBridgeLogger(config, node, level, buffer);
}