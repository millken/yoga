package yoga

type YGCloneNodeFunc func(
	oldNode *YGNode,
	owner *YGNode,
	childIndex uint32,
) *YGNode

type YGLogger func(
	config *YGConfig,
	node *YGNode,
	level YGLogLevel,
	format string,
	args ...interface{},
) int

type YGConfig struct {
	cloneNodeCallback_    YGCloneNodeFunc
	logger_               YGLogger
	useWebDefaults_       bool
	printTree_            bool
	experimentalFeatures_ EnumBitset
	errata_               YGErrata
	context_              interface{}
	pointScaleFactor_     float32
}

func NewConfig(logger YGLogger) *YGConfig {
	return &YGConfig{
		cloneNodeCallback_:    nil,
		logger_:               logger,
		useWebDefaults_:       false,
		printTree_:            false,
		experimentalFeatures_: NewEnumBitset(),
		errata_:               0,
		pointScaleFactor_:     1.0,
	}
}

func (config *YGConfig) setUseWebDefaults(useWebDefaults bool) {
	config.useWebDefaults_ = useWebDefaults
}

func (config *YGConfig) useWebDefaults() bool {
	return config.useWebDefaults_
}

func (config *YGConfig) setPrintTreeEnabled(printTree bool) {
	config.printTree_ = printTree
}

func (config *YGConfig) shouldPrintTree() bool {
	return config.printTree_
}

func (config *YGConfig) setExperimentalFeatureEnabled(feature YGExperimentalFeature, enabled bool) {
	if enabled {
		config.experimentalFeatures_.Set(int(feature))
	} else {
		config.experimentalFeatures_.Reset(int(feature))
	}
}

func (config *YGConfig) isExperimentalFeatureEnabled(feature YGExperimentalFeature) bool {
	return config.experimentalFeatures_.Test(int(feature))
}

func (config *YGConfig) getEnabledExperiments() EnumBitset {
	return config.experimentalFeatures_
}

// setErrata
func (config *YGConfig) setErrata(errata YGErrata) {
	config.errata_ = errata
}

// addErrata
func (config *YGConfig) addErrata(errata YGErrata) {
	config.errata_ |= errata
}

// removeErrata
func (config *YGConfig) removeErrata(errata YGErrata) {
	config.errata_ &= ^errata
}

// getErrata
func (config *YGConfig) getErrata() YGErrata {
	return config.errata_
}

// hasErrata
func (config *YGConfig) hasErrata(errata YGErrata) bool {
	return config.errata_&errata != YGErrataNone
}

// setPointScaleFactor
func (config *YGConfig) setPointScaleFactor(pointScaleFactor float32) {
	config.pointScaleFactor_ = pointScaleFactor
}

// getPointScaleFactor
func (config *YGConfig) getPointScaleFactor() float32 {
	return config.pointScaleFactor_
}

// setContext
func (config *YGConfig) setContext(context interface{}) {
	config.context_ = context
}

// getContext
func (config *YGConfig) getContext() interface{} {
	return config.context_
}

// setLogger
func (config *YGConfig) setLogger(logger YGLogger) {
	config.logger_ = logger
}

// log
func (config *YGConfig) log(node *YGNode, level YGLogLevel, format string, args ...interface{}) {
	if config.logger_ != nil {
		config.logger_(config, node, level, format, args...)
	}
}

// setCloneNodeCallback
func (config *YGConfig) setCloneNodeCallback(callback YGCloneNodeFunc) {
	config.cloneNodeCallback_ = callback
}

// cloneNode
func (config *YGConfig) cloneNode(oldNode *YGNode, owner *YGNode, childIndex uint32) *YGNode {
	var clone *YGNode
	if config.cloneNodeCallback_ != nil {
		clone = config.cloneNodeCallback_(oldNode, owner, childIndex)
	}
	if clone == nil {
		clone = &YGNode{}
		*clone = *oldNode
		clone.setOwner(nil)
	}
	return clone
}

var defaultConfig YGConfig

func getDefault() *YGConfig {
	return &defaultConfig
}

func init() {
	defaultConfig = YGConfig{
		logger_: DefaultLogger,
	}
}

func configUpdateInvalidatesLayout(oldConfig, newConfig *YGConfig) bool {
	return oldConfig.getErrata() != newConfig.getErrata() ||
		oldConfig.getEnabledExperiments() != newConfig.getEnabledExperiments() ||
		oldConfig.getPointScaleFactor() != newConfig.getPointScaleFactor() ||
		oldConfig.useWebDefaults() != newConfig.useWebDefaults()
}
