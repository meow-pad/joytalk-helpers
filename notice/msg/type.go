package msg

const (
	NoticeTypeGetCharge    = "getChargeNotice"
	NoticeTypeGetFollow    = "getFollowNotice"
	NoticeTypeGetHeartbeat = "getUserHeartBeatNotice"
	NoticeTypeGetMessage   = "getMessageNotice"
	NoticeTypeSendItem     = "getSendItemNotice"
)

var noticeTypeTopics = map[string]string{
	"getMessageNotice":       "getMessageNotice",
	"getSendItemNotice":      "getSendItemNotice",
	"getFollowNotice":        "getFollowNotice",
	"getLoginNotice":         "getLoginNotice",
	"getChargeNotice":        "getChargeNotice",
	"getUserHeartBeatNotice": "getUserHeartBeatNotice",
}

func GetNoticeTypeTopic(noticeType string) (string, bool) {
	topic, exist := noticeTypeTopics[noticeType]
	return topic, exist
}
