package osutil

import (
	"github.com/xxl6097/go-glog/glog"
	"runtime"
	"strings"
)

var debug *bool

func IsDebug() bool {
	glog.Debug("Debug mode", *debug)
	if debug == nil {
		return false
	} else {
		return *debug
	}
}

func SetDebug(_debug bool) {
	debug = &_debug
}

func IsMacOs() bool {
	if strings.Compare(runtime.GOOS, "darwin") == 0 {
		return true
	}
	return false
}

func IsLinux() bool {
	if strings.Compare(runtime.GOOS, "linux") == 0 {
		return true
	}
	return false
}

func IsWindows() bool {
	if strings.Compare(runtime.GOOS, "windows") == 0 {
		return true
	}
	return false
}

func IsFreebsd() bool {
	if strings.Compare(runtime.GOOS, "freebsd") == 0 {
		return true
	}
	return false
}

func IsOpenbsd() bool {
	if strings.Compare(runtime.GOOS, "openbsd") == 0 {
		return true
	}
	return false
}

func IsNetbsd() bool {
	if strings.Compare(runtime.GOOS, "netbsd") == 0 {
		return true
	}
	return false
}

func IsDragonfly() bool {
	if strings.Compare(runtime.GOOS, "dragonfly") == 0 {
		return true
	}
	return false
}

func IsAndroid() bool {
	if strings.Compare(runtime.GOOS, "android") == 0 {
		return true
	}
	return false
}
