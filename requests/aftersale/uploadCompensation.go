package aftersale

import (
	"github.com/bububa/jinritemai-go/client"
)

type UploadCompensationRequest struct {
	OrderID  string `json:"order_id,omitempty"` // 子订单ID
	Comment  string `json:"comment,omitempty"`  // type = 2 时需要选择拒绝原因
	Evidence string `json:"evidence,omitempty"` // type = 2 时需要上传图片凭证，仅支持1张图片
}

func (this UploadCompensationRequest) Method() string {
	return "afterSale/uploadCompensation"
}

func (this UploadCompensationRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"order_id": this.OrderID,
		"comment":  this.Comment,
		"evidence": this.Evidence,
	}
}

// 货到付款订单上传退款凭证
// 货到付款订单，用户申请退货，商家确认收到退货时（final_status=12 or 14时）。如果需要上传退款凭证，需要调此接口！
func UploadCompensation(clt *client.Client, orderID string, comment string, evidence string) error {
	req := &UploadCompensationRequest{
		OrderID:  orderID,
		Comment:  comment,
		Evidence: evidence,
	}
	return clt.Do(req, nil)
}
