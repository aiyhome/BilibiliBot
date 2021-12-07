package event

import (
	_ "fmt"
	"sync"
)

type Dispatcher struct {
	sync.RWMutex
	events map[string][]Listener
}

func NewDispatcher() *Dispatcher {
	dp := &Dispatcher{}
	dp.events = make(map[string][]Listener)
	return dp
}

func (self *Dispatcher) Attach(name string, listener Listener) {
	self.Lock()
	self.events[name] = append(self.events[name], listener)
	self.Unlock()
}

func (self *Dispatcher) Detach(name string, listener Listener) {
	self.Lock()
	var pos int
	if listeners, exist := self.events[name]; exist {
		for _, l := range listeners {
			if l.Equal(listener) {
				self.events[name] = append(listeners[:pos], listeners[pos+1:]...)
				if pos > 0 {
					pos = pos - 1
				}
				break
			}
			pos++
		}

	}
	self.Unlock()
}

func (self *Dispatcher) Event(name string, data interface{}) {
	if listeners, exist := self.events[name]; exist {
		evt := NewEvent(name, data)
		for _, l := range listeners {
			l.Exec(evt)
			if evt.Stoped {
				break
			}
		}
	}
}
