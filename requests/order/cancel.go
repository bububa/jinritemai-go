package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type CancelRequest struct {
	OrderID      string `json:"order_id,omitempty"`      // 父订单id，由orderList接口返回
	CancelReason string `json:"cancel_reason,omitempty"` // 取消订单的原因
}

func (this CancelRequest) Method() string {
	return "order/cancel"
}

func (this CancelRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id":      this.OrderID,
		"cancel_reason": this.CancelReason,
	}
}

// 取消订单
// 取消待确认或备货中的货到付款订单（final_status=1或2）
func Cancel(clt *client.Client, orderID string, cancelReason string) error {
	req := &CancelRequest{
		OrderID:      orderID,
		CancelReason: cancelReason,
	}
	return clt.Do(req, nil)
}
