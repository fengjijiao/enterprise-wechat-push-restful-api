package main

import (
	//"fmt"
	//"os"
	"flag"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/commonio"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/logio"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/sqlhandler"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/httphandler"
	"go.uber.org/zap"
	"github.com/jasonlvhit/gocron"
)

var (
	configFilePath string
	flagQuiet bool
	Closed chan struct{}
)

func init() {	
	logio.Init()
	
	flag.StringVar(&configFilePath, "c", "config.yaml", "configure file location.")
	flag.BoolVar(&flagQuiet, "quiet", false, "quiet for log print.")
    flag.Parse()
	
	if(flagQuiet) {
		logio.Cfg.Level.SetLevel(zap.ErrorLevel)
	}
	
	if !commonio.IsFileExists(configFilePath) {
		logio.Logger.Fatal("configure file not found!")
	}
	
	dat, err := commonio.ReadFile(configFilePath)
	if err != nil {
		logio.Logger.Fatal("read configure file fail!", zap.Error(err))
	}
	
	err = conf.Load(dat)
	if err != nil {
		logio.Logger.Fatal("parse configure file fail!", zap.Error(err))
	}
}

func main() {
	Closed = make(chan struct{})
	gocron.Every(1).Hour().Do(sqlhandler.UpdateTokenAuto)
	go gocron.Start()
	go httphandler.Run()
	sqlhandler.UpdateTokenAuto()
	for range Closed {
        close(Closed)
    }
}