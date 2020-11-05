package product

import (
	"github.com/bububa/jinritemai-go/client"
)

type EditRequest struct {
	ProductID      uint64 `json:"product_id,omitempty"`       // 商品ID
	Name           string `json:"name,omitempty"`             // 商品名称，最多30个字符，不能含emoj表情
	SpecID         uint64 `json:"spec_id,omitempty"`          // 规格id, 要先创建商品通用规格, 如颜色-尺寸; 注意 : 修改后将导致原有的sku绑定关系失效，已上架商品会自动下架，请谨慎操作
	Pic            string `json:"pic,omitempty"`              // 商品轮播图，多张图片用 "|" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M 注："pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
	Description    string `json:"description,omitempty"`      // 商品描述，目前只支持图片。多张图片用 "|" 分开。不能用其他网站的文本粘贴，这样会出现css样式不兼容; 注："pic"、"description"、"spec_pic"这三个字段里的传入的图片数量总和，不得超过50张。超过会报错：图片转链失败
	FirstCid       uint64 `json:"first_cid,omitempty"`        // 一级类目
	SecondCid      uint64 `json:"second_cid,omitempty"`       // 二级类目
	ThirdCid       uint64 `json:"third_cid,omitempty"`        // 三级类目
	Mobile         string `json:"mobile,omitempty"`           // 客服号码
	PresellType    *uint  `json:"presell_type,omitempty"`     // 预售类型，1-全款预售，0-非预售，默认0
	PresellDelay   uint   `json:"presell_delay,omitempty"`    // 预售结束后，几天发货，可以选择2-30
	PresellEndTime string `json:"presell_end_time,omitempty"` // 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
	Commit         bool   `json:"commit,omitempty"`           // "1"：编辑后立即提交审核；"2"：编辑后仅保存，不提审
}

func (this EditRequest) Method() string {
	return "product/edit"
}

func (this EditRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"product_id": this.ProductID,
	}
	if this.Name != "" {
		ret["name"] = this.Name
	}
	if this.SpecID > 0 {
		ret["spec_id"] = this.SpecID
	}
	if this.Pic != "" {
		ret["pic"] = this.Pic
	}
	if this.Description != "" {
		ret["description"] = this.Description
	}
	if this.FirstCid > 0 {
		ret["first_cid"] = this.FirstCid
	}
	if this.SecondCid > 0 {
		ret["second_cid"] = this.SecondCid
	}
	if this.ThirdCid > 0 {
		ret["third_cid"] = this.ThirdCid
	}
	if this.Mobile != "" {
		ret["mobile"] = this.Mobile
	}
	if this.PresellType != nil {
		ret["presell_type"] = this.PresellType
	}
	if this.PresellDelay >= 2 && this.PresellDelay <= 30 {
		ret["presell_delay"] = this.PresellDelay
	}
	if this.PresellEndTime != "" {
		ret["presell_end_time"] = this.PresellEndTime
	}
	if this.Commit {
		ret["commit"] = "1"
	} else {
		ret["commit"] = "0"
	}
	return ret
}

// 编辑商品
// 编辑商品相关信息。编辑提交后默认自动提交审核，审核通过后，更新线上数据。
func Edit(clt *client.Client, req *EditRequest) error {
	return clt.Do(req, nil)
}
