package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

type EditRequest struct {
	OutWarehouseID uint64 `json:"out_warehouse_id,omitempty"` // 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	Name           string `json:"name,omitempty"`             // 仓库名称
	Info           string `json:"info,omitempty"`             // 仓库介绍
}

func (this EditRequest) Method() string {
	return "warehouse/edit"
}

func (this EditRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"out_warehouse_id": this.OutWarehouseID,
		"name":             this.Name,
		"info":             this.Info,
	}
}

// 编辑区域仓
func Edit(clt *client.Client, outWarehouseID uint64, name string, info string) error {
	req := &EditRequest{
		OutWarehouseID: outWarehouseID,
		Name:           name,
		Info:           info,
	}
	return clt.Do(req, nil)
}
