package model

import "time"


type UserInfoReq struct {
	Phone int `form:"phone" binding:"required,len=11"`
	VerifyCode string `form:"verify_code" binding:"omitempty"`
	Article string `form:"article" binding:"required,min=10,max=500"`
}

type UserInvolvesInfo struct {
	Id int `json:"id"`
	Phone int `json:"phone"`
	DrawRight int8 `json:"draw_right"`
	Article string `json:"article"`
	Ctime time.Time `json:"ctime"`
}

type UserInvolvesInfosReply struct {
	List []*UserInvolvesInfo
}

const HaveDrawRight = 1