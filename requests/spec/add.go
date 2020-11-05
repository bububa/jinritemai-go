package spec

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddRequest struct {
	Specs string `json:"specs,omitempty"` // 店铺通用规格，能被同类商品通用。多规格用^分隔，父规格与子规格用|分隔，子规格用,分隔
	Name  string `json:"name,omitempty"`  // 如果不填，则规格名为子规格名用 "-" 自动生成
}

func (this AddRequest) Method() string {
	return "spec/add"
}

func (this AddRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"specs": this.Specs,
		"name":  this.Name,
	}
}

// 添加规格
// 1. 一个规格组下，组合总数不能超过600
// 2. 单个规格值数量不能超过20个，比如“颜色”不能超过20种
func Add(clt *client.Client, req *AddRequest) (*Spec, error) {
	var ret Spec
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
