package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddRequest struct {
	ProductID       uint64 `json:"product_id,omitempty"`       // 商品id
	OutID           uint64 `json:"out_sku_id,omitempty"`       // 业务方自己的sku_id，唯一需为数字字符串，max = int64
	SpecID          uint64 `json:"spec_id,omitempty"`          // 规格id，依赖/spec/list接口的返回
	SpecDetailIDs   string `json:"spec_detail_ids,omitempty"`  // 子规格id,最多3级,如 100041|150041|160041 （ 女款|白色|XL）
	StockNum        int    `json:"stock_num,omitempty"`        // 库存 (必须大于0)
	Price           int    `json:"price,omitempty"`            // 售价 (单位 分)
	SettlementPrice int    `json:"settlement_price,omitempty"` // 结算价格 (单位 分)
	Code            string `json:"code,omitempty"`             // 商品编码
	OutWarehouseID  uint64 `json:"out_warehouse_id,omitempty"` // 外部仓库ID，需要注意，如果设置了此参数，该sku类型会变成区域库存类，原来的全国库存数据会被覆盖
	SupplierID      uint64 `json:"supplier_id,omitempty"`      // 供应商ID
}

func (this AddRequest) Method() string {
	return "sku/add"
}

func (this AddRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"product_id":      this.ProductID,
		"out_sku_id":      this.OutID,
		"spec_id":         this.SpecID,
		"spec_detail_ids": this.SpecDetailIDs,
		"stock_num":       this.StockNum,
		"price":           this.Price,
	}
	if this.SettlementPrice > 0 {
		ret["settlement_price"] = this.SettlementPrice
	}
	if this.Code != "" {
		ret["code"] = this.Code
	}
	if this.OutWarehouseID > 0 {
		ret["out_warehouse_id"] = this.OutWarehouseID
	}
	if this.SupplierID > 0 {
		ret["supplier_id"] = this.SupplierID
	}
	return ret
}

// 添加SKU
// 单个规格可设置的规格值最多是20个
func Add(clt *client.Client, req *AddRequest) (uint64, error) {
	var ret uint64
	err := clt.Do(req, &ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}
