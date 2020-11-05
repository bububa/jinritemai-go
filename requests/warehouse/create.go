package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

type CreateRequest struct {
	OutWarehouseID uint64 `json:"out_warehouse_id,omitempty"` // 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	Name           string `json:"name,omitempty"`             // 仓库名称
	Info           string `json:"info,omitempty"`             // 仓库介绍
}

func (this CreateRequest) Method() string {
	return "warehouse/create"
}

func (this CreateRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"out_warehouse_id": this.OutWarehouseID,
		"name":             this.Name,
		"info":             this.Info,
	}
}

// 创建区域仓
func Create(clt *client.Client, outWarehouseID uint64, name string, info string) (uint64, error) {
	req := &CreateRequest{
		OutWarehouseID: outWarehouseID,
		Name:           name,
		Info:           info,
	}
	var ret uint64
	err := clt.Do(req, &ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}
