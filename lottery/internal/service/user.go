package service

import (
	"context"

	"app/lottery/internal/model"
)

func (s *Svc) Participate (c context.Context, req *model.UserInfoReq) (isOk bool, err error){
	_, err = s.dao.SaveUsers(c, req.Phone, req.Article)
	if err != nil {
		s.dao.Logger.Printf("s.Participate err(%v)", err)
		return
	}
	isOk = true
	return
}




func (s *Svc) GetAllUsers (c context.Context) (data *model.UserInvolvesInfosReply, err error){
	data = new(model.UserInvolvesInfosReply)
	data.List, err = s.dao.FetchUsers(c)
	if err != nil {
		s.dao.Logger.Printf("s.GetAllUsers err(%v)", err)
	}
	return
}