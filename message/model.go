package message

import (
	"encoding/json"

	"github.com/bububa/jinritemai-go/message/refund"
	"github.com/bububa/jinritemai-go/message/trade"
)

var SuccessResponse = Response{Code: 0, Msg: "success"}

type Response struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
}

type Message struct {
	Tag   string `json:"tag"`
	MsgID string `json:"msg_id"`
	Data  string `json:"data"`
}

type MessageData interface {
	Tag() string
}

func (this Message) DecodeData() MessageData {
	var ret MessageData
	switch this.Tag {
	case "100":
		ret = trade.TradeCreate{}
	case "101":
		ret = trade.TradePaid{}
	case "102":
		ret = trade.TradeSellerShip{}
	case "103":
		ret = trade.TradeSuccess{}
	case "104":
		ret = trade.TradeLogisticsChanged{}
	case "105":
		ret = trade.TradeAddressChanged{}
	case "106":
		ret = trade.TradeCanceled{}
	case "200":
		ret = refund.RefundCreate{}
	case "201":
		ret = refund.RefundAggree{}
	case "202":
		ret = refund.ReturnApplyAgree{}
	case "203":
		ret = refund.BuyerReturnGoods{}
	case "204":
		ret = refund.RefundRefused{}
	case "205":
		ret = refund.ReturnApplyRefused{}
	case "207":
		ret = refund.RefundSuccess{}
	case "208":
		ret = refund.RefundClosed{}
	}
	if ret != nil {
		json.Unmarshal([]byte(this.Data), &ret)
	}
	return ret
}

func (this Message) Execute(delegator *Delegator) error {
	fn := delegator.Get(this.Tag)
	if fn == nil {
		return nil
	}
	return fn(this)
}
