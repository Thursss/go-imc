package logic

import (
	"context"

	"user/models"
	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:       in.Id,
		Username: in.Name,
		Mobile:   in.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &user.CreateUserResp{
		Id: in.Id,
	}, nil
}
