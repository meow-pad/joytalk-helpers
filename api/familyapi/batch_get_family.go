package familyapi

import "github.com/meow-pad/joytalk-helpers/api"

const BatchGetClanPath = "family/batchGetFamilyInfoById"

type BatchGetFamilyRequest struct {
	// 家族ID列表，最多 50 条。
	FamilyIds []string `json:"family_id_list" validate:"required"`
}

type FamilyGroup struct {
	GroupId    int32  `json:"group_num"`     // 组编号
	GroupName  string `json:"group_name"`    // 组名称
	GroupElder string `json:"group_manager"` // 组 长老
}

type FamilyData struct {
	ID        string        `json:"family_id"`
	PublicID  string        `json:"public_family_id"`
	Name      string        `json:"family_name"`
	Avatar    string        `json:"family_avatar"`
	MemberCnt int32         `json:"member_cnt"`
	Groups    []FamilyGroup `json:"group_list"`
}

type BatchGetFamilyData struct {
	Families []FamilyData `json:"family_info_list"`
	Fails    []string     `json:"fail_list"`
}

type BatchGetClanResponse = api.Response[BatchGetFamilyData]
