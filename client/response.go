package client

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	ErrNo   int             `json:"err_no,omitempty"`
	Message string          `json:"message,omitempty"`
	data    json.RawMessage `json:"data,omitempty"`
}

func (this *Response) IsError() bool {
	return this.ErrNo != 0
}

func (this Response) Error() string {
	return fmt.Sprintf("CODE:%s, MSG:%s", this.ErrNo, this.Message)
}

func (this *Response) Data(v interface{}) error {
	if v == nil {
		return nil
	}
	return json.Unmarshal(this.data, v)
}
