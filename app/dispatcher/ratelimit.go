// +build !confonly

package dispatcher

import (
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/features/ratelimit"
)

type LimitWriter struct {
	Limiter ratelimit.Limiter
	Writer  buf.Writer
}

func (w *LimitWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	return w.Writer.WriteMultiBuffer(mb)
}

func (w *LimitWriter) Close() error {
	return common.Close(w.Writer)
}

func (w *LimitWriter) Interrupt() {
	common.Interrupt(w.Writer)
}
