package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

type RemoveAddrRequest struct {
	OutWarehouseID uint64         `json:"out_warehouse_id,omitempty"` // 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	Addr           *WarehouseAddr `json:"addr,omitempty"`             // 仓库地址定义
}

func (this RemoveAddrRequest) Method() string {
	return "warehouse/removeAddr"
}

func (this RemoveAddrRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"out_warehouse_id": this.OutWarehouseID,
		"addr":             this.Addr,
	}
}

// 地址与区域仓解绑
func RemoveAddr(clt *client.Client, outWarehouseID uint64, addr *WarehouseAddr) error {
	req := &RemoveAddrRequest{
		OutWarehouseID: outWarehouseID,
		Addr:           addr,
	}
	return clt.Do(req, nil)
}
