package spec

import (
	"github.com/bububa/jinritemai-go/client"
)

type ListRequest struct {
}

func (this ListRequest) Method() string {
	return "spec/list"
}

func (this ListRequest) Params() map[string]interface{} {
	return nil
}

func GetList(clt *client.Client) ([]SpecDetail, error) {
	req := &ListRequest{}
	var ret []SpecDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
