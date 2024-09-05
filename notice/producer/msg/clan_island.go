package msg

const (
	clanIslandNotice = "clanIslandNotice"
)

type ClanIsland struct {
	ClanId    string `json:"clanId"`
	TimeMs    int64  `json:"timeMs"`
	StarIncr  int32  `json:"starIncr"`
	TotalStar int32  `json:"totalStar"`
}

func (clanIsland *ClanIsland) Key() string {
	return clanIsland.ClanId
}

func (clanIsland *ClanIsland) Topic() string {
	return topicPrefix + clanIslandNotice
}

func (clanIsland *ClanIsland) Notice() string {
	return clanIslandNotice
}

func (clanIsland *ClanIsland) Data() any {
	return clanIsland
}

func (clanIsland *ClanIsland) ToNoticeBytes() ([]byte, error) {
	return toNoticeBytes(clanIsland)
}
