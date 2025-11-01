package yoga

import (
	"testing"

	"github.com/dnsoa/go/assert"
)

func TestConfig(t *testing.T) {
	r := assert.New(t)
	c := NewConfig()
	defer c.Destroy()
	c.SetPointScaleFactor(1.0)
	c.SetErrata(ErrataNone)
	r.Equal(ErrataNone, c.GetErrata())
	c.SetExperimentalFeatureEnabled(ExperimentalFeatureWebFlexBasis, true)
	r.True(c.IsExperimentalFeatureEnabled(ExperimentalFeatureWebFlexBasis))
	c.SetExperimentalFeatureEnabled(ExperimentalFeatureWebFlexBasis, false)
	r.False(c.IsExperimentalFeatureEnabled(ExperimentalFeatureWebFlexBasis))
	r.False(c.UseWebDefaults())
	c.SetUseWebDefaults(true)
	r.True(c.UseWebDefaults())
	c.SetPointScaleFactor(2.0)
	r.Equal(float32(2.0), c.PointScaleFactor())

	// 测试 Logger 功能
	c.SetLogger(vlog)
	r.NotNil(c.GetLogger())
}

func TestConfigLogger(t *testing.T) {
	c := NewConfig()
	defer c.Destroy()

	// Test basic logger functionality
	var logMessages []struct {
		config  *Config
		node    *Node
		level   LogLevel
		message string
	}

	testLogger := func(config *Config, node *Node, level LogLevel, message string) int {
		logMessages = append(logMessages, struct {
			config  *Config
			node    *Node
			level   LogLevel
			message string
		}{
			config:  config,
			node:    node,
			level:   level,
			message: message,
		})
		return 1 // Return success
	}

	// Set logger function
	c.SetLogger(testLogger)

	// Verify logger was set
	retrievedLogger := c.GetLogger()
	if retrievedLogger == nil {
		t.Fatal("Expected non-nil logger, got nil")
	}

	// Test logger replacement
	newTestLogger := func(config *Config, node *Node, level LogLevel, message string) int {
		return 2 // Different return value to indicate replacement
	}

	c.SetLogger(newTestLogger)
	retrievedLogger = c.GetLogger()
	if retrievedLogger == nil {
		t.Fatal("Expected non-nil logger after replacement, got nil")
	}

	// Test unset logger
	c.UnsetLogger()
	retrievedLogger = c.GetLogger()
	if retrievedLogger != nil {
		t.Errorf("Expected nil logger after unset, got %v", retrievedLogger)
	}

	// Test setting nil logger directly
	c.SetLogger(nil)
	retrievedLogger = c.GetLogger()
	if retrievedLogger != nil {
		t.Errorf("Expected nil logger after setting nil, got %v", retrievedLogger)
	}
}

func TestConfigContextWithNodeLogger(t *testing.T) {
	c := NewConfig()
	defer c.Destroy()
	node := NewNode()
	defer node.Destroy()

	// Test that both Config Logger and Node MeasureFunc can coexist
	measureCount := 0
	testMeasureFunc := func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		measureCount++
		return Size{Width: 50, Height: 60}
	}

	logCount := 0
	testLoggerFunc := func(config *Config, node *Node, level LogLevel, message string) int {
		logCount++
		return 1
	}

	// Set both callbacks
	node.SetMeasureFunc(testMeasureFunc)
	c.SetLogger(testLoggerFunc)

	// Verify both are set
	if node.GetMeasureFunc() == nil {
		t.Fatal("Expected non-nil measure func, got nil")
	}
	if c.GetLogger() == nil {
		t.Fatal("Expected non-nil logger func, got nil")
	}

	// Test measure functionality still works
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if measureCount == 0 {
		t.Error("Expected measure func to be called")
	}

	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 50 || h != 60 {
		t.Errorf("Expected size (50,60), got (%v,%v)", w, h)
	}

	// Test unset measure func while keeping logger
	node.UnsetMeasureFunc()
	if node.GetMeasureFunc() != nil {
		t.Error("Expected nil measure func after unset")
	}
	if c.GetLogger() == nil {
		t.Error("Expected logger func to remain after measure func unset")
	}

	// Test unset logger func
	c.UnsetLogger()
	if c.GetLogger() != nil {
		t.Error("Expected nil logger func after unset")
	}
}
