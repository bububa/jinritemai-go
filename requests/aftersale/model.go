package aftersale

type RefundProcessDetail struct {
	Order             *OrderInfo           `json:"order_info,omitempty"`
	Process           *ProcessInfo         `json:"process_info,omitempty"`        // 退款的当前实时处理信息
	Aftersale         *AftersaleInfo       `json:"aftersale_info,omitempty"`      // 售后单相关信息
	Logs              []Log                `json:"logs,omitempty"`                // 退款的历史处理日志信息
	RefundTotalAmount int                  `json:"refund_total_amount,omitempty"` // 退款总金额，单位是分
	RefundPostAmount  int                  `json:"refund_post_amount,omitempty"`  // 退的运费金额，单位是分
	ProcessList       []ProcessInfoWithLog `json:"process_info_list,omitempty"`   // 当前子订单的售后记录（1次或多次）
}

type OrderInfo struct {
	OrderID     uint64 `json:"order_id,omitempty"`     // 订单号
	Pid         uint64 `json:"pid,omitempty"`          // 父订单号
	OrderStatus uint   `json:"order_status,omitempty"` // order_status
	FinalStatus uint   `json:"final_status,omitempty"`
	StatusDesc  string `json:"status_desc,omitempty"`  // 退款状态对应的描述文案
	CreateTime  string `json:"create_time,omitempty"`  // 订单创建事件
	ReceiptTime string `json:"receipt_time,omitempty"` // 订单确认收货时间,可能为空字符串
	ComboNum    int    `json:"combo_num,omitempty"`    // 下单的sku购买数量
	ComboAmount int    `json:"combo_amount,omitempty"` // 下单时的sku单价
	TotalAmount int    `json:"total_amount,omitempty"` // 下单时sku对应的总价
	PayAmount   int    `json:"pay_amount,omitempty"`   // 下单时改单实际支付的金额(sku总价扣除优惠后的)
	ShopID      uint64 `json:"shop_id,omitempty"`      // 店铺id
	ProductID   uint64 `json:"product_id,omitempty"`   // 商品id
	ProductName string `json:"product_name,omitempty"` // 商品名字
	ProductImg  string `json:"product_img,omitempty"`  // 商品图片
}

/*
aftersale_status：售后单状态，具体释义见下表
6   待商家处理
7   待买家退货
11  待商家收货
12  商家同意退款：
1、发货前退款，商家同意退款
2、发货后仅退款，商家同意退款
3、发货后退货退款，商家确认收货
27  拒绝售后申请
28  售后关闭
29  退货后商家拒绝

refund_type：表示金额怎么退给买家，具体释义见下表
0   原路退回
1   货到付款线下退款
2   备用金退款
3   保证金退款
4   无需退款

refund_status：表示退款到账进度，具体释义见下表
0   无需退款
1   待退款
2   退款中
3   退款成功
4   退款失败
*/
type ProcessInfo struct {
	ApplyTime       string   `json:"apply_time,omitempty"`
	Reason          string   `json:"reason,omitempty"`
	TypeDesc        string   `json:"type_desc,omitempty"`
	ApplyRemark     string   `json:"apply_remark,omitempty"`
	Evidence        []string `json:"evidence,omitempty"`
	LogisticsTime   string   `json:"logistics_time,omitempty"`   // 买家填写退货物流时间
	LogisticsCode   string   `json:"logistics_code,omitempty"`   // 退货物流单号
	LogisticsName   string   `json:"logistics_name,omitempty"`   // 退货快递公司
	AftersaleID     uint64   `json:"aftersale_id,omitempty"`     // 售后单ID
	AftersaleStatus uint     `json:"aftersale_status,omitempty"` // 售后单状态
	RefundType      uint     `json:"refund_type,omitempty"`      // 表示金额怎么退给买家
	RefundStatus    uint     `json:"refund_status,omitempty"`    // 表示退款到账进度
}

/*
aftersale_type：售后类型，具体释义见下表
0   售后退货退款
1   售后仅退款
2   售前仅退款
*/
type AftersaleInfo struct {
	AftersaleType     uint   `json:"aftersale_type,omitempty"`      // 售后类型：0售后退货退款，1售后仅退款，2售前仅退款
	AfterSaleTypeText string `json:"aftersale_type_text,omitempty"` //
}

type Log struct {
	CreateTime string   `json:"create_time,omitempty"`
	Action     string   `json:"action,omitempty"`
	Desc       string   `json:"desc,omitempty"`
	Operator   string   `json:"operator,omitempty"`
	Evidence   []string `json:"evidence,omitempty"`
}

type ProcessInfoWithLog struct {
	Process *ProcessInfo `json:"process_info,omitempty"`
	Logs    []Log        `json:"logs,omitempty"`
}
