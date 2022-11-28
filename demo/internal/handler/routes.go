// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"docker-go-zero/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: HelloHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: VersionHandler(serverCtx),
			},
		},
	)
}
