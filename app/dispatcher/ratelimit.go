// +build !confonly

package dispatcher

import (
	"time"
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/features/ratelimit"
)

type LimitWriter struct {
	Limiter ratelimit.Limiter
	Writer  buf.Writer
}

func (w *LimitWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	bucket := w.Limiter.GetBucket()
	if bucket == nil {
		return w.Writer.WriteMultiBuffer(mb)
	}
	size := int64(bucket.Rate())
	var waitSize int64 = 0
	for {
		if mb.IsEmpty() {
			break
		}
		mbl := int64(mb.Len())
		if mbl < size {
			waitSize = mbl
		} else {
			waitSize = size
		}
		//bucket.Wait(waitSize)
		//mbr, mbw := buf.SplitSize(mb, int32(waitSize))
		//err := w.Writer.WriteMultiBuffer(mbw)
		//if err != nil {
		//	return err
		//}
		//mb = mbr
		cnt := bucket.TakeAvailable(waitSize)
		if cnt == 0 {
			time.Sleep(50 * time.Millisecond)
			continue
		}
		mbr, mbw := buf.SplitSize(mb, int32(cnt))
		err := w.Writer.WriteMultiBuffer(mbw)
		if err != nil {
			return err
		}
		mb = mbr
	}
	return nil
}

func (w *LimitWriter) Close() error {
	return common.Close(w.Writer)
}

func (w *LimitWriter) Interrupt() {
	common.Interrupt(w.Writer)
}
