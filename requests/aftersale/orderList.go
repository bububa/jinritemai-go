package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
	"github.com/bububa/jinritemai-go/requests/order"
)

/*
type值对应表
2   6: 退货申请待商家处理
30: 仅退款申请待商家处理
3   11: 买家已退货，待商家收货
4   8: 商家拒绝退货申请，待客服仲裁
13: 商家拒绝收货，待客服仲裁
34: 商家拒绝仅退款申请，待客服仲裁
5   22: 在线支付订单，退货退款成功
24: 货到付款订单，退货退款成功
39: 仅退款成功
7   7 & 10: 待买家填写退货物流
*/
type OrderListRequest struct {
	Page      int    `json:"page,omitempty"`       // 第几页（第一页为0，最大为99）
	Size      int    `json:"size,omitempty"`       // 每页返回条数，最多支持100条
	Type      uint   `json:"type,omitempty"`       // 类型(1.全部售后单 2.待商家处理 3.待商家收货 4.待客服仲裁 5.退款成功 7.待买家退货)默认为1.全部该接口无法获取到待买家处理的订单
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间，必须大于等于开始时间
	OrderBy   string `json:"order_by,omitempty"`   // 1、默认按订单创建时间搜索 2、值为“create_time”：按订单创建时间；值为“update_time”：按订单更新时间
	IsDesc    *bool  `json:"is_desc,omitempty"`    // 订单排序方式：0(is_desc，最近的在前)， 1(asc，最近的在后) 默认为1
}

func (this OrderListRequest) Method() string {
	return "afterSale/orderList"
}

func (this OrderListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"start_time": this.StartTime,
		"end_time":   this.EndTime,
		"order_by":   this.OrderBy,
		"page":       this.Page,
	}
	if this.Type > 0 {
		ret["type"] = this.Type
	}
	if this.IsDesc != nil && *this.IsDesc {
		ret["is_desc"] = 1
	}
	if this.Size > 0 {
		ret["size"] = this.Size
	}
	return ret
}

// 获取已发货且有售后的订单列表
/*
订单已发货，通过该接口可拉取有售后的订单：

1. 售后仅退款
2. 售后退货
*/
func GetOrderList(clt *client.Client, req *OrderListRequest) (*order.ListResponse, error) {
	var ret order.ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
