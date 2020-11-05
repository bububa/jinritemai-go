package client

type Request interface {
	Method() string
	Params() map[string]interface{}
}
