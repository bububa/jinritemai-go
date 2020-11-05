package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type EditCodeRequest struct {
	ProductID uint64 `json:"product_id,omitempty"` // 商品id
	SkuID     uint64 `json:"sku_id,omitempty"`     // SKU ID
	Code      string `json:"code,omitempty"`       // 售价 (单位 分)
}

func (this EditCodeRequest) Method() string {
	return "sku/editCode"
}

func (this EditCodeRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"product_id": this.ProductID,
		"sku_id":     this.SkuID,
		"code":       this.Code,
	}
}

// 修改sku编码
func EditCode(clt *client.Client, productID uint64, skuID uint64, code string) error {
	req := &EditCodeRequest{
		ProductID: productID,
		SkuID:     skuID,
		Code:      code,
	}
	return clt.Do(req, nil)
}
