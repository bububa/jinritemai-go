package product

import (
	"github.com/bububa/jinritemai-go/client"
	"github.com/bububa/jinritemai-go/requests/category"
)

type GetCatePropertyRequest struct {
	FirstCid  uint64 `json:"first_cid,omitempty"`
	SecondCid uint64 `json:"second_cid,omitempty"`
	ThirdCid  uint64 `json:"third_cid,omitempty"`
}

func (this GetCatePropertyRequest) Method() string {
	return "product/getCateProperty"
}

func (this GetCatePropertyRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"first_cid":  this.FirstCid,
		"second_cid": this.SecondCid,
		"third_cid":  this.ThirdCid,
	}
}

// 根据商品分类获取对应的属性列表
// 调用API接口创建商品时，入参product_format（属性对）依赖此接口返回的值
func GetCateProperty(clt *client.Client, firstCid uint64, secondCid uint64, thirdCid uint64) ([]category.CategoryProperty, error) {
	req := &GetCatePropertyRequest{
		FirstCid:  firstCid,
		SecondCid: secondCid,
		ThirdCid:  thirdCid,
	}
	var ret []category.CategoryProperty
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
