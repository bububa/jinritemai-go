package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

/*
comment值对应表
1   商品退回后才能退款   是   需提供发货物流进度截图
2   买家已签收   是   需提供发货物流进度截图
3   买家误操作/取消申请  否   请保留好短信确认凭证，如后续买家投诉无法提供短信凭证，将支持买家并处罚商家
4   问题已解决，待用户收货 否   请保留好短信确认凭证，如后续买家投诉无法提供短信凭证，将支持买家并处罚商家
5   协商一致，用户取消退款 否   请保留好短信确认凭证，如后续买家投诉无法提供短信凭证，将支持买家并处罚商家
6   商品影响二次销售    是   请提供影响二次销售凭证
7   三方卡券订单，卡券状态变更，退款失败，请重新申请    否   无
*/
type BuyerRefundRequest struct {
	OrderID  string `json:"order_id,omitempty"` // 子订单ID
	Agree    bool   `json:"agree,omitempty"`    // 处理方式：1：同意退款; 2：拒绝用户售后仅退款申请
	Comment  uint   `json:"comment,omitempty"`  // type = 2 时需要选择拒绝原因
	Evidence string `json:"evidence,omitempty"` // type = 2 时需要上传图片凭证，仅支持1张图片
}

func (this BuyerRefundRequest) Method() string {
	return "afterSale/buyerRefund"
}

func (this BuyerRefundRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id": this.OrderID,
		"evidence": this.Evidence,
	}
	if this.Agree {
		ret["type"] = 1
	} else {
		ret["type"] = 2
	}
	if this.Comment > 0 && this.Comment <= 6 {
		ret["comment"] = this.Comment
	}
	return ret
}

// 商家确认是否收到退货
// 用户填写退货物流后，商家处理，确认是否收到退货
// 注：货到付款订单，调此接口确认收货，必须上传退款凭证图片。售后状态会变为24（退货成功-商户已退款）
func BuyerRefund(clt *client.Client, req *BuyerRefundRequest) error {
	return clt.Do(req, nil)
}
