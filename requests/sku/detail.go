package sku

import (
	"github.com/bububa/jinritemai-go/client"
)

type DetailRequest struct {
	SkuID uint64 `json:"sku_id,omitempty"`
}

func (this DetailRequest) Method() string {
	return "sku/detail"
}

func (this DetailRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"sku_id": this.SkuID,
	}
}

func GetDetail(clt *client.Client, skuID uint64) (*SkuDetail, error) {
	req := &DetailRequest{
		SkuID: skuID,
	}
	var ret SkuDetail
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
