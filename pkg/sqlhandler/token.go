package sqlhandler

import (
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/commonio"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
	"path/filepath"
	"time"
	"errors"
)

const (
	TokenFileName string = "token.cache"
)

func getTokenFilePath() string {
	return filepath.Join(conf.Config.WorkDir, TokenFileName)
}

func UpdateToken(token string) error {
	return commonio.WriteToFile(getTokenFilePath(), []byte(token))
}

func GetToken() (string, error) {
	timestamp, err := commonio.GetFileModifyTime(getTokenFilePath())
	if err != nil {
		return "", err
	}
	if time.Now().Unix() - timestamp > 6000 {
		return "", errors.New("token was expired.")
	}
	dat, err := commonio.ReadFile(getTokenFilePath())
	if err != nil {
		return "", err
	}
	return string(dat), nil
}