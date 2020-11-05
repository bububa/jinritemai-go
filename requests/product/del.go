package product

import (
	"github.com/bububa/jinritemai-go/client"
)

type DelRequest struct {
	ProductID uint64 `json:"product_id,omitempty"` // 商品ID
}

func (this DelRequest) Method() string {
	return "product/del"
}

func (this DelRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"product_id": this.ProductID,
	}
}

// 删除商品
func Del(clt *client.Client, productID uint64) error {
	req := &DelRequest{
		ProductID: productID,
	}
	return clt.Do(req, nil)
}
