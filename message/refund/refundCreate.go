package refund

// 买家发起售后申请消息
/*
触发场景
买家发起交易逆向申请时，推送此消息。具体场景如下：

订单未发货，买家申请整单退款时
订单已发货，买家申请售后仅退款时
订单已发货，买家申请售后退货时

aftersale_status售后单状态码列表
6   待商家处理
7   待买家退货
11  待商家收货
12  同意退款
27  拒绝售后申请
28  售后关闭
29  退货后拒绝退款

reason_code申请售后原因码列表
{
  "1": {
  "name": "多拍 / 错拍 / 不想要",
  "code": 1,
  "is_no_reason": true
  },
  "2": {
  "name": "未按约定时间发货",
  "code": 2
  },
  "3": {
  "name": "快递一直未送达",
  "code": 3
  },
  "4": {
  "name": "快递无跟踪记录",
  "code": 4
  },
  "5": {
  "name": "收到商品少件 / 错件 / 空包裹",
  "code": 5,
    "refund_part_type":2
  },
  "6": {
  "name": "不喜欢 / 效果不好",
  "code": 6,
  "is_no_reason": true
  },
  "7": {
  "name": "做工粗糙 / 有瑕疵 / 破损或污渍",
  "code": 7
  },
  "8": {
  "name": "功能故障",
  "code": 8
  },
  "9": {
  "name": "效果与商品描述不符",
  "code": 9
  },
  "10": {
  "name": "商品材质 / 品牌 / 外观等描述不符",
  "code": 10
  },
  "11": {
  "name": "生产日期 / 保质期 / 规格等描述不符",
  "code": 11
  },
  "12": {
  "name": "假冒品牌",
  "code": 12
  },
  "13": {
  "name": "协商一致退款",
  "code": 13
  },
  "14": {
  "name": "缺货",
  "code": 14
  },
  "15": {
  "name": "其他",
  "code": 15
  },
    "16":{
  "name": "大小／尺寸／重量与商品描述不符",
  "code": 16
    },
    "17":{
  "name": "生产日期／保质期与商品描述不符",
  "code": 17
    },
    "18":{
  "name": "品种／规格／成分等描述不符",
  "code": 18
    },
    "19":{
  "name": "商品腐烂/变质",
  "code": 19,
    "refund_part_type":1
    },
    "20":{
  "name": "少件／漏发",
  "code": 20,
    "refund_part_type":2
    },
    "21":{
  "name": "包装／商品破损",
  "code": 21,
    "refund_part_type":2

    },
    "22":{
  "name": "商家发错货",
  "code": 22
    },
    "23":{
  "name": "与商家协商一致退款",
  "code": 23
    },
    "24":{
  "name": "退运费",
  "code": 24
    },
    "25":{
  "name": "品种／产品／规格／成分等描述不符",
  "code": 25
    },
    "26":{
  "name": "商品腐烂／变质／死亡",
  "code": 26,
    "refund_part_type":1
    },
    "27":{
  "name": "商品变质／过期",
  "code": 27
    },
     "28":{
  "name": "规格等描述不符",
  "code": 28
    },
     "29":{
  "name": "收到商品少件／错件／空包裹",
  "code": 29,
    "refund_part_type":2
    },
         "30":{
  "name": "盗版",
  "code": 30
    },
  "31": {
  "name": "做工粗糙 / 有瑕疵 / 有污渍",
  "code": 31
  },
  "32": {
  "name": "产品破损 / 损坏",
  "code": 32
  },
  "33": {
  "name": "商品材质 / 外观等描述不符",
  "code": 33
  },
  "34": {
  "name": "使用后过敏",
  "code": 34
  }
}
*/
type RefundCreate struct {
	Sid              uint64 `json:"s_id,omitempty"`               // 子订单ID
	Pid              uint64 `json:"p_id,omitempty"`               // 父订单ID
	ShopID           uint64 `json:"shop_id,omitempty"`            // 店铺ID
	AftersaleID      uint64 `json:"aftersale_id,omitempty"`       // 售后单ID
	AftersaleStatus  uint   `json:"aftersale_status,omitempty"`   // 售后状态
	RefundAmount     int    `json:"refund_amount,omitempty"`      // 申请退款的金额（含运费）
	RefundPostAmount int    `json:"refund_post_amount,omitempty"` // 申请退的运费金额
	RefundVoucherNum int    `json:"refund_voucher_num,omitempty"` // 申请退款的卡券的数量
	ReasonCode       uint   `json:"reason_code,omitempty"`        // 申请售后原因码，枚举值
	ApplyTime        int64  `json:"apply_time,omitempty"`         // 售后申请时间
	LatestRecord     string `json:"latest_record,omitempty"`      // 最近一条操作记录
}

func (this RefundCreate) Tag() string {
	return "200"
}
