package payapi

import "github.com/meow-pad/joytalk-helpers/api"

const OrderRewardPath = "gamehall/order/reward"

type OrderRewardRequest struct {
	JoytalkId int64  `json:"joytalkId"`
	Amount    int    `json:"amount"`
	OrderId   string `json:"orderId"`
}

type OrderRewardResponse = api.Response[any]
