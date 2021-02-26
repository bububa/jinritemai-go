package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type StockNumRequest struct {
	SkuID          uint64 `json:"sku_id,omitempty"`
	OutWarehouseID uint64 `json:"out_warehouse_id,omitempty"`
}

func (this StockNumRequest) Method() string {
	return "sku/stockNum"
}

func (this StockNumRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"sku_id": this.SkuID,
	}
	if this.OutWarehouseID > 0 {
		ret["out_warehouse_id"] = this.OutWarehouseID
	}
	return ret
}

// 查询库存
func GetStockNum(clt *client.Client, skuID uint64, outWarehouseID uint64) (*SkuStock, error) {
	req := &StockNumRequest{
		SkuID:          skuID,
		OutWarehouseID: outWarehouseID,
	}
	var ret SkuStock
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
