package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

type RefundProcessDetailRequest struct {
	OrderID string `json:"order_id,omitempty"` // 子订单ID，不带字母A
}

func (this RefundProcessDetailRequest) Method() string {
	return "afterSale/refundProcessDetail"
}

func (this RefundProcessDetailRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
	}
}

// 根据子订单ID查询退款详情
/*
通过该接口，根据子订单ID查询退款详情信息
1、订单未发货，买家申请整单退款
2、订单已发货，买家申请发货后仅退款
3、订单已发货，买家申请发货后退货
*/
func GetRefundProcessDetail(clt *client.Client, orderID string) (*RefundProcessDetail, error) {
	req := &RefundProcessDetailRequest{
		OrderID: orderID,
	}
	var ret RefundProcessDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
