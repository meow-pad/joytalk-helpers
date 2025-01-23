package gamehallapi

import "github.com/meow-pad/joytalk-helpers/api"

const GetGameStatus = "gamehall/getStatus"

type GetGameStatusRequest struct {
	RoomId string `json:"roomId" validate:"required"`
}

type GameHallStatus struct {
	GameStatus         int     `json:"gameStatus"`         // 这里只会返回进行中的游戏
	GamePanelStatus    int     `json:"gamePanelStatus"`    // 查看常量定义，是否需要展示当前游戏面板
	GameParticipantNum int     `json:"gameParticipantNum"` // 加入游戏的人数（上游戏麦玩家数）
	GameParticipantUid []int64 `json:"gameParticipantUid"` // 加入游戏的人（上游戏麦玩家）
}

type GetGameStatusData struct {
	Status []GameHallStatus `json:"status"`
}

type GetGameStatusResponse = api.Response[GetGameStatusData]
