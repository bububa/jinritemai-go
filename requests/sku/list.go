package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type ListRequest struct {
	ProductID uint64 `json:"product_id,omitempty"`
}

func (this ListRequest) Method() string {
	return "sku/list"
}

func (this ListRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"product_id": this.ProductID,
	}
}

func GetList(clt *client.Client, productID uint64) ([]SkuDetail, error) {
	req := &ListRequest{
		ProductID: productID,
	}
	var ret []SkuDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
