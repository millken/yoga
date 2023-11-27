package yoga_test

import (
	"testing"

	"github.com/millken/yoga"
	"github.com/stretchr/testify/require"
)

func TestMaina(t *testing.T) {
	r := require.New(t)
	a := struct {
		Foo int32
	}{
		Foo: 0x1234ABCD,
	}
	t.Logf("%p", &a)
	b := &a
	t.Logf("%p", b)
	b.Foo = 0x1234
	r.Equal(&a, b)
	t.Logf("%x", a.Foo)

}
func TestConfig(t *testing.T) {
	r := require.New(t)
	config := yoga.ConfigNew()
	r.NotNil(config)
	r.False(config.UseWebDefaults())
	r.Equal(config.GetPointScaleFactor(), float32(1.0))
	r.False(config.IsExperimentalFeatureEnabled(yoga.YGExperimentalFeatureWebFlexBasis))
	r.False(config.IsExperimentalFeatureEnabled(yoga.YGExperimentalFeatureAbsolutePercentageAgainstPaddingEdge))
	config.SetUseWebDefaults(true)
	config.SetPointScaleFactor(2.0)
	config.SetExperimentalFeatureEnabled(yoga.YGExperimentalFeatureWebFlexBasis, true)
	r.Nil(config.GetContext())

	r.True(config.UseWebDefaults())
	r.Equal(config.GetPointScaleFactor(), float32(2.0))
	r.True(config.IsExperimentalFeatureEnabled(yoga.YGExperimentalFeatureWebFlexBasis))
	r.False(config.IsExperimentalFeatureEnabled(yoga.YGExperimentalFeatureAbsolutePercentageAgainstPaddingEdge))
	ctx := &struct {
		Foo int32
	}{
		Foo: 0x1234ABCD,
	}
	config.SetContext(ctx)
	r.Equal(config.GetContext(), ctx)
}
