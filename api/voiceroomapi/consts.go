package voiceroomapi

const (
	SeatUidNone = int64(0)

	StatusClosed    = 0 // 关闭
	StatusOpen      = 1 // 开启
	StatusBeingLess = 2 // 未开播过

	VoiceRoomSeatStatusEmpty  = 0 // 空位
	VoiceRoomSeatStatusLocked = 1 // 被锁定
	VoiceRoomSeatStatusSeated = 2 // 已入座

	VoiceRoomMicStatusNormal   = 0 // 可发言
	VoiceRoomMicStatusBanned   = 1 // 被禁言
	VoiceRoomMicStatusBlanking = 2 // 自己闭麦

	VoiceRoomSeatIndexStart = 1     // 房间位置起始编号
	VoiceRoomSeatIndexHost  = 10000 // 主持位
)
