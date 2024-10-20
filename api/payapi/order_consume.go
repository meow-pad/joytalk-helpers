package payapi

import "github.com/meow-pad/joytalk-helpers/api"

const OrderConsumerPath = "gamehall/order/consume"

type OrderConsumeRequest struct {
	JoytalkId int64  `json:"joytalkId"`
	Amount    int    `json:"amount"`
	UserToken string `json:"userToken"`
	OrderId   string `json:"orderId"`
}

type OrderConsumeResponse = api.Response[any]
