package msg

const (
	userClanFundsNotice = "userClanFundsNotice"
)

type UserClanFunds struct {
	Uid        string `json:"uid"`
	TimeSec    int64  `json:"timeSec"`
	FundsIncr  int64  `json:"fundsIncr"`
	TotalFunds int64  `json:"totalFunds"`
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
