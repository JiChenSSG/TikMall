package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-chassis/foundation/validator"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       string
	Kitex     Kitex     `yaml:"kitex"`
	Consul    Consul    `yaml:"consul"`
	Server    Server    `yaml:"server"`
	Telemetry Telemetry `yaml:"telemetry"`
	Mysql     Mysql     `yaml:"mysql"`
	Metrics   Metrics   `yaml:"metrics"`
	Order     Order     `yaml:"order"`
}

type Kitex struct {
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
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}

type Telemetry struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Dsn      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Metrics struct {
	Port int `yaml:"port"`
}

type Order struct {
	CancelTimeout int `yaml:"cancel_timeout"`
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
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
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
				klog.Error("load env error - %v", err)
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

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
