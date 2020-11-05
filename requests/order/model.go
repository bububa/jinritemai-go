package order

import (
	"encoding/json"
)

type OrderDetail struct {
	ID               string       `json:"order_id,omitempty"`           // 父订单ID，带大写字母A
	ShopID           uint64       `json:"shop_id,omitempty"`            // 店铺ID
	OpenID           string       `json:"open_id,omitempty"`            // 在抖音小程序下单时，买家的抖音小程序ID
	OrderStatus      int          `json:"order_status,omitempty"`       // 订单状态
	OrderType        uint         `json:"order_type,omitempty"`         // 订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)
	Child            []ChildOrder `json:"child,omitempty"`              // 子订单列表
	ChildNum         int          `json:"child_num,omitempty"`          // 子订单数量
	PostAddrs        []PostAddr   `json:"post_addr,omitempty"`          // 收件人地址
	PostCode         string       `json:"post_code,omitempty"`          // 邮政编码
	PostReceiver     string       `json:"post_receiver,omitempty"`      // 收件人姓名
	PostTel          string       `json:"post_tel,omitempty"`           // 收件人电话
	BuyerWords       string       `json:"buyer_words,omitempty"`        // 买家备注
	SellerWords      string       `json:"seller_words,omitempty"`       // 卖家备注
	LogisticsID      uint64       `json:"logistics_id,omitempty"`       // 物流公司ID 各ID对应的物流公司
	LogisticsCode    string       `json:"logistics_code,omitempty"`     // 物流单号
	LogisticsTime    json.Number  `json:"logistics_time,omitempty"`     // 发货时间。未发货时为"0"，已发货返回秒级时间戳
	ReceiptTime      json.Number  `json:"receipt_time,omitempty"`       // 收货时间。未收货时为"0"，已发货返回秒级时间戳
	CreateTime       json.Number  `json:"create_time,omitempty"`        // 订单创建时间，例如 "1512553757"
	UpdateTime       json.Number  `json:"update_time,omitempty"`        // 订单更新时间
	ExpShipTime      json.Number  `json:"exp_ship_time,omitempty"`      // 订单最晚发货时间，例如 1512553887
	CancelReason     string       `json:"cancel_reason,omitempty"`      // 订单取消原因
	PayType          uint         `json:"pay_type,omitempty"`           // 支付类型 (0：货到付款，1：微信，2：支付宝）
	PayTime          string       `json:"pay_time,omitempty"`           // 支付时间 (pay_type为0货到付款时, 此字段为空)，例如"2018-06-01 12:00:00"
	PostAmount       int          `json:"pay_amount,omitempty"`         // 邮费金额 (单位: 分)
	CouponAmount     int          `json:"coupon_amount,omitempty"`      // 平台优惠券金额 (单位: 分)
	ShopCouponAmount int          `json:"shop_coupon_amount,omitempty"` // 商家优惠券金额 (单位: 分)
	CouponMetaID     json.Number  `json:"coupon_meta_id,omitempty"`     // 优惠券id
	Coupons          []CouponInfo `json:"coupon_info,omitempty"`        // 优惠券详情 (type为优惠券类型，具体如下表所示, credit为优惠金额,单位分)
	TotalAmount      int          `json:"order_total_amount,omitempty"` // 订单实付金额（不包含运费）
	IsComment        string       `json:"is_comment,omitempty"`         // 是否评价 :1已评价，0未评价
	UrgeCnt          int          `json:"urge_cnt,omitempty"`           // 催单次数
	BType            string       `json:"b_type,omitempty"`             // 订单APP渠道
	SubBType         uint         `json:"sub_b_type,omitempty"`         // 订单来源类型: 0:未知; 1:app; 2:小程序; 3:h5
	CBiz             uint         `json:"c_biz,omitempty"`              // 订单业务类型
	IsInsurance      string       `json:"is_insurance,omitempty"`       // 是否有退货运费险
}

