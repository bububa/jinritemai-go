package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type SyncStockRequest struct {
	ProductID      uint64 `json:"product_id,omitempty"`       // 商品id
	SkuID          uint64 `json:"sku_id,omitempty"`           // SKU ID
	OutWarehouseID uint64 `json:"out_warehouse_id,omitempty"` // 外部仓库ID
	SupplierID     uint64 `json:"supplier_id,omitempty"`      // 供应商ID
	IdempotentID   int    `json:"idempotent_id,omitempty"`    // 幂等ID，仅incremental=true时有用
	Incremental    bool   `json:"incremental,omitempty"`      // true表示增量库存，false表示全量库存，默认为false
	StockNum       *int   `json:"stock_num,omitempty"`        // 库存 (可以为0)
}

func (this SyncStockRequest) Method() string {
	return "sku/syncStock"
}

func (this SyncStockRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"product_id": this.ProductID,
		"sku_id":     this.SkuID,
	}
	if this.OutWarehouseID > 0 {
		ret["out_warehouse_id"] = this.OutWarehouseID
	}
	if this.Incremental {
		ret["incremental"] = this.Incremental
		if this.IdempotentID > 0 {
			ret["idempotent_id"] = this.IdempotentID
		}
	}
	if this.SupplierID > 0 {
		ret["supplier_id"] = this.SupplierID
	}
	if this.StockNum != nil {
		ret["stock_num"] = *this.StockNum
	}
	return ret
}

// 修改sku库存
// 注：同步库存时请注意sku对应商品的状态status和check_status, 下架、删除、禁封等状态的商品不予同步sku库存
func SyncStock(clt *client.Client, req *SyncStockRequest) error {
	return clt.Do(req, nil)
}
