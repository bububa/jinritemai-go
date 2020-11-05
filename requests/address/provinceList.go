package address

import (
	"github.com/bububa/jinritemai-go/client"
)

type ProvinceListRequest struct{}

func (this ProvinceListRequest) Method() string {
	return "address/provinceList"
}

func (this ProvinceListRequest) Params() map[string]interface{} {
	return nil
}

// 获取省列表
// 获取平台支持的省列表
func GetProvinceList(clt *client.Client) ([]Province, error) {
	req := &ProvinceListRequest{}
	var ret []Province
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
