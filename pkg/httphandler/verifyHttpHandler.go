package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/enterprise-wechat-push-restful-api/pkg/conf"
	"sort"
	"strings"
	"crypto/sha1"
)

func verifyHttpHandler(w http.ResponseWriter, req *http.Request) {
	signature := req.URL.Query().Get("msg_signature")
	timestamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	echostr := req.URL.Query().Get("echostr")
	if len(signature) == 0 || len(timestamp) == 0 || len(nonce) == 0 || len(echostr) == 0 {
		fmt.Fprintf(w, "verification failed\n")
		return
	}
	stringList := []string {conf.Config.WechatToken, timestamp, nonce}
	sort.Strings(stringList)
	h := sha1.New()
    h.Write([]byte(strings.Join(stringList, "")))
    bs := h.Sum(nil)
	if fmt.Sprintf("%x", bs) == signature {
		fmt.Fprintf(w, echostr)
		return
	}
	fmt.Fprintf(w, "verification failed\n")
}