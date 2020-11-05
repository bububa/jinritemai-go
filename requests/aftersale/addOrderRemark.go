package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddOrderRemarkRequest struct {
	OrderID string `json:"order_id,omitempty"` // 子订单ID
	Remark  string `json:"remark,omitempty"`   // 售后备注内容
}

func (this AddOrderRemarkRequest) Method() string {
	return "afterSale/addOrderRemark"
}

func (this AddOrderRemarkRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
		"remark":   this.Remark,
	}
}

// 商家为订单添加售后备注
// 此接口可用于给备货中退款的订单、已发货退货/仅退款的订单，添加售后备注
func AddOrderRemark(clt *client.Client, orderID string, remark string) error {
	req := &AddOrderRemarkRequest{
		OrderID: orderID,
		Remark:  remark,
	}
	return clt.Do(req, nil)
}
