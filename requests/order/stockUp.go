package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type StockUpRequest struct {
	OrderID string `json:"order_id,omitempty"` // 父订单id，由orderList接口返回
}

func (this StockUpRequest) Method() string {
	return "order/stockUp"
}

func (this StockUpRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
	}
}

// 确认货到付款订单
// 当货到付款订单是待确认状态（final_status=1）时，可调此接口确认订单。确认后，订单状态更新为“备货中”
func StockUp(clt *client.Client, orderID string) error {
	req := &StockUpRequest{
		OrderID: orderID,
	}
	return clt.Do(req, nil)
}
