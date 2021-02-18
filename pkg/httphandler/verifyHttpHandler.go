package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/wxbizmsgcrypt"
)

func verifyHttpHandler(w http.ResponseWriter, req *http.Request) {
	msg_signature := req.URL.Query().Get("msg_signature")
	timestamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	echostr := req.URL.Query().Get("echostr")
	if len(msg_signature) == 0 || len(timestamp) == 0 || len(nonce) == 0 || len(echostr) == 0 {
		fmt.Fprintf(w, "verification failed\n")
		return
	}
	wxBizMsgCrypt := wxbizmsgcrypt.NewWXBizMsgCrypt(conf.Config.WechatToken, conf.Config.WechatEnCodingAesKey, conf.Config.WechatCorpId, wxbizmsgcrypt.XmlType)
	res, err := wxBizMsgCrypt.VerifyURL(msg_signature, timestamp, nonce, echostr)
	if err != nil {
		fmt.Fprintf(w, "verification failed\n")
		return
	}
	w.Write(res)
}