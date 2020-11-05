package order

import (
	"errors"

	"github.com/bububa/jinritemai-go/client"
)

type DetailRequest struct {
	OrderID string `json:"order_id,omitempty"` // 父订单，订单号后面需要加大写字母"A"
}

func (this DetailRequest) Method() string {
	return "order/detail"
}

func (this DetailRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
	}
}

// 获取订单详情
// 根据订单ID，获取单个订单的详情信息
func GetDetail(clt *client.Client, orderID string) (*OrderDetail, error) {
	req := &DetailRequest{
		OrderID: orderID,
	}
	var ret ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	if len(ret.List) == 0 {
		return nil, errors.New("not found")
	}
	return &ret.List[0], nil
}
