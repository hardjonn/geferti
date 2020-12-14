package zerolog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/errs"
	"github.com/hardjonn/geferti/pkg/logging"
	"github.com/rs/zerolog"
)

// New creates a new looger instance
func New(config *config.Logger) (logging.Logger, error) {
	level := logLevel(config.Level)
	zerolog.SetGlobalLevel(level)

	outputs, err := outputWriters(config.Output, config.Path)
	if err != nil {
		return nil, errs.E(errs.Op("logger.New"), err)
	}

	multi := zerolog.MultiLevelWriter(outputs...)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.
		Debug().
		Str("output", config.Output).
		Str("path", config.Path).
		Str("handler", config.Handler).
		Msg("logger initialized")

	return &loggerWrapper{logger, &logging.Fields{}}, nil
}

func outputWriters(output string, path string) ([]io.Writer, error) {
	var writers []io.Writer

	output = strings.ToUpper(output)

	if output == "FILE" || output == "MIXED" {
		fileWriter, err := fileWriter(path)
		if err != nil {
			return nil, errs.E(errs.Op("outputWriters.file"), err)
		}

		writers = append(writers, fileWriter)
	}

	if output == "CONSOLE" || output == "MIXED" {
		consoleWriter, _ := consoleWriter()
		writers = append(writers, consoleWriter)
	}

	if len(writers) == 0 {
		return nil, errs.E(errs.Op("outputWriters.output"), errs.StatusInvalid, "invalid output")
	}

	return writers, nil
}

func consoleWriter() (zerolog.ConsoleWriter, error) {
	return zerolog.ConsoleWriter{Out: os.Stdout}, nil
}

func fileWriter(path string) (io.Writer, error) {
	dirPath := filepath.Dir(path)

	if err := makeLogDir(dirPath); err != nil {
		return nil, errs.E(errs.Op("fileWriter.makeDir"), errs.StatusIO, err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return nil, errs.E(errs.Op("fileWriter.openFile"), errs.StatusIO, fmt.Errorf("path: %s. error: %v", path, err))
	}

	return file, nil
}

func makeLogDir(path string) error {
	err := os.MkdirAll(path, 0700)

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}

func logLevel(level string) zerolog.Level {
	level = strings.ToUpper(level)

	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "TRACE":
		return zerolog.TraceLevel
	case "PANIC":
		return zerolog.PanicLevel
	case "NONE":
		return zerolog.NoLevel
	}

	return zerolog.ErrorLevel
}
