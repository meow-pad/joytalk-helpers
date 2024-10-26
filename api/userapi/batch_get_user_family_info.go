package userapi

import (
	"github.com/meow-pad/joytalk-helpers/api"
	"github.com/meow-pad/joytalk-helpers/api/familyapi"
)

const BatchGetUserFamilyInfoPath = "family/batchGetUserFamilyInfo"

type BatchGetUserFamilyInfoRequest struct {
	// 用户ID列表，最多 50 条。
	UserIds []string `json:"user_id_list" validate:"required"`
}

type UserFamilyData struct {
	Uid            string `json:"user_id"`
	FamilyAvatar   string `json:"family_avatar"`
	FamilyId       string `json:"family_id"`
	PublicFamilyId string `json:"public_family_id"`
	FamilyName     string `json:"family_name"`
	MemberCnt      int32  `json:"member_cnt"`
	MemberType     int32  `json:"member_type"`
	GroupId        int32  `json:"group_num"`
}

func (data *UserFamilyData) ToFamilyData() *familyapi.FamilyData {
	return &familyapi.FamilyData{
		Avatar:    data.FamilyAvatar,
		ID:        data.FamilyId,
		PublicID:  data.PublicFamilyId,
		Name:      data.FamilyName,
		MemberCnt: data.MemberCnt,
	}
}

type BatchGetUserFamilyInfoData struct {
	UserFamilies []UserFamilyData `json:"family_info_list"`
	Fails        []string         `json:"fail_list"`
}

type BatchGetUserFamilyInfoResponse = api.Response[BatchGetUserData]
