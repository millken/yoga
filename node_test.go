package yoga

import (
	"strings"
	"testing"
)

func TestNodeNewAPIs(t *testing.T) {
	// Test GetFlex method
	node := NewNode()
	defer node.Destroy()

	node.SetFlex(1.5)
	flex := node.GetFlex()
	if flex != 1.5 {
		t.Errorf("Expected flex to be 1.5, got %f", flex)
	}

	// Test HasBaselineFunc
	hasBaseline := node.HasBaselineFunc()
	if hasBaseline {
		t.Errorf("Expected HasBaselineFunc to be false by default")
	}

	// Test Finalize
	testNode := NewNode()
	testNode.Finalize() // Should not crash
}

func TestNodeClone(t *testing.T) {
	node := NewNode()
	defer node.Destroy()

	node.SetFlex(1.0)
	node.SetFlexGrow(2.0)

	// Test Clone
	clonedNode := node.Clone()
	if clonedNode == nil {
		t.Fatal("Clone returned nil")
	}
	defer clonedNode.Destroy()

	// Verify the cloned node has the same properties
	originalFlex := node.GetFlex()
	clonedFlex := clonedNode.GetFlex()

	if originalFlex != clonedFlex {
		t.Errorf("Expected cloned flex to be %f, got %f", originalFlex, clonedFlex)
	}
}

func TestNodeNewLayoutAPIs(t *testing.T) {
	config := NewConfig()
	defer config.Destroy()

	node := NewNodeWithConfig(config)
	defer node.Destroy()

	// Set some basic properties
	node.SetWidth(100)
	node.SetHeight(100)
	node.SetPadding(EdgeAll, 10)

	// Calculate layout
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)

	// Test new layout APIs
	layoutDir := node.GetLayoutDirection()
	if layoutDir != DirectionLTR {
		t.Errorf("Expected layout direction LTR, got %v", layoutDir)
	}

	// Test raw width/height APIs
	rawWidth := node.GetRawWidth()
	rawHeight := node.GetRawHeight()

	if rawWidth <= 0 {
		t.Errorf("Expected raw width > 0, got %f", rawWidth)
	}

	if rawHeight <= 0 {
		t.Errorf("Expected raw height > 0, got %f", rawHeight)
	}

	// Test GetDirtiedFunc (should return nil for default)
	dirtiedFunc := node.GetDirtiedFunc()
	if dirtiedFunc != nil {
		t.Errorf("Expected GetDirtiedFunc to return nil by default, got %v", dirtiedFunc)
	}

	// Test that computed layout still works
	computedWidth := node.GetComputedWidth()
	computedHeight := node.GetComputedHeight()

	if computedWidth <= 0 {
		t.Errorf("Expected computed width > 0, got %f", computedWidth)
	}

	if computedHeight <= 0 {
		t.Errorf("Expected computed height > 0, got %f", computedHeight)
	}
	htmlStr := node.ToYogaLayout()
	if len(htmlStr) == 0 {
		t.Errorf("Expected non-empty HTML string representation of node")
	}
}

func TestYogaConstantsAndUtils(t *testing.T) {
	// Test Zero constant
	if Zero != 0.0 {
		t.Errorf("Expected Zero to be 0.0, got %f", Zero)
	}

	// Test Undefined constant
	if !IsNaN(Undefined) {
		t.Errorf("Expected Undefined to be NaN")
	}

	// Test FloatIsUndefined function
	if !FloatIsUndefined(Undefined) {
		t.Errorf("Expected FloatIsUndefined(Undefined) to return true")
	}

	if FloatIsUndefined(0.0) {
		t.Errorf("Expected FloatIsUndefined(0.0) to return false")
	}

	if FloatIsUndefined(100.0) {
		t.Errorf("Expected FloatIsUndefined(100.0) to return false")
	}
}

func TestNodeMeasureFuncBasic(t *testing.T) {
	node := NewNode()
	defer node.Destroy()

	expected := Size{Width: 50, Height: 40}
	node.SetMeasureFunc(func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		return expected
	})

	node.CalculateLayout(Undefined, Undefined, DirectionLTR)

	if got := node.GetComputedWidth(); got != expected.Width {
		t.Errorf("expected width %v, got %v", expected.Width, got)
	}
	if got := node.GetComputedHeight(); got != expected.Height {
		t.Errorf("expected height %v, got %v", expected.Height, got)
	}
}

func TestNodeMeasureFuncReplaceAndUnset(t *testing.T) {
	node := NewNode()
	defer node.Destroy()

	node.SetMeasureFunc(func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		return Size{Width: 10, Height: 20}
	})
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 10 || h != 20 {
		t.Fatalf("first measure expected (10,20), got (%v,%v)", w, h)
	}

	// 替换回调后，再次布局应采用新回调结果
	node.SetMeasureFunc(func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		return Size{Width: 30, Height: 40}
	})
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 30 || h != 40 {
		t.Fatalf("replaced measure expected (30,40), got (%v,%v)", w, h)
	}

	// 取消回调并设置明确样式尺寸，应不再调用回调
	node.UnsetMeasureFunc()
	node.SetWidth(77)
	node.SetHeight(88)
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 77 || h != 88 {
		t.Fatalf("after unset measure expected (77,88), got (%v,%v)", w, h)
	}
}

