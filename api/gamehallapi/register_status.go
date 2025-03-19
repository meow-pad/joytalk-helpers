package gamehallapi

import "github.com/meow-pad/joytalk-helpers/api"

const RegisterStatusPath = "gamehall/registerStatus"

type RegisterStatusRequest struct {
	RoomId      int64 `json:"roomId"`
	GameId      int   `json:"gameId"`
	GameStatus  int   `json:"gameStatus"` // 看常量定义
	GameSeatNum int   `json:"gameSeatNum"`
}

type RegisterStatusResponse = api.Response[any]
