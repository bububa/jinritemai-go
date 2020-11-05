package refund

// 拒绝退款消息
/*
触发场景
买家在发货前申请整单退款，卖家拒绝
买家在发货后申请仅退款，卖家拒绝
买家在发货后申请申请退货，买家发货后，卖家拒绝收货
*/
type RefundRefused struct {
	Sid              uint64 `json:"s_id,omitempty"`               // 子订单ID
	Pid              uint64 `json:"p_id,omitempty"`               // 父订单ID
	ShopID           uint64 `json:"shop_id,omitempty"`            // 店铺ID
	AftersaleID      uint64 `json:"aftersale_id,omitempty"`       // 售后单ID
	AftersaleStatus  uint   `json:"aftersale_status,omitempty"`   // 售后状态
	AftersaleType    uint   `json:"aftersale_type,omitempty"`     // 售后类型：0: 退货; 1: 售后仅退款; 2: 发货前整单退款
	RefundAmount     int    `json:"refund_amount,omitempty"`      // 申请退款的金额（含运费）
	RefundPostAmount int    `json:"refund_post_amount,omitempty"` // 申请退的运费金额
	RefundVoucherNum int    `json:"refund_voucher_num,omitempty"` // 申请退款的卡券的数量
	ReasonCode       uint   `json:"reason_code,omitempty"`        // 申请售后原因码，枚举值
	RefuseTime       int64  `json:"refuse_time,omitempty"`        // 拒绝退款时间
	LatestRecord     string `json:"latest_record,omitempty"`      // 最近一条操作记录
}

func (this RefundRefused) Tag() string {
	return "204"
}
