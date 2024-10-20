package payapi

import "github.com/meow-pad/joytalk-helpers/api"

const BalancePath = "gamehall/balance"

type BalanceRequest struct {
	JoytalkId int64  `json:"joytalkId"`
	UserToken string `json:"userToken"`
}

type BalanceData struct {
	Balance string `json:"balance"` // 余额值
}

type BalanceResponse = api.Response[BalanceData]
