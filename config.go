package yoga

type YGCloneNodeFunc func(
	oldNode *Node,
	owner *Node,
	childIndex uint32,
) *Node

type YGLogger func(
	config *Config,
	node *Node,
	level YGLogLevel,
	format string,
	args ...any,
) int

type Config struct {
	cloneNodeCallback_    YGCloneNodeFunc
	logger_               YGLogger
	useWebDefaults_       bool
	printTree_            bool
	experimentalFeatures_ EnumBitset
	errata_               YGErrata
	context_              any
	pointScaleFactor_     float32
}

func ConfigNew() *Config {
	return NewConfig(DefaultLogger)
}

func NewConfig(logger YGLogger) *Config {
	return &Config{
		cloneNodeCallback_:    nil,
		logger_:               logger,
		useWebDefaults_:       false,
		printTree_:            false,
		experimentalFeatures_: NewEnumBitset(),
		errata_:               0,
		pointScaleFactor_:     1.0,
	}
}

func (config *Config) SetUseWebDefaults(useWebDefaults bool) {
	config.useWebDefaults_ = useWebDefaults
}

func (config *Config) UseWebDefaults() bool {
	return config.useWebDefaults_
}

func (config *Config) setPrintTreeEnabled(printTree bool) {
	config.printTree_ = printTree
}

func (config *Config) shouldPrintTree() bool {
	return config.printTree_
}

func (config *Config) SetExperimentalFeatureEnabled(feature YGExperimentalFeature, enabled bool) {
	if enabled {
		config.experimentalFeatures_.Set(int(feature))
	} else {
		config.experimentalFeatures_.Reset(int(feature))
	}
}

func (config *Config) IsExperimentalFeatureEnabled(feature YGExperimentalFeature) bool {
	return config.experimentalFeatures_.Test(int(feature))
}

func (config *Config) GetEnabledExperiments() EnumBitset {
	return config.experimentalFeatures_
}

// setErrata
func (config *Config) setErrata(errata YGErrata) {
	config.errata_ = errata
}

// addErrata
func (config *Config) addErrata(errata YGErrata) {
	config.errata_ |= errata
}

// removeErrata
func (config *Config) removeErrata(errata YGErrata) {
	config.errata_ &= ^errata
}

// getErrata
func (config *Config) getErrata() YGErrata {
	return config.errata_
}

// hasErrata
func (config *Config) hasErrata(errata YGErrata) bool {
	return config.errata_&errata != YGErrataNone
}

// SetPointScaleFactor
func (config *Config) SetPointScaleFactor(pointScaleFactor float32) {
	config.pointScaleFactor_ = pointScaleFactor
}

// GetPointScaleFactor
func (config *Config) GetPointScaleFactor() float32 {
	return config.pointScaleFactor_
}

// SetContext
func (config *Config) SetContext(context any) {
	config.context_ = context
}

// GetContext
func (config *Config) GetContext() any {
	return config.context_
}

// SetLogger
func (config *Config) SetLogger(logger YGLogger) {
	config.logger_ = logger
}

// log
func (config *Config) log(node *Node, level YGLogLevel, format string, args ...any) {
	if config.logger_ != nil {
		config.logger_(config, node, level, format, args...)
	}
}

// setCloneNodeCallback
func (config *Config) setCloneNodeCallback(callback YGCloneNodeFunc) {
	config.cloneNodeCallback_ = callback
}

// cloneNode
func (config *Config) cloneNode(oldNode *Node, owner *Node, childIndex uint32) *Node {
	var clone *Node
	if config.cloneNodeCallback_ != nil {
		clone = config.cloneNodeCallback_(oldNode, owner, childIndex)
	}
	if clone == nil {
		clone = &Node{}
		*clone = *oldNode
		clone.setOwner(nil)
	}
	return clone
}

var defaultConfig Config

func getDefault() *Config {
	return &defaultConfig
}

func init() {
	defaultConfig = Config{
		logger_: DefaultLogger,
	}
}

func configUpdateInvalidatesLayout(oldConfig, newConfig *Config) bool {
	return oldConfig.getErrata() != newConfig.getErrata() ||
		oldConfig.GetEnabledExperiments() != newConfig.GetEnabledExperiments() ||
		oldConfig.GetPointScaleFactor() != newConfig.GetPointScaleFactor() ||
		oldConfig.UseWebDefaults() != newConfig.UseWebDefaults()
}
