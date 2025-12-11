package util

import (
	"runtime/debug"
	"time"

	"go-webmvc/pkg/logger"

	"go.uber.org/zap"
)

func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// SafeGo 启动一个 goroutine 并在内部 recover panic，防止整个进程因为未捕获的 panic 崩溃。
// 使用示例：
//
//	util.SafeGo(func(){
//	    // 异步任务逻辑
//	})
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				// 使用项目 logger 记录详细信息，便于排查
				logger.Error("goroutine panic recovered",
					zap.Any("panic", r),
					zap.ByteString("stack", stack),
				)
			}
		}()
		fn()
	}()
}
