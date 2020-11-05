package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type LogisticsEditRequest struct {
	OrderID       string `json:"order_id,omitempty"`       // 父订单id，由orderList接口返回
	LogisticsID   uint64 `json:"logistics_id,omitempty"`   // 物流公司ID，由接口/order/logisticsCompanyList返回的物流公司列表中对应的ID
	LogisticsCode string `json:"logistics_code,omitempty"` // 运单号
	Company       string `json:"company,omitempty"`        // 物流公司名称
}

func (this LogisticsEditRequest) Method() string {
	return "order/logisticsEdit"
}

func (this LogisticsEditRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id":       this.OrderID,
		"logistics_id":   this.LogisticsID,
		"logistics_code": this.LogisticsCode,
	}
	if this.Company != "" {
		ret["company"] = this.Company
	}
	return ret
}

// 修改发货物流
// 修改已发货订单（final_status=3）的发货物流信息
func LogisticsEdit(clt *client.Client, req *LogisticsEditRequest) error {
	return clt.Do(req, nil)
}
