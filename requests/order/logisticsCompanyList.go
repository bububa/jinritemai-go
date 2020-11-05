package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type LogisticsCompanyListRequest struct{}

func (this LogisticsCompanyListRequest) Method() string {
	return "order/logisticsCompanyList"
}

func (this LogisticsCompanyListRequest) Params() map[string]interface{} {
	return nil
}

// 获取快递公司列表
// 获取平台支持的快递公司列表
// Tips：开发者需自行维护快递公司ID的映射关系
func GetLogisticsCompanyList(clt *client.Client) ([]LogisticsCompany, error) {
	req := &LogisticsCompanyListRequest{}
	var ret []LogisticsCompany
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