type ChildOrder struct {
	ID                    string                 `json:"order_id,omitempty"`               // 父订单ID，带大写字母A
	ShopID                uint64                 `json:"shop_id,omitempty"`                // 店铺ID
	Pid                   string                 `json:"pid,omitempty"`                    // 该子订单对应的父订单ID
	FinalStatus           uint                   `json:"final_status,omitempty"`           // 子订单状态
	OpenID                string                 `json:"open_id,omitempty"`                // 在抖音小程序下单时，买家的抖音小程序ID
	ProductID             uint64                 `json:"product_id,omitempty"`             // 商品ID
	ProductName           string                 `json:"product_name,omitempty"`           // 商品名称
	ProductPic            string                 `json:"product_pic,omitempty"`            // 头图，商品主图第一张
	ComboID               uint64                 `json:"combo_id,omitempty"`               // 该子订单购买的商品 sku_id
	ComboAmount           int                    `json:"combo_amount,omitempty"`           // 该子订单所购买的sku的售价
	ComboNum              int                    `json:"combo_num,omitempty"`              // 该子订单所购买的sku的数量
	Code                  string                 `json:"code,omitempty"`                   // 该子订单购买的商品的编码 code
	SpecDesc              []SpecDesc             `json:"spec_desc,omitempty"`              // 该子订单所属商品规格描述
	PostAddrs             []PostAddr             `json:"post_addr,omitempty"`              // 收件人地址
	PostCode              string                 `json:"post_code,omitempty"`              // 邮政编码
	PostReceiver          string                 `json:"post_receiver,omitempty"`          // 收件人姓名
	PostTel               string                 `json:"post_tel,omitempty"`               // 收件人电话
	BuyerWords            string                 `json:"buyer_words,omitempty"`            // 买家备注
	SellerWords           string                 `json:"seller_words,omitempty"`           // 卖家备注
	LogisticsID           uint64                 `json:"logistics_id,omitempty"`           // 物流公司ID 各ID对应的物流公司
	LogisticsCode         string                 `json:"logistics_code,omitempty"`         // 物流单号
	LogisticsTime         json.Number            `json:"logistics_time,omitempty"`         // 发货时间。未发货时为"0"，已发货返回秒级时间戳
	ReceiptTime           json.Number            `json:"receipt_time,omitempty"`           // 收货时间。未收货时为"0"，已发货返回秒级时间戳
	OrderType             uint                   `json:"order_type,omitempty"`             // 订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)
	PreSaleType           uint                   `json:"pre_sale_type,omitempty"`          // 订单预售类型 (1:全款预售订单)
	CreateTime            json.Number            `json:"create_time,omitempty"`            // 订单创建时间，例如 "1512553757"
	UpdateTime            json.Number            `json:"update_time,omitempty"`            // 订单更新时间
	ExpShipTime           json.Number            `json:"exp_ship_time,omitempty"`          // 订单最晚发货时间，例如 1512553887
	CancelReason          string                 `json:"cancel_reason,omitempty"`          // 订单取消原因
	PayType               uint                   `json:"pay_type,omitempty"`               // 支付类型 (0：货到付款，1：微信，2：支付宝）
	PayTime               string                 `json:"pay_time,omitempty"`               // 支付时间 (pay_type为0货到付款时, 此字段为空)，例如"2018-06-01 12:00:00"
	PostAmount            int                    `json:"pay_amount,omitempty"`             // 邮费金额 (单位: 分)
	CouponAmount          int                    `json:"coupon_amount,omitempty"`          // 平台优惠券金额 (单位: 分)
	ShopCouponAmount      int                    `json:"shop_coupon_amount,omitempty"`     // 商家优惠券金额 (单位: 分)
	CouponMetaID          json.Number            `json:"coupon_meta_id,omitempty"`         // 优惠券id
	Coupons               []CouponInfo           `json:"coupon_info,omitempty"`            // 优惠券详情 (type为优惠券类型，具体如下表所示, credit为优惠金额,单位分)
	TotalAmount           int                    `json:"total_amount,omitempty"`           // 订单实付金额（不包含运费）
	IsComment             string                 `json:"is_comment,omitempty"`             // 是否评价 :1已评价，0未评价
	UrgeCnt               int                    `json:"urge_cnt,omitempty"`               // 催单次数
	BType                 string                 `json:"b_type,omitempty"`                 // 订单APP渠道
	SubBType              uint                   `json:"sub_b_type,omitempty"`             // 订单来源类型: 0:未知; 1:app; 2:小程序; 3:h5
	CBiz                  uint                   `json:"c_biz,omitempty"`                  // 订单业务类型
	IsInsurance           string                 `json:"is_insurance,omitempty"`           // 是否有退货运费险
	WarehouseID           uint64                 `json:"warehouse_id,omitempty"`           // 仓库ID
	OutWarehouseID        uint64                 `json:"out_warehouse_id,omitempty"`       // 仓库外部ID
	WarehouseSuplier      string                 `json:"wharehouse_suplier,omitempty"`     // 供应商ID
	ShopFullCampaigns     []ShopFullCampaign     `json:"shop_full_campaign,omitempty"`     // 该子订单所使用的店铺满减优惠信息
	ShopCampaignID        uint64                 `json:"shop_campaign_id,omitempty"`       // 店铺满减活动ID
	ShopFullAmount        int                    `json:"shop_full_amount,omitempty"`       // 分摊到该子订单上的满减金额，单位：分
	PlatformFullCampaigns []PlatformFullCampaign `json:"platform_full_campaign,omitempty"` // 该子订单所使用的平台满减优惠信息
	PlatformCampaignID    uint64                 `json:"platform_campaign_id,omitempty"`   // 平台满减活动ID
	PlatformFullAmount    int                    `json:"platform_full_amount,omitempty"`   // 该子订单所使用的平台满减金额，单位：分
	CostSources           []CostSource           `json:"cost_source,omitempty"`            // 平台满减分摊详情: 1、"promotion_amount"：店铺或平台承担的金额，单位：分; 2、"source_type"：1-店铺承担，2-平台承担
	PromotionAmount       int                    `json:"promotion_amount,omitempty"`       // 店铺或平台承担的金额，单位：分
	SourceType            uint                   `json:"source_type,omitempty"`            // 优惠由谁承担：1-店铺承担，2-平台承担
}

