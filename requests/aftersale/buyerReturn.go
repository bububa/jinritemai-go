package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

/*
comment值对应表
1   未收到货（未填写退货单号）
2   退货与原订单不符（商品不符、退货地址不符）
3   退回商品影响二次销售', '订单超出售后时效（订单完成超7天）
4   买家误操作/取消申请
5   已与买家协商补偿，包括差价、赠品、额外补偿
6   已与买家协商补发商品', '已与买家协商换货
*/
type BuyerReturnRequest struct {
	OrderID          string `json:"order_id,omitempty"`          // 子订单ID
	Agree            bool   `json:"agree,omitempty"`             // 处理方式：1：同意退款; 2：不同意退款
	Sms              bool   `json:"sms,omitempty"`               // 是否使用模版短信发送短信: 1：是; 0：否
	Comment          uint   `json:"comment,omitempty"`           // type = 2 时需要选择拒绝原因
	Evidence         string `json:"evidence,omitempty"`          // type = 2 时需要上传图片凭证，仅支持1张图片
	AddressID        uint64 `json:"address_id,omitempty"`        // 商家退货物流收货地址id,不传则使用默认售后收货地址, type=1时，可以不入参address_id，直接入参退货地址的详细文本信息
	ReceiverName     string `json:"receiver_name,omitempty"`     // 退回商品的收货人姓名
	ReceiverTel      string `json:"receiver_tel,omitempty"`      // 退回商品的收货手机号
	ReceiverProvince string `json:"receiver_province,omitempty"` // 退货地址的省（直辖市也必须填，比如北京市）
	ReceiverCity     string `json:"receiver_city,omitempty"`     // 退货地址的市
	ReceiverDistrict string `json:"receiver_district,omitempty"` // 退货地址的区
	ReceiverAddress  string `json:"receiver_address,omitempty"`  // 退货详细地址
}

func (this BuyerReturnRequest) Method() string {
	return "afterSale/buyerReturn"
}

func (this BuyerReturnRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"order_id": this.OrderID,
	}
	if this.Agree {
		ret["type"] = 1
	} else {
		ret["type"] = 2
	}
	if this.Sms {
		ret["sms"] = 1
	} else {
		ret["sms"] = 0
	}
	if this.Comment > 0 && this.Comment <= 6 {
		ret["comment"] = this.Comment
	}
	if this.Evidence != "" {
		ret["evidence"] = this.Evidence
	}
	if this.AddressID > 0 {
		ret["address_id"] = this.AddressID
	}
	if this.ReceiverName != "" {
		ret["receiver_name"] = this.ReceiverName
	}
	if this.ReceiverTel != "" {
		ret["receiver_tel"] = this.ReceiverTel
	}
	if this.ReceiverProvince != "" {
		ret["receiver_province"] = this.ReceiverProvince
	}
	if this.ReceiverCity != "" {
		ret["receiver_city"] = this.ReceiverCity
	}
	if this.ReceiverDistrict != "" {
		ret["receiver_district"] = this.ReceiverDistrict
	}
	if this.ReceiverAddress != "" {
		ret["receiver_address"] = this.ReceiverAddress
	}
	return ret
}

// 商家处理退货申请
/*
订单已发货，用户申请售后退货，商家处理。商家拒绝退货申请时，address_id是选填入参。

1、address_id不为空时，取address_id对应的地址库地址，作为退货地址
2、address_id为空时：
 * 如果退货地址文本信息也为空，则取店铺地址库默认退货地址**
 * 如果退货地址文本信息不为空，则取详细地址文本**
*/
func BuyerReturn(clt *client.Client, req *BuyerReturnRequest) error {
	return clt.Do(req, nil)
}
