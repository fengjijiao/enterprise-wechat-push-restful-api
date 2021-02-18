package conf

import (
	//"os"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/logio"
)

type ConfInfo struct {
	WorkDir string `yaml:"work-dir"`
	WechatCorpId string `yaml:"corp-id"`
	WechatCorpSecret string `yaml:"corp-secret"`
	WechatAgentId	int	`yaml:"agent-id"`
	HttpServerListen string `yaml:"http-server-listen"`
	BaseUrlPath string `yaml:"base-url-path"`
	SecurityPrefix string `yaml:"security-prefix"`
	WechatToken string `yaml:"token"`
	WechatEnCodingAesKey string	`yaml:"encoding-aeskey"`
}

func (ci *ConfInfo) setDefaults() {
	if ci.WorkDir == "" {
		ci.WorkDir = "./"
	}
	if ci.WechatCorpId == "" {
		logio.Logger.Fatal("[setDefaults]: WechatCorpId can not be empty!")
	}
	if ci.WechatCorpSecret == "" {
		logio.Logger.Fatal("[setDefaults]: WechatCorpSecret can not be empty!")
	}
	if ci.HttpServerListen == "" {
		ci.HttpServerListen = ":9465"
	}
	if ci.WechatToken == "" {
		logio.Logger.Fatal("[setDefaults]: WechatToken can not be empty!")
	}
	if ci.WechatEnCodingAesKey == "" {
		logio.Logger.Fatal("[setDefaults]: WechatEnCodingAesKey can not be empty!")
	}
}