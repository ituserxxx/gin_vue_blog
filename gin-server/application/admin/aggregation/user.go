package aggregation

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/entity"
	"gin-server/utils"
)

var User *userAggregation

type userAggregation struct {
	Ent     *entity.User
	EntList []entity.User
}

func (ua *userAggregation) NewAddAdminAggregation(req *in_out.AddAdminReq) *entity.User {
	r := &userAggregation{
		Ent: &entity.User{
			Username: req.Username,
			Password: utils.EncodeMD5(req.Password),
		},
	}
	return r.Ent
}

func (ua *userAggregation) NewDelAdminAggregation(req *in_out.AdminUserIDReq) *entity.User {
	r := &userAggregation{
		Ent: &entity.User{
			Id: req.UID,
		},
	}
	return r.Ent
}
func (ua *userAggregation) NewAdminLoginAggregation(req *in_out.LoginReq) *entity.User {
	r := &userAggregation{
		Ent: &entity.User{
			Username: req.Username,
			Password: req.Password,
		},
	}
	return r.Ent
}
func (ua *userAggregation) NewGetAdminInfoAggregation(uid int) *entity.User {
	r := &userAggregation{
		Ent: &entity.User{
			Id: uid,
		},
	}
	return r.Ent
}
