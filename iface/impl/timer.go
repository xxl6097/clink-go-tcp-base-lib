package impl

import (
	"github.com/xxl6097/clink-go-tcp-base-lib/iface"
	"github.com/xxl6097/go-glog/glog"
	"time"
)

type Timer struct {
	iface.ITimer
	ticker  *time.Ticker
	trigger chan bool
	son     func()
}

func New(duration time.Duration, _son func()) iface.ITimer {
	if _son == nil {
		glog.Fatal("son is nil")
	}
	return &Timer{
		ticker:  time.NewTicker(duration),
		trigger: make(chan bool),
		son:     _son,
	}
}
func (this *Timer) run() {
	if this.son == nil {
		return
	}
	this.son()
	for {
		select {
		case ch := <-this.trigger:
			glog.Info("停止 TimerPkg ", ch)
			if ch {
				return
			} else {
				//用户行为触发
				this.son()
			}
		case <-this.ticker.C:
			//定时行为触发
			this.son()
		}
	}
}

func (this *Timer) ReSetTimer(duration time.Duration) {
	this.ticker.Reset(duration)
}

func (this *Timer) Trigger() {
	this.trigger <- false
}
func (this *Timer) Start() {
	go this.run()
}
func (this *Timer) Stop() {
	this.trigger <- true
	this.ticker.Stop()
}
