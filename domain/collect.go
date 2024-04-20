package domain

import collectv1 "github.com/MuxiKeStack/be-api/gen/proto/collect/v1"

type Collection struct {
	Id    int64
	Uid   int64
	Biz   collectv1.Biz
	BizId int64
}
