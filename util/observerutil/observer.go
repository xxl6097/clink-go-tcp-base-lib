package observerutil

import (
	"sync"
)

var instance *observerSingleton
var lock = &sync.Mutex{}

type observerSingleton struct {
	observers []Observer
}

func GetObserver() *observerSingleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &observerSingleton{
				observers: make([]Observer, 0),
			}
		}
	}
	return instance
}

// Subject 主题接口，定义了添加、删除和通知观察者的方法
type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify()
}

// Observer 观察者接口，定义了更新方法
type Observer interface {
	Update(data interface{})
}

// RegisterObserver 添加观察者
func (s *observerSingleton) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver 移除观察者
func (s *observerSingleton) Remove(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers 通知所有观察者
func (s *observerSingleton) Notify(data interface{}) {
	for _, _observer := range s.observers {
		_observer.Update(data)
	}
}
