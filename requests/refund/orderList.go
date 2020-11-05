package refund

import (
	"github.com/bububa/jinritemai-go/client"
	"github.com/bububa/jinritemai-go/requests/order"
)

type OrderListRequest struct {
	Page      int    `json:"page,omitempty"`       // 第几页（第一页为0，最大为99）
	Size      int    `json:"size,omitempty"`       // 每页返回条数，最多支持100条
	Type      uint   `json:"type,omitempty"`       // 类型(1.全部 2.待商家处理 5.退款成功) 默认为1
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间，必须大于等于开始时间
	OrderBy   string `json:"order_by,omitempty"`   // 1、默认按订单创建时间搜索 2、值为“create_time”：按订单创建时间；值为“update_time”：按订单更新时间
	IsDesc    *bool  `json:"is_desc,omitempty"`    // 订单排序方式：0(is_desc，最近的在前)， 1(asc，最近的在后) 默认为1
}

func (this OrderListRequest) Method() string {
	return "refund/orderList"
}

func (this OrderListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"start_time": this.StartTime,
		"end_time":   this.EndTime,
		"order_by":   this.OrderBy,
		"page":       this.Page,
	}
	if this.Type == 1 || this.Type == 2 || this.Type == 5 {
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

func GetOrderList(clt *client.Client, req *OrderListRequest) (*order.ListResponse, error) {
	var ret order.ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
