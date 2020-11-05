package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type EditPriceRequest struct {
	ProductID uint64 `json:"product_id,omitempty"` // 商品id
	SkuID     uint64 `json:"sku_id,omitempty"`     // SKU ID
	Price     int    `json:"price,omitempty"`      // 售价 (单位 分)
}

func (this EditPriceRequest) Method() string {
	return "sku/editPrice"
}

func (this EditPriceRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"product_id": this.ProductID,
		"sku_id":     this.SkuID,
		"price":      this.Price,
	}
}

// 编辑sku价格
func EditPrice(clt *client.Client, productID uint64, skuID uint64, price int) error {
	req := &EditPriceRequest{
		ProductID: productID,
		SkuID:     skuID,
		Price:     price,
	}
	return clt.Do(req, nil)
}
