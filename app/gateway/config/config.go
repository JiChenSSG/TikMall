package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-chassis/foundation/validator"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       string
	Hertz     Hertz     `yaml:"hertz"`
	Consul    Consul    `yaml:"consul"`
	Server    Server    `yaml:"server"`
	Endpoint  Endpoint  `yaml:"endpoint"`
	Telemetry Telemetry `yaml:"telemetry"`
	Metrics   Metrics   `yaml:"metrics"`
}

type Hertz struct {
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
}

type Consul struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Server struct {
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`
}

type Endpoint struct {
	Auth string `yaml:"auth"`
}

type Telemetry struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Metrics struct {
	Port int `yaml:"port"`
}

var (
	conf     *Config
	confOnce sync.Once

	envOnce sync.Once
)

func GetConf() *Config {
	confOnce.Do(initConf)
	return conf
}

func initConf() {
	prefix := "config"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}
	conf.Env = GetEnv()
	pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	envOnce.Do(
		func() {
			err := godotenv.Load()
			if err != nil {
				hlog.Error("load env error - %v", err)
				panic(err)
			}
		},
	)

	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "dev"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