type PostAddr struct {
	Town     Location `json:"town,omitempty"`
	City     Location `json:"city,omitempty"`
	Province Location `json:"province,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}

type Location struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CouponInfo struct {
	ID       uint64 `json:"id,omitempty"`          // 优惠券批次号
	Name     string `json:"name,omitempty"`        // 优惠券名称
	Desc     string `json:"description,omitempty"` // 优惠券描述
	Credit   int    `json:"credit,omitempty"`      // 满减/直减券金额(单位分)
	Type     uint   `json:"type,omitempty"`        // 优惠券类型：1-平台折扣券 (平台承担); 2-平台直减券 (平台承担); 3-平台满减券 (平台承担); 11-品类折扣券 (暂未开放); 12-品类直减券 (暂未开放); 13-品类满减券 (暂未开放); 21-店铺折扣券 (店铺承担); 22-店铺直减券 (店铺承担); 23-店铺满减券 (店铺承担); 31-渠道折扣券 (平台承担); 32-渠道直减券 (平台承担); 33-渠道满减券 (平台承担); 41-单品折扣券 (店铺承担); 42-单品直减券 (店铺承担); 43-单品满减券 (店铺承担)
	Discount int    `json:"discount,omitempty"`    // 折扣券折扣
	PayType  uint   `json:"pay_type,omitempty"`    // 支付类型 (0：货到付款，1：微信，2：支付宝）
}

type SpecDesc struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ShopFullCampaign struct {
	ID     uint64 `json:"shop_campaign_id,omitempty"`
	Amount int    `json:"shop_full_amount,omitempty"`
}

type PlatformFullCampaign struct {
	ID     uint64 `json:"platform_campaign_id,omitempty"`
	Amount int    `json:"platform_full_amount,omitempty"`
}

type CostSource struct {
	PromotionAmount int  `json:"promotion_amount,omitempty"` // 店铺或平台承担的金额，单位：分
	SourceType      uint `json:"source_type,omitempty"`      // 1-店铺承担，2-平台承担
}

type ServiceDetail struct {
	ID                uint64 `json:"id,omitempty"`                  // 服务请求ID
	OrderID           uint64 `json:"order_id,omitempty"`            // 订单号
	Reply             string `json:"reply,omitempty"`               // 商家答复内容
	Detail            string `json:"detail,omitempty"`              // 服务单详情
	CreateTime        string `json:"create_time,omitempty"`         // 服务创建时间
	OperateStatus     uint   `json:"operate_status,omitempty"`      // 操作状态
	OperateStatusDesc string `json:"operate_status_desc,omitempty"` // 操作状态含义
	OperatorID        uint64 `json:"operator_id,omitempty"`         // 操作人id(客服)
	ReplyTime         string `json:"reply_time,omitempty"`          // 回复时间
	ShopID            uint64 `json:"shop_id,omitempty"`             // 店铺id
	SupplyID          uint64 `json:"supply_id,omitempty"`           // 供应商id
	ReplyOpID         uint64 `json:"reply_op_id,omitempty"`         // 回复人id
}

type LogisticsCompany struct {
	ID   uint64 `json:"id,omitempty"`   // 快递公司ID
	Name string `json:"name,omitempty"` // 快递公司名称
}

type InsPolicy struct {
	InsPolicyNo          string `json:"ins_policy_no,omitempty"`          // 保单号
	ApproximateReturnFee int    `json:"approximate_return_fee,omitempty"` // 预计退保费用，单位分
	TotalPremium         int    `json:"total_premium,omitempty"`          // 总保费，单位分
	UserPremium          int    `json:"user_premium,omitempty"`           // 用户支付的保费，单位分
	MerchantPremium      int    `json:"merchant_premium,omitempty"`       // 商家支付的保费，单位分
	PlatformPremium      int    `json:"platform_premium,omitempty"`       // 平台承担的保费，单位分
	Status               uint   `json:"status,omitempty"`                 // 保单状态: 0: 初始化; 1: 待生效; 2: 保障中; 3: 赔付完成; 4: 保单失效; 5: 已经取消
	AppealStatus         uint   `json:"appeal_status,omitempty"`          // 申述状态: 0: 初始化; 1: 申诉处理中; 2: 申诉成功; 3: 申诉失败
	ClaimStatus          uint   `json:"claim_status,omitempty"`           // 理赔状态: 0: 初始化; 1: 理赔中; 2: 理赔成功; 3: 理赔失败
	InsEnsuredTimeBegin  string `json:"ins_ensured_time_begin,omitempty"` // 出保时间
	InsEnsuredTimeEnd    string `json:"ins_ensured_time_end,omitempty"`   // 保险过期时间
	CompanyName          string `json:"company_name,omitempty"`           // 保险公司名称
	IsAllowAppeal        bool   `json:"is_allow_appeal,omitempty"`        // 只有在保单状态为赔付失败的时候并允许申诉 true，其他情况 false
	FailReason           string `json:"fail_reason,omitempty"`            // 理赔或申述失败原因
	InsuranceHotline     string `json:"insurance_hotline,omitempty"`      // 保险客服电话
}
