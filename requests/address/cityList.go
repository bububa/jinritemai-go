package address

import (
	"github.com/bububa/jinritemai-go/client"
)

type CityListRequest struct {
	ProvinceID uint64 `json:"province_id,omitempty"` // 省ID
}

func (this CityListRequest) Method() string {
	return "address/cityList"
}

func (this CityListRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"province_id": this.ProvinceID,
	}
}

// 获取市列表
// 获取平台支持的市列表
func GetCityList(clt *client.Client, provinceID uint64) ([]City, error) {
	req := &CityListRequest{
		ProvinceID: provinceID,
	}
	var ret []City
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
