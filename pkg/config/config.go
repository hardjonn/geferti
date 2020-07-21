package config

import (
	"geferti/pkg/errs"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// C defines the whole application config.
type C struct {
	App    *App
	Logger *Logger
	DB     *DB
}

// App defines the application specific config.
type App struct {
	Name string
	Key  string
	Env  string
}

// Logger defines the logger config.
type Logger struct {
	Path   string
	Level  string
	Output string
}

// DB defines the database connection config.
type DB struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// New instantiates config
func New(confPath string, confName string, confType string) (*C, error) {
	viper.SetConfigName(confName)
	viper.SetConfigType(confType)
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()

	if err != nil {
		return nil, errs.E(errs.Op("config init"), errs.StatusIO, err)
	}

	viper.RegisterAlias("app_env", "env")

	viper.RegisterAlias("server_port", "port")
	viper.RegisterAlias("server_host", "host")

	viper.RegisterAlias("log_level", "log-level")
	viper.RegisterAlias("log_path", "log-path")
	viper.RegisterAlias("log_output", "log-output")

	viper.RegisterAlias("db_host", "db-host")
	viper.RegisterAlias("db_port", "db-port")
	viper.RegisterAlias("db_database", "db-database")
	viper.RegisterAlias("db_user", "db-user")
	viper.RegisterAlias("db_password", "db-password")

	flag.StringP("env", "e", "development", "application environment")

	flag.IntP("port", "p", 443, "server port")
	flag.StringP("host", "h", "localhost", "server host")

	flag.StringP("log-path", "L", "./log/geferti.log", "application log path")
	flag.StringP("log-level", "l", "debug", "application log level")
	flag.StringP("log-output", "o", "mixed", "application log output: mixed | console | file")

	flag.IntP("db-port", "P", 3306, "database port")
	flag.StringP("db-host", "H", "localhost", "database host")
	flag.StringP("db-user", "U", "root", "database user")
	flag.StringP("db-database", "B", "geferti", "database name")
	flag.StringP("db-password", "W", "root", "database password")

	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	return &C{
		Logger: newLogger(),
	}, nil
}

func newApp() *App {
	return &App{
		Name: viper.GetString("app_name"),
		Key:  viper.GetString("app_key"),
		Env:  viper.GetString("app_env"),
	}
}

func newLogger() *Logger {
	return &Logger{
		Path:   viper.GetString("log_path"),
		Level:  viper.GetString("log_level"),
		Output: viper.GetString("log_output"),
	}
}

func newDB() *DB {
	return &DB{
		Host:     viper.GetString("db_host"),
		Port:     viper.GetInt("db_port"),
		Database: viper.GetString("db_database"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
	}
}
