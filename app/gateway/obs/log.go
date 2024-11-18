package obs

import (
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/server"
	hertzzap "github.com/hertz-contrib/logger/zap"
	hertzotelzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	"github.com/jichenssg/tikmall/app/gateway/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLog() {
	ioWriter := &lumberjack.Logger{
		Filename:   config.GetConf().Kitex.LogFileName,
		MaxSize:    config.GetConf().Kitex.LogMaxSize,
		MaxBackups: config.GetConf().Kitex.LogMaxBackups,
		MaxAge:     config.GetConf().Kitex.LogMaxAge,
	}

	var opts []hertzzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "online" {
		opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(ioWriter),
			zapcore.AddSync(os.Stdout),
		)
	} else {
		opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ioWriter),
			FlushInterval: time.Minute,
		}
		server.RegisterShutdownHook(func() {
			output.Sync() //nolint:errcheck
		})
	}

	log := hertzotelzap.NewLogger(hertzotelzap.WithLogger(hertzzap.NewLogger(opts...)))
	hlog.SetLogger(log)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(output)
}
