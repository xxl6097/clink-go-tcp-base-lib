package logutil

import (
	"context"
	"fmt"
	"github.com/aceld/zinx/zlog"
	"github.com/xxl6097/go-glog/glog"
)

// User-defined logging method
// The internal engine logging method of zinx can be reset by the logging method of its own business.
// In this example, fmt.Println is used.
// 用户自定义日志方式，
// 可以通过自身业务的日志方式，来重置zinx内部引擎的日志打印方式
// 本例以fmt.Println为例
type TcpServerLog struct{}

// Without context logging interface
func (l *TcpServerLog) InfoF(format string, v ...interface{}) {
	//fmt.Printf(format, v...)
	glog.StdGLog.Infof(format, v...)
}

func (l *TcpServerLog) ErrorF(format string, v ...interface{}) {
	//fmt.Printf(format, v...)
	glog.StdGLog.Errorf(format, v...)
}

func (l *TcpServerLog) DebugF(format string, v ...interface{}) {
	//fmt.Printf(format, v...)
	glog.StdGLog.Debugf(format, v...)
}

// Logging interface with context
func (l *TcpServerLog) InfoFX(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(ctx)
	//fmt.Printf(format, v...)
	glog.StdGLog.Infof(format, v...)
}

func (l *TcpServerLog) ErrorFX(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(ctx)
	//fmt.Printf(format, v...)
	glog.StdGLog.Errorf(format, v...)
}

func (l *TcpServerLog) DebugFX(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(ctx)
	//fmt.Printf(format, v...)
	glog.StdGLog.Debugf(format, v...)
}

func init() {
	zlog.SetLogger(new(TcpServerLog))
	glog.Debug("TcpServerLog init...")
}
