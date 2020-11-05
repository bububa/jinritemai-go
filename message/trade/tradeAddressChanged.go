package trade

// 买家收货信息变更消息
/*
触发场景
当买家收货地址被修改时，推送此消息

order_status订单正向状态码列表
1   1: 在线支付订单待支付
2: 货到付款订单待确认
2   备货中（只有此状态下，才可发货）
3   已发货
4   已取消
5   已完成：
1: 货到付款订单，商家发货后，用户收货、拒收或者15天无物流更新
2: 在线支付订单，用户确认收货
*/
type TradeAddressChanged struct {
	Pid         uint64       `json:"p_id,omitempty"`         // 父订单ID
	SIDs        uint64       `json:"s_ids,omitempty"`        // 子订单ID列表
	ShopID      uint64       `json:"shop_id,omitempty"`      // 店铺ID
	OrderStatus uint         `json:"order_status,omitempty"` // 父订单状态，卖家发货消息的status值为"2"
	OrderType   uint         `json:"order_type,omitempty"`   // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	UpdateTime  int64        `json:"update_time,omitempty"`  // 订单发货时间
	ReceiverMsg *ReceiverMsg `json:"receiver_msg,omitempty"` // 收货人详细信息
}

func (this TradeAddressChanged) Tag() string {
	return "105"
}
