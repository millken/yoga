package yoga

/*
#include <stdint.h>
#include <stdarg.h>
#include <stdio.h>
#include <string.h>
#include <yoga/Yoga.h>

*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// NodeContext 存储节点的回调句柄，避免全局 map 查找
type NodeContext struct {
	measureHandle cgo.Handle
	// 可以扩展其他节点级回调：dirtiedHandle, baselineHandle 等
}

func wrapConfigRef(ref C.YGConfigConstRef) *Config {
	return &Config{
		config: C.YGConfigRef(ref),
	}
}

func wrapNodeRef[T C.YGNodeRef | C.YGNodeConstRef](ref T) *Node {
	if ref == nil {
		return nil
	}
	return &Node{
		node: C.YGNodeRef(ref),
	}
}

//export goMeasureInvoke
func goMeasureInvoke(node C.YGNodeRef, width C.float, widthMode C.YGMeasureMode, height C.float, heightMode C.YGMeasureMode) C.YGSize {
	if node == nil {
		var out C.YGSize
		out.width = 0
		out.height = 0
		return out
	}

	// 从节点的 context 中获取 measure handle
	h := getMeasureHandleByNode(node)
	if h == 0 {
		var out C.YGSize
		out.width = 0
		out.height = 0
		return out
	}

	// 安全地获取handle的值
	var mf MeasureFunc
	if func() bool {
		defer func() {
			if r := recover(); r != nil {
				// Handle无效，返回false
			}
		}()
		val, ok := h.Value().(MeasureFunc)
		if !ok || val == nil {
			return false
		}
		mf = val
		return true
	}() {
		// 成功获取到MeasureFunc
		size := mf(float32(width), MeasureMode(widthMode), float32(height), MeasureMode(heightMode))
		var out C.YGSize
		out.width = C.float(size.Width)
		out.height = C.float(size.Height)
		return out
	}

	// 无法获取有效的MeasureFunc，返回默认值
	var out C.YGSize
	out.width = 0
	out.height = 0
	return out
}

// 获取或创建节点的 NodeContext
func getNodeContext(node C.YGNodeRef) *NodeContext {
	if node == nil {
		return nil
	}

	contextPtr := C.YGNodeGetContext(node)
	if contextPtr != nil {
		return (*NodeContext)(contextPtr)
	}

	// 创建新的 NodeContext
	ctx := &NodeContext{}
	C.YGNodeSetContext(node, unsafe.Pointer(ctx))
	return ctx
}

// 清理节点的 NodeContext
func deleteNodeContext(node C.YGNodeRef) {
	if node == nil {
		return
	}

	contextPtr := C.YGNodeGetContext(node)
	if contextPtr != nil {
		ctx := (*NodeContext)(contextPtr)

		// 清理 measure handle
		if ctx.measureHandle != 0 {
			ctx.measureHandle.Delete()
		}
	}

	C.YGNodeSetContext(node, nil)
}

// 设置节点的 MeasureFunc
func setMeasureHandle(node C.YGNodeRef, measureFunc MeasureFunc) {
	ctx := getNodeContext(node)
	if ctx == nil {
		return
	}

	// 清理旧的 handle
	if ctx.measureHandle != 0 {
		ctx.measureHandle.Delete()
		ctx.measureHandle = 0
	}

	// 设置新的 handle
	if measureFunc != nil {
		ctx.measureHandle = cgo.NewHandle(measureFunc)
	}
}

// 删除节点的 MeasureFunc
func deleteMeasureHandle(node C.YGNodeRef) {
	if node == nil {
		return
	}

	contextPtr := C.YGNodeGetContext(node)
	if contextPtr != nil {
		ctx := (*NodeContext)(contextPtr)
		if ctx.measureHandle != 0 {
			ctx.measureHandle.Delete()
			ctx.measureHandle = 0
		}
	}
}

// 获取节点的 MeasureHandle（用于 goMeasureInvoke）
func getMeasureHandleByNode(node C.YGNodeRef) cgo.Handle {
	if node == nil {
		return 0
	}

	contextPtr := C.YGNodeGetContext(node)
	if contextPtr == nil {
		return 0
	}

	// 尝试将context转换为NodeContext
	ctx := (*NodeContext)(contextPtr)
	return ctx.measureHandle
}

// === ConfigContext 管理函数 ===

// 获取或创建配置的 ConfigContext
func getConfigContext(config C.YGConfigRef) *ConfigContext {
	if config == nil {
		return nil
	}

	contextPtr := C.YGConfigGetContext(config)
	if contextPtr != nil {
		return (*ConfigContext)(contextPtr)
	}

	// 创建新的 ConfigContext
	ctx := &ConfigContext{}
	C.YGConfigSetContext(config, unsafe.Pointer(ctx))
	return ctx
}

// 清理配置的 ConfigContext
func deleteConfigContext(config C.YGConfigRef) {
	if config == nil {
		return
	}

	contextPtr := C.YGConfigGetContext(config)
	if contextPtr != nil {
		ctx := (*ConfigContext)(contextPtr)

		// 清理 logger handle
		if ctx.loggerHandle != 0 {
			ctx.loggerHandle.Delete()
		}
	}

	C.YGConfigSetContext(config, nil)
}

// 设置配置的 LoggerFunc
func setLoggerHandle(config C.YGConfigRef, loggerFunc Logger) {
	ctx := getConfigContext(config)
	if ctx == nil {
		return
	}

	// 清理旧的 handle
	if ctx.loggerHandle != 0 {
		ctx.loggerHandle.Delete()
		ctx.loggerHandle = 0
	}

	// 设置新的 handle
	if loggerFunc != nil {
		ctx.loggerHandle = cgo.NewHandle(loggerFunc)
	}
}

// 删除配置的 LoggerFunc
func deleteLoggerHandle(config C.YGConfigRef) {
	if config == nil {
		return
	}

	contextPtr := C.YGConfigGetContext(config)
	if contextPtr != nil {
		ctx := (*ConfigContext)(contextPtr)
		if ctx.loggerHandle != 0 {
			ctx.loggerHandle.Delete()
			ctx.loggerHandle = 0
		}
	}
}

// 获取配置的 LoggerHandle
func getLoggerHandleByConfig(config C.YGConfigRef) cgo.Handle {
	if config == nil {
		return 0
	}

	contextPtr := C.YGConfigGetContext(config)
	if contextPtr != nil {
		ctx := (*ConfigContext)(contextPtr)
		return ctx.loggerHandle
	}

	return 0
}
