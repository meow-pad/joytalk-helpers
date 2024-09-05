package msg

func init() {
	registerBuilder(NoticeTypeGetCharge, func() Message {
		return &GetCharge{}
	})
}

type GetCharge struct {
	Amount     float64 `json:"amount"`
	PayType    int     `json:"payType"`
	OrderId    string  `json:"orderId"`
	TsMs       int64   `json:"tsMs"`
	ChargeType int     `json:"chargeType"`
	UserId     int64   `json:"userId"`
}

func (msg *GetCharge) Notice() string {
	return NoticeTypeGetCharge
}
