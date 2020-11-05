package trade

// 订单支付/确认消息
/*
触发场景
以下两种场景会推送此消息：

在线支付订单，当买家付款成功时
货到付款订单，当商家确认订单时（实际并无支付行为）

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
type TradePaid struct {
	Pid         uint64 `json:"p_id,omitempty"`         // 父订单ID
	SIDs        uint64 `json:"s_ids,omitempty"`        // 子订单ID列表
	ShopID      uint64 `json:"shop_id,omitempty"`      // 店铺ID
	OrderStatus uint   `json:"order_status,omitempty"` // 父订单状态，订单支付消息的status值为"2"
	OrderType   uint   `json:"order_type,omitempty"`   // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	PayType     uint   `json:"pay_type,omitempty"`     // 订单支付方式：0: 货到付款; 1: 微信; 2: 支付宝
	PayTime     int64  `json:"pay_time,omitempty"`     // 1: 在线订单支付时间; 2: 货到付款订单确认时间
	PayAmount   int    `json:"pay_amount,omitempty"`   // 订单实付金额
	Biz         uint   `json:"biz_id,omitempty"`       // 订单业务类型，表示买家从哪里看到的这个商品、产生了订单: 1: 鲁班广告; 2: 联盟; 4: 商城; 8:自主经营; 10: 线索通支付表单; 12: 抖音门店; 14: 抖+; 15: 穿山甲
}

func (this TradePaid) Tag() string {
	return "101"
}
