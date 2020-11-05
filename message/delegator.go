package message

import (
	"sync"
)

type handlerFunc func(Message) error

type Delegator struct {
	handlers map[string]handlerFunc
	sync.RWMutex
}

func NewDelegator() *Delegator {
	return &Delegator{
		handlers: make(map[string]handlerFunc, 15),
	}
}

func (this *Delegator) Register(tag string, fn handlerFunc) {
	this.Lock()
	defer this.Unlock()
	this.handlers[tag] = fn
}

func (this *Delegator) Unregister(tag string) {
	this.Lock()
	defer this.Unlock()
	delete(this.handlers, tag)
}

func (this *Delegator) Get(tag string) handlerFunc {
	this.RLock()
	defer this.RUnlock()
	if fn, found := this.handlers[tag]; found {
		return fn
	}
	return nil
}
