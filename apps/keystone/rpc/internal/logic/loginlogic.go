package logic

import (
	"context"

	"github.com/ScottCho/goweb/apps/keystone/model"
	"github.com/ScottCho/goweb/apps/keystone/rpc/internal/svc"
	"github.com/ScottCho/goweb/apps/keystone/rpc/pb/keystone"
	"github.com/ScottCho/goweb/pkg/utils"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *keystone.LoginRequest) (*keystone.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &keystone.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
