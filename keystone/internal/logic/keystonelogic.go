package logic

import (
	"context"
	"fmt"

	"github.com/ScottCho/goweb/keystone/internal/svc"
	"github.com/ScottCho/goweb/keystone/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeystoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeystoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeystoneLogic {
	return &KeystoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeystoneLogic) Keystone(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	users, _ := l.svcCtx.UserModel.FindOne(l.ctx, 1)
	fmt.Println(users)
	return
}
