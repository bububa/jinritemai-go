package trade

// 交易完成消息
/*
触发场景
买家确认收货或系统自动确认收货，且父订单状态变为「已完成」时，推送此消息
*/
type TradeSuccess struct {
	Pid          uint64 `json:"p_id,omitempty"`          // 父订单ID
	SIDs         uint64 `json:"s_ids,omitempty"`         // 子订单ID列表
	ShopID       uint64 `json:"shop_id,omitempty"`       // 店铺ID
	CompleteTime int64  `json:"complete_time,omitempty"` // 交易完成时间
	OrderStatus  uint   `json:"order_status,omitempty"`  // 父订单状态，订单创建消息的order_status值为"5"
	OrderType    uint   `json:"order_type,omitempty"`    // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	Biz          uint   `json:"biz_id,omitempty"`        // 订单业务类型，表示买家从哪里看到的这个商品、产生了订单: 1: 鲁班广告; 2: 联盟; 4: 商城; 8:自主经营; 10: 线索通支付表单; 12: 抖音门店; 14: 抖+; 15: 穿山甲
}

func (this TradeSuccess) Tag() string {
	return "103"
}
