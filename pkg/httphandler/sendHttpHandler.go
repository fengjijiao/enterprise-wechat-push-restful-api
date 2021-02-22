package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/logio"
	"go.uber.org/zap"
	//"path"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/sqlhandler"
	"encoding/json"
	"github.com/imroc/req"
)

type SendTextCardInfo struct {
	ToUser   string `json:"touser"`
	ToParty  string `json:"toparty"`
	ToTag    string `json:"totag"`
	MsgType  string `json:"msgtype"`
	AgentId  int    `json:"agentid"`
	TextCard struct {
		Title       string `json:"title"` //标题，不超过128个字节，超过会自动截断
		Description string `json:"description"` //描述，不超过512个字节，超过会自动截断
		URL         string `json:"url"`
		BtnTxt      string `json:"btntxt"`
	} `json:"textcard"`
	EnableIDTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

type SendTextInfo struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIDTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

type ErrorInfo struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func sendTextCardHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hr.ParseForm()
	title := hr.FormValue("title")
	body := hr.FormValue("body")
	touser := hr.FormValue("touser")
	toparty := hr.FormValue("toparty")
	totag := hr.FormValue("totag")
	if len(title) <= 0 || len(body) <= 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	var res ErrorInfo
	var sendInfo SendTextCardInfo
	if len(touser) > 0 {
		sendInfo.ToUser = touser
	}else if len(toparty) > 0 {
		sendInfo.ToParty = toparty
	}else if len(totag) > 0 {
		sendInfo.ToTag = totag
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	sendInfo.MsgType = "textcard"
	sendInfo.TextCard.Title = title
	sendInfo.TextCard.Description = body
	sendInfo.TextCard.URL = "https://www.fengjijiao.cn/?ref=enterprise-wechat"
	sendInfo.TextCard.BtnTxt = "详情"
	sendInfo.AgentId = conf.Config.WechatAgentId
	sendInfo.EnableIDTrans = 0
	sendInfo.EnableDuplicateCheck = 0
	sendInfo.DuplicateCheckInterval = 900
	param, err := json.Marshal(&sendInfo)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	token, err := sqlhandler.GetToken()
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r, err := req.Post(fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s`, token), param)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r.ToJSON(&res)
	if res.ErrCode == 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{0, "send message success!"})
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed!"+res.ErrMsg})
	}
}

func sendTextHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hr.ParseForm()
	context := hr.FormValue("context")
	touser := hr.FormValue("touser")
	toparty := hr.FormValue("toparty")
	totag := hr.FormValue("totag")
	if len(context) <= 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	var res ErrorInfo
	var sendInfo SendTextInfo
	if len(touser) > 0 {
		sendInfo.ToUser = touser
	}else if len(toparty) > 0 {
		sendInfo.ToParty = toparty
	}else if len(totag) > 0 {
		sendInfo.ToTag = totag
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	sendInfo.MsgType = "text"
	sendInfo.AgentId = conf.Config.WechatAgentId
	sendInfo.Text.Content = context
	sendInfo.Safe = 0
	sendInfo.EnableIDTrans = 0
	sendInfo.EnableDuplicateCheck = 0
	sendInfo.DuplicateCheckInterval = 900
	param, err := json.Marshal(&sendInfo)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	token, err := sqlhandler.GetToken()
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r, err := req.Post(fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s`, token), param)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r.ToJSON(&res)
	if res.ErrCode == 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{0, "send message success!"})
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed!"+res.ErrMsg})
	}
}