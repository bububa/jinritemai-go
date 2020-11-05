package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type LogisticsAddRequest struct {
	OrderID       string `json:"order_id,omitempty"`       // 父订单id，由orderList接口返回
	LogisticsID   uint64 `json:"logistics_id,omitempty"`   // 物流公司ID，由接口/order/logisticsCompanyList返回的物流公司列表中对应的ID
	LogisticsCode string `json:"logistics_code,omitempty"` // 运单号
	Company       string `json:"company,omitempty"`        // 物流公司名称
}

func (this LogisticsAddRequest) Method() string {
	return "order/logisticsAdd"
}

func (this LogisticsAddRequest) Params() map[string]interface{} {
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

// 订单发货
// 暂时只支持整单出库，即接口调用时入参只能传父订单号
func LogisticsAdd(clt *client.Client, req *LogisticsAddRequest) error {
	return clt.Do(req, nil)
}
