package logic

import (
	"context"
	"time"

	"docker-go-zero/demo/internal/svc"
	"docker-go-zero/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelayLogic {
	return &DelayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelayLogic) Delay() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	time.Sleep(time.Second * 10)

	var data map[string]interface{}
	data = make(map[string]interface{})

	data["name"] = "delay"
	data["time"] = 10

	resp = new(types.Response)
	resp.Message = "success"
	resp.Values = data

	return
}
