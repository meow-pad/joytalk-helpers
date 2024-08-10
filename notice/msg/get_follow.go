package msg

func init() {
	registerBuilder(NoticeTypeGetFollow, func() Message {
		return &GetFollow{}
	})
}

type GetFollow struct {
	RoomId        string `json:"roomId"`
	PresenterNick string `json:"presenterNick"`
	UserNick      string `json:"userNick"`
	UserAvatarUrl string `json:"userAvatarUrl"`
	UserId        int64  `json:"userId"`
	Gender        int32  `json:"gender"`
}

func (msg *GetFollow) Notice() string {
	return NoticeTypeGetFollow
}
