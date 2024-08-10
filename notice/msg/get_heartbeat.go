package msg

func init() {
	registerBuilder(NoticeTypeGetHeartbeat, func() Message {
		return &GetHeartbeat{}
	})
}

type GetHeartbeat struct {
	UserId    int64 `json:"userId"`
	Timestamp int64 `json:"timestamp"`
}

func (msg *GetHeartbeat) Notice() string {
	return NoticeTypeGetHeartbeat
}
