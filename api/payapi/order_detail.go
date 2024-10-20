package payapi

import "github.com/meow-pad/joytalk-helpers/api"

const OrderDetailPath = "gamehall/order/detail"

const (
	OrderTypeUserConsume = 1 // 用户消费订单
	OrderTypeReward      = 2 // 发奖订单

	OrderStatusNotExist = 1 // 不存在
	OrderStatusSuccess  = 2 // 支付成功
)

type OrderDetailRequest struct {
	JoytalkId int64  `json:"joytalkId"`
	Amount    int    `json:"amount"`
	OrderId   string `json:"orderId"`
}

type OrderDetailData struct {
	OrderStatus int `json:"orderStatus"` // 1 不存在； 2支付成功
	Amount      int `json:"amount"`      // 订单金额
	OrderType   int `json:"orderType"`   // 1 用户消费订单; 2 发奖订单
}

type OrderDetailResponse = api.Response[OrderDetailData]
