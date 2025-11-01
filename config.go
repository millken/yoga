package yoga

/*
#include "cgo_wrapper.h"

*/
import "C"
import (
	"log"
	"runtime"
	"runtime/cgo"
)

// ConfigContext 存储 Config 的回调句柄，类似于 NodeContext
type ConfigContext struct {
	loggerHandle cgo.Handle
	// 可以扩展其他配置级回调
}

// Config 包装YGConfigRef
type Config struct {
	config C.YGConfigRef
}

// NewConfig 创建一个新的配置对象
func NewConfig() *Config {
	c := &Config{
		config: C.YGConfigNew(),
	}
	runtime.SetFinalizer(c, (*Config).Destroy)
	return c
}

// Destroy 释放配置资源
func (c *Config) Destroy() {
	if c.config != nil {
		// 清理 ConfigContext
		deleteConfigContext(c.config)
		C.YGConfigFree(c.config)
		c.config = nil
		runtime.SetFinalizer(c, nil)
	}
}

// SetExperimentalFeatureEnabled 启用或禁用实验性功能
func (c *Config) SetExperimentalFeatureEnabled(feature ExperimentalFeature, enabled bool) {
	if c.config != nil {
		C.YGConfigSetExperimentalFeatureEnabled(c.config, C.YGExperimentalFeature(feature), C.bool(enabled))
	}
}

// IsExperimentalFeatureEnabled 检查实验性功能是否启用
func (c *Config) IsExperimentalFeatureEnabled(feature ExperimentalFeature) bool {
	if c.config != nil {
		return bool(C.YGConfigIsExperimentalFeatureEnabled(c.config, C.YGExperimentalFeature(feature)))
	}
	return false
}

// SetPointScaleFactor 设置点缩放因子
func (c *Config) SetPointScaleFactor(pixelsInPoint float32) {
	if c.config != nil {
		C.YGConfigSetPointScaleFactor(c.config, C.float(pixelsInPoint))
	}
}

// PointScaleFactor 获取点缩放因子
func (c *Config) PointScaleFactor() float32 {
	if c.config != nil {
		return float32(C.YGConfigGetPointScaleFactor(c.config))
	}
	return 0.0
}

// SetErrata 设置错误处理模式
func (c *Config) SetErrata(errata Errata) {
	if c.config != nil {
		C.YGConfigSetErrata(c.config, C.YGErrata(errata))
	}
}

// GetErrata 获取当前的错误处理模式
func (c *Config) GetErrata() Errata {
	if c.config != nil {
		return Errata(C.YGConfigGetErrata(c.config))
	}
	return ErrataClassic // 默认返回经典模式
}

// SetUseWebDefaults 设置是否使用Web默认值
func (c *Config) SetUseWebDefaults(useWebDefaults bool) {
	if c.config != nil {
		C.YGConfigSetUseWebDefaults(c.config, C.bool(useWebDefaults))
	}
}

// UseWebDefaults 检查是否使用Web默认值
func (c *Config) UseWebDefaults() bool {
	if c.config != nil {
		return bool(C.YGConfigGetUseWebDefaults(c.config))
	}
	return false
}

// 内部方法，供Node使用
func (c *Config) ref() C.YGConfigRef {
	return c.config
}

type Logger func(config *Config, node *Node, level LogLevel, message string) int

// SetLogger 设置日志回调函数
func (c *Config) SetLogger(logger Logger) {
	if c.config == nil {
		return
	}

	if logger == nil {
		// 取消回调并清理句柄
		C.YGConfigSetLogger(c.config, nil)
		deleteLoggerHandle(c.config)
		return
	}

	// 存储 logger 句柄，并设置 C 层回调
	setLoggerHandle(c.config, logger)
	C.YGConfigSetLogger(c.config, C.YGLogger(C.c_bridge_yg_logger))
}

// GetLogger 获取当前日志回调函数
func (c *Config) GetLogger() Logger {
	if c.config == nil {
		return nil
	}

	// 从 ConfigContext 中获取 logger handle
	handle := getLoggerHandleByConfig(c.config)
	if handle == 0 {
		return nil
	}
	if logger, ok := handle.Value().(Logger); ok {
		return logger
	}
	return nil
}

// UnsetLogger 取消日志回调函数
func (c *Config) UnsetLogger() {
	if c.config != nil {
		C.YGConfigSetLogger(c.config, nil)
		// 清理 logger handle，但保留 ConfigContext 结构
		deleteLoggerHandle(c.config)
	}
}

//export goBridgeLogger
func goBridgeLogger(
	config C.YGConfigConstRef,
	node C.YGNodeConstRef,
	level C.YGLogLevel,
	message *C.char,
) C.int {
	// 将C字符串转换为Go字符串
	goMessage := C.GoString(message)

	// 从配置中获取 logger
	configHandle := getLoggerHandleByConfig(C.YGConfigRef(config))
	if configHandle != 0 {
		if logger, ok := configHandle.Value().(Logger); ok {
			result := logger(
				wrapConfigRef(config),
				wrapNodeRef(node),
				LogLevel(level),
				goMessage,
			)
			return C.int(result)
		}
	}

	// 如果没有设置logger，默认打印
	log.Printf("[Yoga %d] %s", level, goMessage)
	return 0
}

func vlog(config *Config,
	node *Node,
	level LogLevel,
	message string) int {
	log.Printf("%s", message)
	return 0
}

// // SetLoggerFunc 设置日志回调函数
// func (c *Config) SetLoggerFunc(loggerFunc func(level LogLevel, message string)) {
// 	if c.config != nil {
// 		C.YGConfigSetLogger(c.config, C.YGLogger(C.YGLoggerFunc(C.loggerFunc)))
// 	}
// }
