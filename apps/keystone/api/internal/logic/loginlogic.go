package logic

import (
	"context"
	"time"

	"github.com/ScottCho/goweb/apps/keystone/api/internal/svc"
	"github.com/ScottCho/goweb/apps/keystone/api/internal/types"
	"github.com/ScottCho/goweb/apps/keystone/rpc/keystoneclient"
	"github.com/ScottCho/goweb/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.KeystoneRpc.Login(l.ctx, &keystoneclient.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := utils.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
