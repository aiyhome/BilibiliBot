package event

type handleFunc func(evt: Event, data ...interface{})
type extData interface{}

type Listener struct {
    Event string
    Func  handleFunc
    Data  extData
}

func NewListener(name string, callback handleFunc, data ...interface{}) *Listener {
    l := &Listener{}
    if len(data) > 0 {
        l.Set(name, handleFunc, data[0])
    } else {
        l.Set(name, handleFunc, nil)
    }
    return l
}

func (l *Listener) Set(name string, callback handleFunc, data interface{}) {
    l.Name = name
    l.Func = callback
    l.Data = data
}

func (l *Listener) Exec(data interface{}) {
    evt := NewEvent(name,data)
    l.Func(evt, l.Data)
}
