package logic

import (
	"context"

	"docker-go-zero/demo/internal/svc"
	"docker-go-zero/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnvLogic {
	return &EnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnvLogic) Env() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = "success"
	resp.Values = l.svcCtx.Config.AppEnv

	return
}
