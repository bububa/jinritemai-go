package order

import (
	"github.com/bububa/jinritemai-go/client"
)

/*
order_status入参列表
1   在线支付订单待支付；货到付款订单待确认
2   备货中（只有此状态下，才可发货）
3   已发货
4   已取消：
1.未支付时用户或客服取消订单；
2.用户超时未支付，系统自动取消订单；
3.货到付款订单，用户拒收；
4.订单付款未发货时，触发平台风控，取消订单
5   已完成：1.在线支付订单，商家发货后，用户收货、拒收或者15天无物流更新；2.货到付款订单，用户确认收货
*/
type ListRequest struct {
	Page        int    `json:"page,omitempty"`         // 第几页（第一页为0，最大为99）
	Size        int    `json:"size,omitempty"`         // 每页返回条数，最多支持100条
	OrderStatus *uint  `json:"order_status,omitempty"` // 子订单状态
	StartTime   string `json:"start_time,omitempty"`   // 开始时间
	EndTime     string `json:"end_time,omitempty"`     // 结束时间，必须大于等于开始时间
	OrderBy     string `json:"order_by,omitempty"`     // 1、默认按订单创建时间搜索 2、值为“create_time”：按订单创建时间；值为“update_time”：按订单更新时间
	IsDesc      *bool  `json:"is_desc,omitempty"`      // 订单排序方式：0(is_desc，最近的在前)， 1(asc，最近的在后) 默认为1
}

func (this ListRequest) Method() string {
	return "order/list"
}

func (this ListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"start_time": this.StartTime,
		"end_time":   this.EndTime,
		"order_by":   this.OrderBy,
		"page":       this.Page,
	}
	if this.OrderStatus != nil {
		ret["order_status"] = *this.OrderStatus
	}
	if this.IsDesc != nil && *this.IsDesc {
		ret["is_desc"] = 1
	}
	if this.Size > 0 {
		ret["size"] = this.Size
	}
	return ret
}

type ListResponse struct {
	Total int           `json:"total,omitempty"` // 满足查询条件的订单总数（如果入参传了order_status，则total为子订单总数；否则为父订单总数）
	Count int           `json:"count,omitempty"` // //当前返回里list结构中的父订单数
	List  []OrderDetail `json:"list,omitempty"`
}

// 获取订单列表
// 支持按照子订单状态和订单的创建时间、更新时间来检索订单，获取订单列表
/*
Tips：

此接口做了限流，每个access_token每分钟1000次上限
此接口，page最多返回100页，一页最多100单，故同样的请求，最多能拉到1万单
通过该接口获取到的父订单，平台会在order_id后面会加上字母A做标识（这里和从店铺后台导出的订单有区别，从店铺后台导出的订单末尾没有字母A，但其实际和通过openapi获取的父订单是同一个）
如搜索时指定order_status，则搜索维度为子订单，即按照子订单状态搜索，total也为该状态子订单数量。返回结果的json结构不变，仍为父子订单结构，因此当有父订单包含多子订单时，返回的结构中父订单可能会有重复!
由于订单同步可能会有延迟，建议在使用open接口拉取订单时，使用时间窗口移动的方式拉单，并做兜底处理。比如，在6:00去拉5:30-5:40的数据（20分钟后拉），6:05去拉5:35-5:45的数据，第二天的6:00时再拉取一次5:30-5:40的数据（兜底补拉）
*/
func GetList(clt *client.Client, req *ListRequest) (*ListResponse, error) {
	var ret ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
