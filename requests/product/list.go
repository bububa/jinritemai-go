package product

import (
	"github.com/bububa/jinritemai-go/client"
)

type ListRequest struct {
	Page        int   `json:"page,omitempty"`         // 第几页（第一页为0，最大为99）
	Size        int   `json:"size,omitempty"`         // 每页返回条数，最多支持100条
	Status      *uint `json:"status,omitempty"`       // 指定状态返回商品列表：0上架 1下架
	CheckStatus uint  `json:"check_status,omitempty"` // 指定审核状态返回商品列表：1未提审 2审核中 3审核通过 4审核驳回 5封禁
}

func (this ListRequest) Method() string {
	return "product/list"
}

func (this ListRequest) Params() map[string]interface{} {
	ret := map[string]interface{}{
		"page": this.Page,
		"size": this.Size,
	}
	if this.Status != nil {
		ret["status"] = *this.Status
	}
	if this.CheckStatus >= 1 && this.CheckStatus <= 5 {
		ret["check_status"] = this.CheckStatus
	}
	return ret
}

type ListResponse struct {
	Total       int             `json:"all,omitempty"`       // 商品总数
	TotalPages  int             `json:"all_pages,omitempty"` // 已当前size所得的分页数
	Count       int             `json:"count,omitempty"`     // 当前条件data返回结果数量
	CurrentPage int             `json:"current_page,omitempty"`
	List        []ProductDetail `json:"data,omitempty"`
}

func GetList(clt *client.Client, req *ListRequest) (*ListResponse, error) {
	var ret ListResponse
	err := clt.Do(req, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
