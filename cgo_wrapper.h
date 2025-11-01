#ifndef CGO_WRAPPER_H
#define CGO_WRAPPER_H

#include <stdint.h>
#include <yoga/Yoga.h>
#include <stdarg.h>

// Go functions that are called from C
extern YGSize goMeasureInvoke(YGNodeRef node, float width, YGMeasureMode widthMode, float height, YGMeasureMode heightMode);
extern int goBridgeLogger(YGConfigConstRef config, YGNodeConstRef node, YGLogLevel level, char* message);

// C functions that are called from Go
extern int c_bridge_yg_logger(YGConfigConstRef config, YGNodeConstRef node, YGLogLevel level, const char* format, va_list args);

#endif // CGO_WRAPPER_H