package refund

import (
	"github.com/bububa/jinritemai-go/client"
)

type ShopRefundRequest struct {
	OrderID       string `json:"order_id,omitempty"`       // 父订单id，由orderList接口返回
	Agree         bool   `json:"agree,omitempty"`          // 处理方式：1：同意退款; 2：不同意退款
	LogisticsID   uint64 `json:"logistics_id,omitempty"`   // type = 2 时必须填写物流公司ID，由接口/order/logisticsCompanyList返回的物流公司列表中对应的ID
	LogisticsCode string `json:"logistics_code,omitempty"` // type = 2 时必须填写物流单号
}

func (this ShopRefundRequest) Method() string {
	return "refund/shopRefund"
}

func (this ShopRefundRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id":       this.OrderID,
		"logistics_code": this.LogisticsCode,
	}
	if this.Agree {
		ret["type"] = 1
	} else {
		ret["type"] = 2
	}
	if this.LogisticsID > 0 {
		ret["logistics_id"] = this.LogisticsID
	}
	if this.LogisticsCode != "" {
		ret["logistics_code"] = this.LogisticsCode
	}
	return ret
}

// 商家处理备货中退款申请
// 订单备货中，用户申请退款，商家处理。
func ShopRefund(clt *client.Client, req *ShopRefundRequest) error {
	return clt.Do(req, nil)
}
