package trade

// 卖家发货消息
/*
触发场景
买家付款后，卖家对所有商品发货完成，且父订单状态为「已发货」时，推送此消息

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
type TradeSellerShip struct {
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

func (this TradeSellerShip) Tag() string {
	return "102"
}

type ReceiverMsg struct {
	Name string `json:"name,omitempty"` // 收货人姓名
	Tel  string `json:"tel,omitempty"`  // 收货人手机号
	Addr string `json:"addr,omitempty"` // 收货地址
}

type LogisticsMsg struct {
	ExpressCompanyID uint64 `json:"express_company_id,omitempty"` // 发货快递公司
	LogisticsCode    string `json:"logistics_code,omitempty"`     // 发货物流单号
}
