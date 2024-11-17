package obs

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLog initializes the log.
// ioWriter is the output of the log.
// logLevel is the level of the log.
func InitLog(ioWriter io.Writer, logLevel klog.Level) {
	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "prod" {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(ioWriter),
			zapcore.AddSync(os.Stdout),
		)
	} else {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ioWriter),
			FlushInterval: time.Minute,
		}
	}

	// opts = append(opts,
	// 	otelzap.WithRecordStackTraceInSpan(true),
	// 	otelzap.WithTraceErrorSpanLevel(zapcore.WarnLevel),
	// )

	server.RegisterShutdownHook(func() {
		output.Sync() //nolint:errcheck
		log.Println("log sync done")
	})

	log := kitexzap.NewLogger(opts...)

	klog.SetLogger(log)
	klog.SetLevel(logLevel)
	klog.SetOutput(output)
}
