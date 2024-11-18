package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-chassis/foundation/validator"
	"github.com/casbin/casbin/v2"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"

	"github.com/joho/godotenv"
)

var Enforcer *casbin.Enforcer

type Config struct {
	Env       string
	Kitex     Kitex     `yaml:"kitex"`
	Consul    Consul    `yaml:"consul"`
	Server    Server    `yaml:"server"`
	Telemetry Telemetry `yaml:"telemetry"`
	Metrics   Metrics   `yaml:"metrics"`
	Redis     Redis     `yaml:"redis"`
	Token     Token     `yaml:"token"`
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

type Metrics struct {
	Port int `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type Token struct {
	AccessExpire  int `yaml:"access_expire"`
	RefreshExpire int `yaml:"refresh_expire"`
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

	casbinModelPath := filepath.Join(prefix, "model.conf")
	casbinPolicyPath := filepath.Join(prefix, filepath.Join(GetEnv(), "policy.csv"))
	Enforcer, err = casbin.NewEnforcer(casbinModelPath, casbinPolicyPath)
	if err != nil {
		klog.Error("load casbin error - %v", err)
		panic(err)
	}
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
