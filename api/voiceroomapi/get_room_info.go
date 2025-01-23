package voiceroomapi

import "github.com/meow-pad/joytalk-helpers/api"

const GetRoomInfo = "voiceroom/getRoomInfo"

type GetRoomInfoRequest struct {
	RoomId int64 `json:"roomId" validate:"required"`
}

type SeatData struct {
	Uid        int64 `json:"uid"`
	Index      int   `json:"index"`
	SeatStatus int   `json:"seaStatus"` // 查看常量定义
	MicStatus  int   `json:"micStatus"` // 查看常量定义
}

type GetRoomInfoData struct {
	SessionId string     `json:"sessionId"`
	Status    int        `json:"status"` // 查看常量定义
	Seat      []SeatData `json:"seat"`
	HostSeat  SeatData   `json:"hostSeat"`
}

type GetRoomInfoResponse = api.Response[GetRoomInfoData]
