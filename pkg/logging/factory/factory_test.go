package factory

import (
	"testing"

	"github.com/hardjonn/geferti/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnLogFactoryBuilder(t *testing.T) {
	for handler, factory := range logFactoryBuilderMap {
		f, err := GetLogFactoryBuilder(handler)
		assert.Equal(t, factory, f)
		assert.Nil(t, err)
	}

	f, err := GetLogFactoryBuilder("some_random_handler")
	assert.Equal(t, f, nil)
	assert.Error(t, err)
}

func TestShouldBuildLog(t *testing.T) {
	c := &config.Logger{
		Path:    "/tmp/geferti/geferti.log",
		Level:   "info",
		Output:  "file",
		Handler: "ZEROLOG",
	}

	l, err := Build(c)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

func TestShouldNotBuildLog(t *testing.T) {
	c := &config.Logger{
		Path:    "/tmp/geferti/geferti.log",
		Level:   "info",
		Output:  "file",
		Handler: "SOME_UNDEFINED_LOGGER",
	}

	l, err := Build(c)
	assert.Error(t, err)
	assert.Nil(t, l)
}
