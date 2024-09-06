package msg

const (
	clanStarNotice = "clanStarNotice"
)

type ClanStar struct {
	ClanId    string `json:"clanId"`
	TimeSec   int64  `json:"timeSec"`
	StarIncr  int32  `json:"starIncr"`
	TotalStar int32  `json:"totalStar"`
}

func (clanStar *ClanStar) Key() string {
	return clanStar.ClanId
}

func (clanStar *ClanStar) Topic() string {
	return topicPrefix + clanStarNotice
}

func (clanStar *ClanStar) Notice() string {
	return clanStarNotice
}

func (clanStar *ClanStar) Data() any {
	return clanStar
}

func (clanStar *ClanStar) ToNoticeBytes() ([]byte, error) {
	return toNoticeBytes(clanStar)
}
