// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserRes) (resp *types.UserResq, err error) {
	// todo: add your logic here and delete this line
	//resp = &types.UserResq{
	//	Id:    "666",
	//	Name:  "666",
	//	Phone: "6666666666",
	//}
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
}
