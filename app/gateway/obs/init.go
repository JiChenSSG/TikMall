package obs

import (
	"sync"

	"github.com/cloudwego/hertz/pkg/route"
)

var once sync.Once

var Hooks []route.CtxCallback

func Init() {
	once.Do(
		func() {
			Hooks = append(Hooks, initTracer(), initMetric())
			initLog()
		},
	)
}
