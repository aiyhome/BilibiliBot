package event

import (
    "reflect"
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

func (dp *Dispatcher) Attach(name string, listener Listener) {
    dp.Lock()
    dp.events[name] = append(dp.events[name], listener)
    dp.Unlock()
}

func (dp *Dispatcher) Detach(name string, listener Listener) {
    dp.Lock()
    var pos int
    if listeners, exist = dp.events[name]; exist {
        for _, l := range listeners {
            if l == listener {
                dp.events[name] = append(slice[:pos], slice[pos+1:]...)
                if pos > 0 {
                    pos = pos - 1
                }
                break
            }
            pos++
        }

    }
    dp.Unlock()
}

func (dp *Dispatcher) Event(name string, data interface{}) {
    if listeners, exist = dp.events[name]; exist {
        for _, l := range listeners {
            l.RunWith(data)
        }
    }
}
