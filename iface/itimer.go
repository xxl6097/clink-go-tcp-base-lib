package iface

import "time"

type ITimer interface {
	Run()
	ReSetTimer(duration time.Duration)
	Trigger()
	Start()
	Stop()
}
