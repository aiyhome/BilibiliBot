package event

type Event struct {
	Name   string
	Data   interface{}
	Stoped bool
}

func NewEvent(name string, data interface{}) *Event {
	evt := &Event{}
	evt.Set(name, data)
	return evt
}

func (self *Event) Set(name string, data interface{}) {
	self.Name = name
	self.Data = data
}

func (self *Event) stopPropagation() {
	self.Stoped = true
}
