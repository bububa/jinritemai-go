package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

/*
comment值对应表
1   未收到货（未填写退货单号）
2   退货与原订单不符（商品不符、退货地址不符）
3   退回商品影响二次销售/订单超出售后时效（订单完成超7天）
4   买家误操作/取消申请
5   已与买家协商补偿，包括差价、赠品、额外补偿
6   已与买家协商补发商品/已与买家协商换货
*/
type FirmReceiveRequest struct {
	OrderID  string `json:"order_id,omitempty"` // 子订单ID
	Agree    bool   `json:"agree,omitempty"`    // 处理方式：1：确认收货并退款; 2：拒绝
	Comment  uint   `json:"comment,omitempty"`  // type = 2 时需要选择拒绝原因
	Evidence string `json:"evidence,omitempty"` // type = 2 时需要上传图片凭证，仅支持1张图片
}

func (this FirmReceiveRequest) Method() string {
	return "afterSale/firmReceive"
}

func (this FirmReceiveRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id": this.OrderID,
		"register": 1,
		"package":  1,
		"facade":   1,
		"function": 1,
	}
	if this.Agree {
		ret["type"] = 1
	} else {
		ret["type"] = 2
	}
	if this.Comment > 0 && this.Comment <= 6 {
		ret["comment"] = this.Comment
	}
	if this.Evidence != "" {
		ret["evidence"] = this.Evidence
	}
	return ret
}

// 商家确认是否收到退货
// 用户填写退货物流后，商家处理，确认是否收到退货
// 注：货到付款订单，调此接口确认收货，必须上传退款凭证图片。售后状态会变为24（退货成功-商户已退款）
func FirmReceive(clt *client.Client, req *FirmReceiveRequest) error {
	return clt.Do(req, nil)
}
