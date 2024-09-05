package msg

const (
	userClanFundsNotice = "userClanFundsNotice"
)

type UserClanFunds struct {
	Uid        string `json:"uid"`
	TimeMs     int64  `json:"timeMs"`
	FundsIncr  int32  `json:"fundsIncr"`
	TotalFunds int32  `json:"totalFunds"`
}

func (userClanFunds *UserClanFunds) Key() string {
	return userClanFunds.Uid
}

func (userClanFunds *UserClanFunds) Topic() string {
	return topicPrefix + userClanFundsNotice
}

func (userClanFunds *UserClanFunds) Notice() string {
	return userClanFundsNotice
}

func (userClanFunds *UserClanFunds) Data() any {
	return userClanFunds
}

func (userClanFunds *UserClanFunds) ToNoticeBytes() ([]byte, error) {
	return toNoticeBytes(userClanFunds)
}
