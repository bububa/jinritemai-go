package trade

// 订单取消消息
/*
触发场景
订单被取消时推送此消息。取消订单的场景如下：

货到付款订单且订单状态为「待确认」，买家和商家可取消订单
货到付款订单且订单状态为「备货中」，买家、商家和平台客服可取消订单
货到付款订单且订单状态为「已发货」，物流状态为拒收或退回
在线支付订单且订单状态为「待付款」，买家或平台客服可取消该订单
在线支付订单且订单状态为「待付款」，超时未支付，订单自动取消
在线支付订单且订单状态为「备货中」，触发平台风控规则，订单被取消
*/
type TradeCanceled struct {
	Pid          uint64 `json:"p_id,omitempty"`          // 父订单ID
	SIDs         uint64 `json:"s_ids,omitempty"`         // 子订单ID列表
	ShopID       uint64 `json:"shop_id,omitempty"`       // 店铺ID
	CancelTime   int64  `json:"cancel_time,omitempty"`   // 订单取消时间
	CancelReason string `json:"cancel_reason,omitempty"` // 取消原因
	OrderStatus  uint   `json:"order_status,omitempty"`  // 父订单状态，订单创建消息的order_status值为"4"
	OrderType    uint   `json:"order_type,omitempty"`    // 订单类型：0: 实物; 2: 普通虚拟; 4: poi核销; 5: 三方核销; 6: 服务市场
	Biz          uint   `json:"biz_id,omitempty"`        // 订单业务类型，表示买家从哪里看到的这个商品、产生了订单: 1: 鲁班广告; 2: 联盟; 4: 商城; 8:自主经营; 10: 线索通支付表单; 12: 抖音门店; 14: 抖+; 15: 穿山甲
}

func (this TradeCanceled) Tag() string {
	return "106"
}
