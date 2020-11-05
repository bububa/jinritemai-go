package product

import (
	"github.com/bububa/jinritemai-go/client"
)

type AddRequest struct {
	OutID           uint64 `json:"out_product_id,omitempty"`    // 外部商品id，接入方的商品id 需为数字字符串，max = int64
	Name            string `json:"name,omitempty"`              // 商品名称，最多30个字符，不能含emoj表情
	Pic             string `json:"pic,omitempty"`               // 商品轮播图，多张图片用 "|" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M 注："pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
	Description     string `json:"description,omitempty"`       // 商品描述，目前只支持图片。多张图片用 "|" 分开。不能用其他网站的文本粘贴，这样会出现css样式不兼容; 注："pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
	MarketPrice     int    `json:"market_price,omitempty"`      // 划线价，单位分
	DiscountPrice   int    `json:"discount_price,omitempty"`    // 售价，单位分
	FirstCid        uint64 `json:"first_cid,omitempty"`         // 一级类目
	SecondCid       uint64 `json:"second_cid,omitempty"`        // 二级类目
	ThirdCid        uint64 `json:"third_cid,omitempty"`         // 三级类目
	SpecID          uint64 `json:"spec_id,omitempty"`           // 规格id, 要先创建商品通用规格, 如颜色-尺寸
	SpecPic         string `json:"spec_pic,omitempty"`          // 主规格id, 如颜色-尺寸, 颜色就是主规格, 颜色如黑,白,黄,它们的id|图片url; 注："pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
	Mobile          string `json:"mobile,omitempty"`            // 客服号码
	Weight          int64  `json:"weight,omitempty"`            // 商品重量 (单位:克)。范围: 10克 - 9999990克
	ProductFormat   string `json:"product_format,omitempty"`    // 属性名称|属性值; 之间用|分隔, 多组之间用^分开
	PayType         uint   `json:"pay_type,omitempty"`          // 支持的支付方式：0货到付款 1在线支付 2两者都支持
	RecommendRemark string `json:"recommend_remark,omitempty"`  // 商家推荐语
	BrandID         uint64 `json:"brand_id,omitempty"`          // 品牌id (请求店铺授权品牌接口获取) (效果即为商品详情页面中的品牌字段) 哪些类目必需上传品牌
	PresellType     uint   `json:"presell_type,omitempty"`      // 预售类型，1-全款预售，0-非预售，默认0
	PresellDelay    uint   `json:"presell_delay,omitempty"`     // 预售结束后，几天发货，可以选择2-30
	PresellEndTime  string `json:"presell_end_time,omitempty"`  // 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
	DeliverDelayDay uint   `json:"deliver_delay_day,omitempty"` // 承诺发货时间，单位是天，可选值为: 2、3、5、7、10、15; 不传则默认为2天
	QuantityReport  string `json:"quantity_report,omitempty"`   // 商品创建和编辑操作支持设置质检报告链接,多个图片以逗号分隔
	ClassQuantity   string `json:"class_quantity,omitempty"`    // 商品创建和编辑操作支持设置品类资质链接,多个图片以逗号分隔
}

func (this AddRequest) Method() string {
	return "product/add"
}

func (this AddRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"name":           this.Name,
		"pic":            this.Pic,
		"description":    this.Description,
		"market_price":   this.MarketPrice,
		"discount_price": this.DiscountPrice,
		"first_cid":      this.FirstCid,
		"second_cid":     this.SecondCid,
		"third_cid":      this.ThirdCid,
		"spec_id":        this.SpecID,
		"weight":         this.Weight,
		"product_format": this.ProductFormat,
		"pay_type":       this.PayType,
	}
	if this.OutID > 0 {
		ret["out_product_id"] = this.OutID
	}
	if this.SpecPic != "" {
		ret["spec_pic"] = this.SpecPic
	}
	if this.Mobile != "" {
		ret["mobile"] = this.Mobile
	}
	if this.RecommendRemark != "" {
		ret["recommend_remark"] = this.RecommendRemark
	}
	if this.BrandID > 0 {
		ret["brand_id"] = this.BrandID
	}
	if this.PresellType != 0 {
		ret["presell_type"] = 1
	} else {
		ret["presell_type"] = 0
	}
	if this.PresellDelay >= 2 && this.PresellDelay <= 30 {
		ret["presell_delay"] = this.PresellDelay
	}
	if this.PresellEndTime != "" {
		ret["presell_end_time"] = this.PresellEndTime
	}
	if this.DeliverDelayDay == 2 ||
		this.DeliverDelayDay == 3 ||
		this.DeliverDelayDay == 5 ||
		this.DeliverDelayDay == 7 ||
		this.DeliverDelayDay == 10 ||
		this.DeliverDelayDay == 15 {
		ret["deliver_delay_day"] = this.DeliverDelayDay
	}
	if this.QuantityReport != "" {
		ret["quantity_report"] = this.QuantityReport
	}
	if this.ClassQuantity != "" {
		ret["class_quantity"] = this.ClassQuantity
	}
	return ret
}

// 添加商品
// 创建商品的接口，商品添加成功后会自动进入平台的异步机审流程，机审完成后将自动上架。如果添加成功后，没有在后台页面上找到商品，可前往草稿箱查看。
// Tips：
// 1、"pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
// 2、此接口只能创建普通实物商品，不能创建虚拟商品
func Add(clt *client.Client, req *AddRequest) (*ProductDetail, error) {
	var ret ProductDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
