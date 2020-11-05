package product

import (
	"github.com/bububa/jinritemai-go/client"
	"github.com/bububa/jinritemai-go/requests/category"
)

type GetGoodsCategoryRequest struct {
	Cid uint64 `json:"cid,omitempty"`
}

func (this GetGoodsCategoryRequest) Method() string {
	return "product/getGoodsCategory"
}

func (this GetGoodsCategoryRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"cid": this.Cid,
	}
}

// 获取商品分类列表
// 根据父分类id获取子分类
func GetGoodsCategory(clt *client.Client, cid uint64) ([]category.GoodsCategory, error) {
	req := &GetGoodsCategoryRequest{
		Cid: cid,
	}
	var ret []category.GoodsCategory
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
