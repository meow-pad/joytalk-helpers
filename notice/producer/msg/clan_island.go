package msg

const (
	clanIslandNotice = "clanIslandNotice"
)

type ClanIsland struct {
	ClanId      string `json:"clanId"`
	TimeSec     int64  `json:"timeSec"`
	IslandGrade uint32 `json:"islandGrade"`
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