func TestNodeContext(t *testing.T) {
	node := NewNode()
	defer node.Destroy()

	// Test basic string context
	testString := "Hello, Yoga!"
	node.SetContext(testString)

	retrievedContext := node.GetContext()
	if retrievedContext == nil {
		t.Fatal("Expected non-nil context, got nil")
	}

	retrievedString, ok := retrievedContext.(string)
	if !ok {
		t.Fatalf("Expected string context, got %T", retrievedContext)
	}

	if retrievedString != testString {
		t.Errorf("Expected context '%s', got '%s'", testString, retrievedString)
	}

	// Test integer context
	testInt := 42
	node.SetContext(testInt)

	retrievedIntContext := node.GetContext()
	if retrievedIntContext == nil {
		t.Fatal("Expected non-nil context, got nil")
	}

	retrievedInt, ok := retrievedIntContext.(int)
	if !ok {
		t.Fatalf("Expected int context, got %T", retrievedIntContext)
	}

	if retrievedInt != testInt {
		t.Errorf("Expected context %d, got %d", testInt, retrievedInt)
	}

	// Test struct context
	type TestStruct struct {
		Name    string
		Value   int
		Enabled bool
	}

	testStruct := TestStruct{
		Name:    "TestNode",
		Value:   123,
		Enabled: true,
	}

	node.SetContext(testStruct)

	retrievedStructContext := node.GetContext()
	if retrievedStructContext == nil {
		t.Fatal("Expected non-nil context, got nil")
	}

	retrievedStruct, ok := retrievedStructContext.(TestStruct)
	if !ok {
		t.Fatalf("Expected TestStruct context, got %T", retrievedStructContext)
	}

	if retrievedStruct.Name != testStruct.Name {
		t.Errorf("Expected Name '%s', got '%s'", testStruct.Name, retrievedStruct.Name)
	}
	if retrievedStruct.Value != testStruct.Value {
		t.Errorf("Expected Value %d, got %d", testStruct.Value, retrievedStruct.Value)
	}
	if retrievedStruct.Enabled != testStruct.Enabled {
		t.Errorf("Expected Enabled %v, got %v", testStruct.Enabled, retrievedStruct.Enabled)
	}

	// Test nil context
	node.SetContext(nil)
	nilContext := node.GetContext()
	if nilContext != nil {
		t.Errorf("Expected nil context after setting nil, got %v", nilContext)
	}

	// Test context with cloned node
	node.SetContext("original context")
	clonedNode := node.Clone()
	defer clonedNode.Destroy()

	clonedContext := clonedNode.GetContext()
	if clonedContext == nil {
		t.Fatal("Expected non-nil context on cloned node, got nil")
	}

	clonedString, ok := clonedContext.(string)
	if !ok {
		t.Fatalf("Expected string context on cloned node, got %T", clonedContext)
	}

	if clonedString != "original context" {
		t.Errorf("Expected cloned context 'original context', got '%s'", clonedString)
	}

	// Test that context survives layout calculations
	node.SetContext("layout test")
	node.SetWidth(100)
	node.SetHeight(100)
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)

	afterLayoutContext := node.GetContext()
	if afterLayoutContext == nil {
		t.Fatal("Expected non-nil context after layout, got nil")
	}

	afterLayoutString, ok := afterLayoutContext.(string)
	if !ok {
		t.Fatalf("Expected string context after layout, got %T", afterLayoutContext)
	}

	if afterLayoutString != "layout test" {
		t.Errorf("Expected context 'layout test' after layout, got '%s'", afterLayoutString)
	}
}

