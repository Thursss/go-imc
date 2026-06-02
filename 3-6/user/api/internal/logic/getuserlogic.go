// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserRes) (resp *types.UserResq, err error) {
	user, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserRes{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}
	return &types.UserResq{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil

	return
}
