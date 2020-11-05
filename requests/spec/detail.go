package spec

import (
	"github.com/bububa/jinritemai-go/client"
)

type DetailRequest struct {
	ID uint64 `json:"id,omitempty"` // 规格id (spec_id)
}

func (this DetailRequest) Method() string {
	return "spec/specDetail"
}

func (this DetailRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"id": this.ID,
	}
}

func GetDetail(clt *client.Client, specID uint64) (*SpecDetail, error) {
	req := &DetailRequest{
		ID: specID,
	}
	var ret SpecDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