func TestNodeToYogaLayout(t *testing.T) {
	// Test classic mobile layout: Header, Content, Footer
	root := NewNode()
	defer root.Destroy()

	// Root container - full screen
	root.SetWidth(375)  // iPhone width
	root.SetHeight(667) // iPhone height
	// Note: flexDirection defaults to column, which is what we want for mobile layout

	// Header - 60px height
	header := NewNode()
	defer header.Destroy()
	header.SetHeight(60)
	header.SetFlexDirection(FlexDirectionRow)
	header.SetAlignItems(AlignCenter)
	header.SetPadding(EdgeHorizontal, 16)

	// Header title (flexible space)
	headerTitle := NewNode()
	defer headerTitle.Destroy()
	headerTitle.SetFlexGrow(1)

	// Header menu button
	headerMenu := NewNode()
	defer headerMenu.Destroy()
	headerMenu.SetWidth(44)
	headerMenu.SetHeight(44)

	// Content - flexible
	content := NewNode()
	defer content.Destroy()
	content.SetFlexGrow(1)
	content.SetPadding(EdgeAll, 16)
	content.SetFlexDirection(FlexDirectionColumn)

	// Content items
	for i := 0; i < 5; i++ {
		item := NewNode()
		defer item.Destroy()
		item.SetHeight(80)
		item.SetMargin(EdgeBottom, 12)
		content.InsertChild(item, uint32(i))
	}

	// Footer - 80px height
	footer := NewNode()
	defer footer.Destroy()
	footer.SetHeight(80)
	footer.SetFlexDirection(FlexDirectionRow)
	footer.SetJustifyContent(JustifySpaceAround)
	footer.SetAlignItems(AlignCenter)
	footer.SetPadding(EdgeHorizontal, 20)

	// Footer navigation items
	for i := 0; i < 4; i++ {
		navItem := NewNode()
		defer navItem.Destroy()
		navItem.SetWidth(50)
		navItem.SetHeight(50)
		footer.InsertChild(navItem, uint32(i))
	}

	// Assemble the header first
	header.InsertChild(headerTitle, 0)
	header.InsertChild(headerMenu, 1)

	// Assemble the layout
	root.InsertChild(header, 0)
	root.InsertChild(content, 1)
	root.InsertChild(footer, 2)

	// Calculate layout
	root.CalculateLayout(Undefined, Undefined, DirectionLTR)

	// Generate Yoga Layout format
	html := root.ToYogaLayout()

	// Verify the generated Yoga Layout format
	if !strings.Contains(html, "<Layout config={{useWebDefaults: false}}>") {
		t.Errorf("Expected Layout container")
	}
	if !strings.Contains(html, "</Layout>") {
		t.Errorf("Expected closing Layout tag")
	}

	// Check for main structure - note that column is default, so it may not appear in output
	// We expect at least one non-default flexDirection (row for header/footer)
	if !strings.Contains(html, "flexDirection: 'row'") {
		t.Errorf("Expected row flex direction for header or footer")
	}

	// Look for key structural elements
	hasRoot := strings.Contains(html, "width: 375")
	hasHeader := strings.Contains(html, "height: 60")
	hasContent := strings.Contains(html, "flex: 1")
	hasFooter := strings.Contains(html, "justifyContent: 'space-around'")

	if !hasRoot || !hasHeader || !hasContent || !hasFooter {
		t.Errorf("Expected complete mobile layout structure (root: %v, header: %v, content: %v, footer: %v)",
			hasRoot, hasHeader, hasContent, hasFooter)
	}
}

func TestNodeContextExtension(t *testing.T) {
	node := NewNode()
	defer node.Destroy()

	// Test that NodeContext structure works properly with MeasureFunc
	measureCount := 0
	testMeasureFunc := func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		measureCount++
		return Size{Width: 50, Height: 60}
	}

	// Set measure func
	node.SetMeasureFunc(testMeasureFunc)

	// Verify measure func is set
	if node.GetMeasureFunc() == nil {
		t.Fatal("Expected non-nil measure func, got nil")
	}

	// Test measure functionality works
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if measureCount == 0 {
		t.Error("Expected measure func to be called")
	}

	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 50 || h != 60 {
		t.Errorf("Expected size (50,60), got (%v,%v)", w, h)
	}

	// Test replace measure func
	replaceCount := 0
	replaceMeasureFunc := func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		replaceCount++
		return Size{Width: 80, Height: 90}
	}

	node.SetMeasureFunc(replaceMeasureFunc)
	if node.GetMeasureFunc() == nil {
		t.Fatal("Expected non-nil measure func after replacement, got nil")
	}

	// Test replacement works
	node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	if replaceCount == 0 {
		t.Error("Expected replaced measure func to be called")
	}

	if w, h := node.GetComputedWidth(), node.GetComputedHeight(); w != 80 || h != 90 {
		t.Errorf("Expected size (80,90), got (%v,%v)", w, h)
	}

	// Test unset measure func
	node.UnsetMeasureFunc()
	if node.GetMeasureFunc() != nil {
		t.Error("Expected nil measure func after unset")
	}
}

func BenchmarkMeasureFunc(b *testing.B) {
	node := NewNode()
	defer node.Destroy()

	node.SetMeasureFunc(func(w float32, wm MeasureMode, h float32, hm MeasureMode) Size {
		return Size{Width: 100, Height: 100}
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node.CalculateLayout(Undefined, Undefined, DirectionLTR)
	}
}

func BenchmarkNodeContextOperations(b *testing.B) {
	node := NewNode()
	defer node.Destroy()

	testMeasureFunc := func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		return Size{Width: 100, Height: 100}
	}

	b.Run("SetMeasureFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			node.SetMeasureFunc(testMeasureFunc)
		}
	})

	b.Run("GetMeasureFunc", func(b *testing.B) {
		node.SetMeasureFunc(testMeasureFunc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = node.GetMeasureFunc()
		}
	})
}
