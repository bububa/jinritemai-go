package product

import (
	"github.com/bububa/jinritemai-go/requests/spec"
)

type ProductDetail struct {
	ID              uint64      `json:"product_id,omitempty"`       // 商品ID，整型格式
	IDStr           string      `json:"product_id_str,omitempty"`   // 商品ID，字符串格式
	OutID           uint64      `json:"out_product_id,omitempty"`   // 商品外部ID
	Name            string      `json:"name,omitempty"`             // 商品名称
	Description     string      `json:"description,omitempty"`      // 详情html
	MarketPrice     int         `json:"market_price,omitempty"`     // 划线价，单位分
	DiscountPrice   int         `json:"discount_price,omitempty"`   // 售价，单位分
	Status          uint        `json:"status,omitempty"`           // 商品上下架状态：0上架 1下架
	SpecID          uint64      `json:"spec_id,omitempty"`          // 规格id
	CheckStatus     uint        `json:"check_status,omitempty"`     // 商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁
	Mobile          uint        `json:"mobile,omitempty"`           // 手机号
	FirstCid        uint64      `json:"first_cid,omitempty"`        // 一级类目
	SecondCid       uint64      `json:"second_cid,omitempty"`       // 二级类目
	ThirdCid        uint64      `json:"third_cid,omitempty"`        // 三级类目
	PayType         uint        `json:"pay_type,omitempty"`         // 支持的支付方式：0货到付款 1在线支付 2两者都支持
	RecommendRemark string      `json:"recommend_remark,omitempty"` // 商家推荐语
	Extra           string      `json:"extra,omitempty"`            // 扩展字段
	CreateTime      string      `json:"create_time,omitempty"`      // 创建时间
	UpdateTime      string      `json:"update_time,omitempty"`      // 更新时间
	Pics            []string    `json:"pic,omitempty"`              // 商品主图
	ProductFormat   string      `json:"product_format,omitempty"`   // 属性名称|属性值之间用|分隔, 多组之间用^分开
	SpecPics        []SpecPic   `json:"spec_pics,omitempty"`        // 规格图片
	SpecPrices      []SpecPrice `json:"spec_prices,omitempty"`      // sku详细信息
	Specs           []spec.Spec `json:"specs,omitempty"`            //
	Image           string      `json:"img,omitempty"`              // 头图，主图第一张
}

type SpecPic struct {
	SpecDetailID uint64 `json:"spec_detail_id,omitempty"` // 子规格ID，比如“颜色-尺寸”这个规格组里的“颜色”里的“白色”对应的ID
	Pic          string `json:"pic,omitempty"`            // 规格图片url
}

type SpecPrice struct {
	SkuID         uint64   `json:"sku_id,omitempty"`          // skuID
	OutID         uint64   `json:"out_sku_id,omitempty"`      // sku外部
	SpecDetailIDs []uint64 `json:"spec_detail_ids,omitempty"` // 上述sku所使用的子规格ID，比如“白色”、“大”、“女式”的ID
	StockNum      int      `json:"stock_num,omitempty"`       // stock_num
	Price         int      `json:"price,omitempty"`           // sku售价
	Code          string   `json:"code,omitempty"`            // 商家外部编码
}
