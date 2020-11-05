package spec

import (
	"github.com/bububa/jinritemai-go/client"
)

type DelRequest struct {
	ID uint64 `json:"id,omitempty"` // 规格id (spec_id)
}

func (this DelRequest) Method() string {
	return "spec/del"
}

func (this DelRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"id": this.ID,
	}
}

func Del(clt *client.Client, specID uint64) error {
	req := &DelRequest{
		ID: specID,
	}
	return clt.Do(req, nil)
}
