package voiceroomapi

import "github.com/meow-pad/joytalk-helpers/api"

const GetGameRoomInfo = "voiceroom/getGameRoomInfo"

type GetGameRoomInfoRequest struct {
	RoomId int64 `json:"roomId" validate:"required"`
	GameId int64 `json:"gameId" validate:"required"`
}

type GetGameRoomInfoData struct {
	SessionId string     `json:"sessionId"`
	Status    int        `json:"status"` // 查看常量定义
	Seat      []SeatData `json:"seat"`
	HostSeat  SeatData   `json:"hostSeat"`
}

type GetGameRoomInfoResponse = api.Response[GetGameRoomInfoData]
