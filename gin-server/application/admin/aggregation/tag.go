package aggregation

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/entity"
)

var Tag *tagAggregation

type tagAggregation struct {
	Ent     *entity.Tag
	EntList []entity.Tag
}

func (a *tagAggregation) NewAddTagAggregation(req *in_out.AddTagReq) *entity.Tag {
	r := &tagAggregation{
		Ent: &entity.Tag{
			TagName: req.TagName,
		},
	}
	return r.Ent
}

func (a *tagAggregation) NewUpdateTagAggregation(req *in_out.UpdateTagReq) *entity.Tag {
	r := &tagAggregation{
		Ent: &entity.Tag{
			ID:      req.ID,
			TagName: req.TagName,
		},
	}
	return r.Ent
}
