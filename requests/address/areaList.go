package address

import (
	"github.com/bububa/jinritemai-go/client"
)

type AreaListRequest struct {
	CityID uint64 `json:"city_id,omitempty"` // 市ID
}

func (this AreaListRequest) Method() string {
	return "address/areaList"
}

func (this AreaListRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"city_id": this.CityID,
	}
}

// 获取区列表
// 获取平台支持的区列表
func GetAreaList(clt *client.Client, cityID uint64) ([]Area, error) {
	req := &AreaListRequest{
		CityID: cityID,
	}
	var ret []Area
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
