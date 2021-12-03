package event

type Event struct {
    Name string
    Data interface{}
}

func NewEvent(name string, data interface{}) *Event {
    evt := &Event{}
    evt.Set(name, data)
    return evt
}

func (evt *Event) Set(name string, data interface{}) {
    evt.Name = name
    evt.Data = data
}
