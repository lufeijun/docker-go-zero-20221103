package logic

import (
	"context"
	"net"
	"os"

	"docker-go-zero/demo/internal/svc"
	"docker-go-zero/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OsLogic {
	return &OsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OsLogic) Os() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	var data map[string]interface{}
	data = make(map[string]interface{})

	data["name"] = getHostName()
	data["ip"] = getLocalIp()

	resp = new(types.Response)
	resp.Message = "success"
	resp.Values = data

	return
}

func getHostName() (name string) {
	name, err := os.Hostname()

	if err != nil {
		name = "none"
	}

	return
}

func getLocalIp() (ip string) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ip = "none"
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}

	return
}
