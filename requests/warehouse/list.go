package warehouse

import (
	"github.com/bububa/jinritemai-go/client"
)

/*
type值对应表
2   6: 退货申请待商家处理
30: 仅退款申请待商家处理
3   11: 买家已退货，待商家收货
4   8: 商家拒绝退货申请，待客服仲裁
13: 商家拒绝收货，待客服仲裁
34: 商家拒绝仅退款申请，待客服仲裁
5   22: 在线支付订单，退货退款成功
24: 货到付款订单，退货退款成功
39: 仅退款成功
7   7 & 10: 待买家填写退货物流
*/
type ListRequest struct {
	OutWarehouseID uint64         `json:"out_warehouse_id,omitempty"` // 外部仓库ID
	WarehouseName  string         `json:"warehouse_name,omitempty"`   // 仓库名称
	Page           int            `json:"page,omitempty"`             // 第几页（第一页为0，最大为99）
	Size           int            `json:"size,omitempty"`             // 每页返回条数，最多支持100条
	OrderBy        string         `json:"order_by,omitempty"`         // 排序方式，可选create_time、update_time
	Rank           string         `json:"rank,omitempty"`             // 顺序，可选desc、asc，与order_by同时生效
	Addr           *WarehouseAddr `json:"addr,omitempty"`
}

func (this ListRequest) Method() string {
	return "warehouse/list"
}

func (this ListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"size": this.Size,
		"page": this.Page,
	}
	if this.OutWarehouseID > 0 {
		ret["out_warehouse_id"] = this.OutWarehouseID
	}
	if this.WarehouseName != "" {
		ret["warehouse_name"] = this.WarehouseName
	}
	if this.OrderBy != "" {
		ret["order_by"] = this.OrderBy
	}
	if this.Rank != "" {
		ret["rank"] = this.Rank
	}
	if this.Addr != nil {
		ret["addr"] = this.Addr
	}
	return ret
}

type ListResponse struct {
	Total int             `json:"total,omitempty"`      // 总数
	List  []WarehouseInfo `json:"warehouses,omitempty"` // 区域仓数组
}

// 批量查询区域仓
func GetList(clt *client.Client, req *ListRequest) (*ListResponse, error) {
	var ret ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
