package logger

import (
	"geferti/pkg/config"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type logLevelTest struct {
	levelIn  string
	levelOut zerolog.Level
}

func TestLogLevel(t *testing.T) {
	logLevelTests := []logLevelTest{
		{"debug", zerolog.DebugLevel},
		{"info", zerolog.InfoLevel},
		{"warn", zerolog.WarnLevel},
		{"error", zerolog.ErrorLevel},
		{"fatal", zerolog.FatalLevel},
		{"trace", zerolog.TraceLevel},
		{"panic", zerolog.PanicLevel},
		{"none", zerolog.NoLevel},
		{"non existing", zerolog.ErrorLevel},
	}

	for _, test := range logLevelTests {
		level := logLevel(test.levelIn)
		assert.Equal(t, level, test.levelOut, "log level should match")

		level = logLevel(strings.ToUpper(test.levelIn))
		assert.Equal(t, level, test.levelOut, "log level should match and be case insensitive")
	}
}

func TestMkDir(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)

	path := tmpDir + "/test/log"

	err = makeLogDir(path)
	assert.NoError(t, err, "should create a deep nested folder")

	err = makeLogDir(path)
	assert.NoError(t, err, "should return no error if the folder already exists")

	_, err = os.Stat(path)
	assert.NoError(t, err, "folder should exist")
}

type outWritersTest struct {
	output string
	wCount int
}

func TestOutputWriters(t *testing.T) {
	outWritersTests := []outWritersTest{
		{"file", 1},
		{"console", 1},
		{"mixed", 2},
		{"invalid", 0},
	}

	path := "/tmp/geferti/geferti.log"

	for _, test := range outWritersTests {
		outputWritersCaseIns(t, test.output, path, test.wCount)
		outputWritersCaseIns(t, strings.ToUpper(test.output), path, test.wCount)
	}
}

func outputWritersCaseIns(t *testing.T, output string, path string, wCount int) {
	writers, err := outputWriters(output, path)

	if wCount == 0 {
		assert.Error(t, err, "should return an error for invalid output")
		return
	}

	assert.Equal(t, wCount, len(writers), "amount of writers should match")
}

type loggerTest struct {
	config  config.Logger
	isError bool
}

func TestLogger(t *testing.T) {
	loggerTests := []loggerTest{
		{config.Logger{Path: "/tmp/geferti/geferti.log", Level: "info", Output: "file"}, false},
		{config.Logger{Path: "/tmp/geferti/geferti.log", Level: "info", Output: "console"}, false},
		{config.Logger{Path: "/tmp/geferti/geferti.log", Level: "info", Output: "mixed"}, false},
		{config.Logger{Path: "/tmp/geferti/geferti.log", Level: "info", Output: "invalid"}, true},
		{config.Logger{Path: "/etc/geferti/geferti.log", Level: "info", Output: "mixed"}, true},
	}

	for _, test := range loggerTests {
		_, err := New(&test.config)
		if test.isError {
			assert.Error(t, err, "should return an error")
		} else {
			assert.NoError(t, err, "should not return any errors")
		}
	}
}
