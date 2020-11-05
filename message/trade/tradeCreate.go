package trade

// 订单创建消息
/*
触发场景
当买家下单，系统生成订单时，推送此消息

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
type TradeCreate struct {
	Pid         uint64 `json:"p_id,omitempty"`         // 父订单ID
	SIDs        uint64 `json:"s_ids,omitempty"`        // 子订单ID列表
	ShopID      uint64 `json:"shop_id,omitempty"`      // 店铺ID
	CreateTime  int64  `json:"create_time,omitempty"`  // 订单创建时间
	OrderStatus uint   `json:"order_status,omitempty"` // 父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint   `json:"order_type,omitempty"`   // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	Biz         uint   `json:"biz_id,omitempty"`       // 订单业务类型，表示买家从哪里看到的这个商品、产生了订单: 1: 鲁班广告; 2: 联盟; 4: 商城; 8:自主经营; 10: 线索通支付表单; 12: 抖音门店; 14: 抖+; 15: 穿山甲
}

func (this TradeCreate) Tag() string {
	return "100"
}
