package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type InsuranceRequest struct {
	OrderID string `json:"order_id,omitempty"` // 父订单，订单号后面需要加大写字母"A"
}

func (this InsuranceRequest) Method() string {
	return "order/insurance"
}

func (this InsuranceRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
	}
}

type InsuranceResponse struct {
	List []InsPolicy `json:"policy_info_list,omitempty"`
}

// 获取运费险保单详情
// 根据订单ID，获取该订单对应的运费险保单的详细信息
func GetInsurance(clt *client.Client, orderID string) ([]InsPolicy, error) {
	req := &InsuranceRequest{
		OrderID: orderID,
	}
	var ret InsuranceResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
