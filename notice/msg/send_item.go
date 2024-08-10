package msg

func init() {
	registerBuilder(NoticeTypeSendItem, func() Message {
		return &SendItem{}
	})
}

type SendItem struct {
	ItemId            int     `json:"itemId"`
	ItemCount         int     `json:"itemCount"`
	PayType           int     `json:"payType"`
	PresenterNick     string  `json:"presenterNick"`
	RoomId            int64   `json:"roomId"`
	SendItemComboHits int     `json:"sendItemComboHits"`
	SendNick          string  `json:"sendNick"`
	SendAvatarUrl     string  `json:"sendAvatarUrl"`
	SendTimeStamp     int64   `json:"sendTimeStamp"`
	UserId            string  `json:"userId"`
	TotalGet          float32 `json:"totalGet"`
	TotalPay          float32 `json:"totalPay"`
}

func (msg *SendItem) Notice() string {
	return NoticeTypeSendItem
}
