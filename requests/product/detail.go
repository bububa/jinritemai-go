package product

import (
	"github.com/bububa/jinritemai-go/client"
)

type DetailRequest struct {
	ProductID uint64 `json:"product_id,omitempty"` // 商品id
	ShowDraft bool   `json:"show_draft,omitempty"` // "true"：读取草稿数据；"false"：读取上架数据
}

func (this DetailRequest) Method() string {
	return "product/detail"
}

func (this DetailRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"product_id": this.ProductID,
		"show_draft": this.ShowDraft,
	}
}

func GetDetail(clt *client.Client, productID uint64, showDraft bool) (*ProductDetail, error) {
	req := &DetailRequest{
		ProductID: productID,
		ShowDraft: showDraft,
	}
	var ret ProductDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
