package logic

import (
	"context"

	"zero/user/rpc/internal/svc"
	"zero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.UserInfo) (*user.ResponseMsg, error) {
	// todo: add your logic here and delete this line

	return &user.ResponseMsg{}, nil
}
