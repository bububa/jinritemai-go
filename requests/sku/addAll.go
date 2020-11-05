package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddAllRequest struct {
	ProductID       string `json:"product_id,omitempty"`       // 商品id
	OutID           string `json:"out_sku_id,omitempty"`       // 业务方自己的sku_id，唯一需为数字字符串，max = int64
	SpecID          string `json:"spec_id,omitempty"`          // 规格id，依赖/spec/list接口的返回; 数量与spec_detail_ids的一致值必须要能与spec_detail_ids的每一项对应得起来
	SpecDetailIDs   string `json:"spec_detail_ids,omitempty"`  // 子规格id，最多3级,如 100041|150041|160041 （ 女款|白色|XL）; 子规格ID的顺序必须严格按照spec_id详情里的顺序，比如颜色在前、尺寸在后
	StockNum        string `json:"stock_num,omitempty"`        // 库存 (必须大于0)
	Price           string `json:"price,omitempty"`            // 售价 (单位 分)
	SettlementPrice string `json:"settlement_price,omitempty"` // 结算价格 (单位 分)
	Code            string `json:"code,omitempty"`             // 商品编码
	OutWarehouseID  string `json:"out_warehouse_id,omitempty"` // 数量需要与sku数量一致，用
	SupplierID      string `json:"supplier_id,omitempty"`      // 数量需要与sku数量一致，用
}

func (this AddAllRequest) Method() string {
	return "sku/addAll"
}

func (this AddAllRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"product_id":      this.ProductID,
		"out_sku_id":      this.OutID,
		"spec_id":         this.SpecID,
		"spec_detail_ids": this.SpecDetailIDs,
		"stock_num":       this.StockNum,
		"price":           this.Price,
	}
	if this.SettlementPrice != "" {
		ret["settlement_price"] = this.SettlementPrice
	}
	if this.Code != "" {
		ret["code"] = this.Code
	}
	if this.OutWarehouseID != "" {
		ret["out_warehouse_id"] = this.OutWarehouseID
	}
	if this.SupplierID != "" {
		ret["supplier_id"] = this.SupplierID
	}
	return ret
}

// 批量添加sku
// 批量添加商品sku（每次接口调用最多添加100个）
func AddAll(clt *client.Client, req *AddAllRequest) (map[string]uint64, error) {
	ret := make(map[string]uint64)
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
