package userapi

import "github.com/meow-pad/joytalk-helpers/api"

const BatchGetUserPath = "user/batchGetUserInfoById"

type BatchGetUserRequest struct {
	api.BaseRequest
	// 用户ID列表，最多 50 条。
	UserIds []string `json:"user_id_list" validate:"required"`
}

type UserData struct {
	Uid    string `json:"user_id"`
	Nick   string `json:"nick"`
	Avatar string `json:"avatar"`
	ClanId string `json:"family_id"`
}

type BatchGetUserData struct {
	Users []UserData `json:"user_info_list"`
	Fails []string   `json:"fail_list"`
}

type BatchGetUserResponse = api.Response[BatchGetUserData]
