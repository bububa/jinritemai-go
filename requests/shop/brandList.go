package shop

import (
	"github.com/bububa/jinritemai-go/client"
)

type BrandListRequest struct{}

func (this BrandListRequest) Method() string {
	return "shop/brandList"
}

func (this BrandListRequest) Params() map[string]interface{} {
	return nil
}

type Brand struct {
	ID          uint64 `json:"id,omitempty"`                 // 品牌ID
	ChineseName string `json:"brand_chinese_name,omitempty"` // 品牌中文名
	EnglishName string `json:"brand_english_name,omitempty"` // 品牌英文名
	RegNum      string `json:"brand_reg_num,omitempty"`      // 商标注册号
}

func GetBrandList(clt *client.Client) ([]Brand, error) {
	req := &BrandListRequest{}
	var ret []Brand
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
