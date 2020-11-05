package order

import (
	"github.com/bububa/jinritemai-go/client"
)

type ReplyServiceRequest struct {
	ID    uint64 `json:"id,omitempty"`    // 服务请求列表中获取的id
	Reply string `json:"reply,omitempty"` // 回复内容
}

func (this ReplyServiceRequest) Method() string {
	return "order/replyService"
}

func (this ReplyServiceRequest) Params() map[string]interface{} {
	return map[string]interface{}{
		"id":    this.ID,
		"reply": this.Reply,
	}
}

// 回复服务请求
// 回复客服向店铺发起的服务请求
func ReplyService(clt *client.Client, serviceID uint64, reply string) error {
	req := &ReplyServiceRequest{
		ID:    serviceID,
		Reply: reply,
	}
	return clt.Do(req, nil)
}
