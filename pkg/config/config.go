package config

import (
	"geferti/pkg/errs"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Logger *Logger
}

// Logger ...
type Logger struct {
	Path   string
	Level  string
	Output string
}

// New instantiates config
func New(confPath string, confName string, confType string) (*Config, error) {
	viper.SetConfigName(confName)
	viper.SetConfigType(confType)
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()

	if err != nil {
		return nil, errs.E(errs.Op("config init"), errs.StatusIO, err)
	}

	viper.RegisterAlias("server_port", "port")
	viper.RegisterAlias("server_host", "host")
	viper.RegisterAlias("log_level", "log-level")
	viper.RegisterAlias("log_path", "log-path")
	viper.RegisterAlias("log_output", "log-output")

	flag.IntP("port", "P", 443, "server port")
	flag.StringP("host", "h", "localhost", "server host")
	flag.StringP("log-level", "l", "debug", "application log level")
	flag.StringP("log-path", "L", "/var/log/geferti/geferti.log", "application log path")
	flag.StringP("log-output", "o", "mixed", "application log output: mixed | console | file")

	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	return &Config{
		Logger: newLogger(),
	}, nil
}

func newLogger() *Logger {
	return &Logger{
		Path:   viper.GetString("log_path"),
		Level:  viper.GetString("log_level"),
		Output: viper.GetString("log_output"),
	}
}
