package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type ServiceListRequest struct {
	Page      int   `json:"page,omitempty"`       // 第几页（第一页为0，最大为99）
	Size      int   `json:"size,omitempty"`       // 每页返回条数，最多支持100条
	Status    *uint `json:"status,omitempty"`     // 1、不传代表获取全部服务请求 2、操作状态：0待处理，1已处理
	StartTime int64 `json:"start_time,omitempty"` // 开始时间
	EndTime   int64 `json:"end_time,omitempty"`   // 结束时间，必须大于等于开始时间
	Supply    uint  `json:"supply,omitempty"`     // 是否获取分销商服务申请，0获取本店铺的服务申请，1获取分销商的服务申请
}

func (this ServiceListRequest) Method() string {
	return "order/serviceList"
}

func (this ServiceListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"start_time": this.StartTime,
		"end_time":   this.EndTime,
		"supply":     this.Supply,
		"page":       this.Page,
	}
	if this.Status != nil {
		ret["status"] = *this.Status
	}
	if this.Size > 0 {
		ret["size"] = this.Size
	}
	return ret
}

type ServiceListResponse struct {
	Total int             `json:"total,omitempty"`
	List  []ServiceDetail `json:"list,omitempty"`
}

// 获取服务请求列表
// 获取客服向店铺发起的服务请求列表
func GetServiceList(clt *client.Client, req *ServiceListRequest) (*ServiceListResponse, error) {
	var ret ServiceListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
