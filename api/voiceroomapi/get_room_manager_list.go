package voiceroomapi

import "github.com/meow-pad/joytalk-helpers/api"

const GetRoomManagerList = "voiceroom/getRoomManagerList"

type GetRoomManagerListRequest struct {
	RoomId int64 `json:"roomId" validate:"required"`
}

type GetRoomManagerListData struct {
	ManagerList []int64 `json:"managerList"`
}

type GetRoomManagerListResponse = api.Response[GetRoomManagerListData]
