package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

type InfoRequest struct {
	OutWarehouseID uint64 `json:"out_warehouse_id,omitempty"` // 外部仓库ID
}

func (this InfoRequest) Method() string {
	return "warehouse/info"
}

func (this InfoRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"out_warehouse_id": this.OutWarehouseID,
	}
}

// 查询区域仓
func Info(clt *client.Client, outWarehouseID uint64) (*WarehouseInfo, error) {
	req := &InfoRequest{
		OutWarehouseID: outWarehouseID,
	}
	var ret WarehouseInfo
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
