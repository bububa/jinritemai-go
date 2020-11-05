package trade

// 发货物流变更消息
/*
触发场景
订单已发货，发货物流信息变更时，推送此消息
*/
type TradeLogisticsChanged struct {
	Pid          uint64        `json:"p_id,omitempty"`          // 父订单ID
	SIDs         uint64        `json:"s_ids,omitempty"`         // 子订单ID列表
	ShopID       uint64        `json:"shop_id,omitempty"`       // 店铺ID
	OrderStatus  uint          `json:"order_status,omitempty"`  // 父订单状态，卖家发货消息的status值为"3"
	OrderType    uint          `json:"order_type,omitempty"`    // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	UpdateTime   int64         `json:"update_time,omitempty"`   // 订单发货时间
	ReceiverMsg  *ReceiverMsg  `json:"receiver_msg,omitempty"`  // 收货人详细信息
	LogisticsMsg *LogisticsMsg `json:"logistics_msg,omitempty"` // 发货物流信息
	Biz          uint          `json:"biz_id,omitempty"`        // 订单业务类型，表示买家从哪里看到的这个商品、产生了订单: 1: 鲁班广告; 2: 联盟; 4: 商城; 8:自主经营; 10: 线索通支付表单; 12: 抖音门店; 14: 抖+; 15: 穿山甲
}

func (this TradeLogisticsChanged) Tag() string {
	return "104"
}
