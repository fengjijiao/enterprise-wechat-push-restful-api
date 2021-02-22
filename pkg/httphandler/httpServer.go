package httphandler

import (
	"path"
    "net/http"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
)

func Run() error {
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, "/"), defaultHttpHandler)
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, conf.Config.SecurityPrefix, "verify"), verifyHttpHandler)
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, conf.Config.SecurityPrefix, "send"), sendTextCardHttpHandler)
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, conf.Config.SecurityPrefix, "textcard", "send"), sendTextCardHttpHandler)
	http.HandleFunc(path.Join(conf.Config.BaseUrlPath, conf.Config.SecurityPrefix, "text", "send"), sendTextHttpHandler)
    return http.ListenAndServe(conf.Config.HttpServerListen, nil)
}