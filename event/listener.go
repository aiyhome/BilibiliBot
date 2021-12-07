package event

type handleFunc func(evt *Event, data ...interface{})
type extData interface{}

type Listener struct {
	Func handleFunc
	Data extData
}

func NewListener(callback handleFunc, data ...interface{}) *Listener {
	l := &Listener{}
	if len(data) > 0 {
		l.Set(callback, data[0])
	} else {
		l.Set(callback, nil)
	}
	return l
}

func (self *Listener) Set(callback handleFunc, data interface{}) {
	self.Func = callback
	self.Data = data
}

func (self *Listener) Exec(evt *Event) {
	self.Func(evt, self.Data)
}
