package msg

func init() {
	registerBuilder(NoticeTypeGetMessage, func() Message {
		return &GetMsg{}
	})
}

type GetMsg struct {
	Content         string `json:"content"`
	FanLevel        int    `json:"fanLevel"`
	Gender          int    `json:"gender"`
	MsgType         int    `json:"msgType"`
	RoomId          int    `json:"roomId"` // 房间号
	RoyalLevel      int    `json:"royalLevel"`
	SenderNick      string `json:"senderNick"`
	SenderAvatarUrl string `json:"senderAvatarUrl"`
	SendTimeStamp   int64  `json:"sendTimeStamp"`
	VipLevel        int    `json:"vipLevel"`
	UserId          int64  `json:"userId"`
}

func (msg *GetMsg) Notice() string {
	return NoticeTypeGetMessage
}
