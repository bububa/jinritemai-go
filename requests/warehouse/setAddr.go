package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

type SetAddrRequest struct {
	OutWarehouseID uint64         `json:"out_warehouse_id,omitempty"` // 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	Addr           *WarehouseAddr `json:"addr,omitempty"`             // 仓库地址定义
}

func (this SetAddrRequest) Method() string {
	return "warehouse/setAddr"
}

func (this SetAddrRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"out_warehouse_id": this.OutWarehouseID,
		"addr":             this.Addr,
	}
}

// 地址与区域仓绑定
func SetAddr(clt *client.Client, outWarehouseID uint64, addr *WarehouseAddr) error {
	req := &SetAddrRequest{
		OutWarehouseID: outWarehouseID,
		Addr:           addr,
	}
	return clt.Do(req, nil)
}
