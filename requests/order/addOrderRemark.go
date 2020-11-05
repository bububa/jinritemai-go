package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddOrderRemarkRequest struct {
	OrderID   string `json:"order_id,omitempty"`    // 父订单id，由orderList接口返回
	Remark    string `json:"remark,omitempty"`      // 备注内容，最大不得超过60个字符
	IsAddStar bool   `json:"is_add_star,omitempty"` // 是否加旗标，不填则默认为否; true：需要加旗标; false：不加旗标
	Star      uint   `json:"star,omitempty"`        // 标星等级，范围0～5 0为灰色旗标，5为红色旗标，数字越大颜色越深
}

func (this AddOrderRemarkRequest) Method() string {
	return "order/addOrderRemark"
}

func (this AddOrderRemarkRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id":    this.OrderID,
		"remark":      this.Remark,
		"is_add_star": this.IsAddStar,
	}
	if this.IsAddStar && this.Star <= 5 {
		ret["star"] = this.Star
	}
	return ret
}

// 添加订单备注
// 添加订单备注，给订单加旗标
func AddOrderRemark(clt *client.Client, req *AddOrderRemarkRequest) error {
	return clt.Do(req, nil)
}
