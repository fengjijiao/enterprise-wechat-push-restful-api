package httphandler

import (
	"testing"
	"sort"
	"strings"
	"crypto/sha1"
	"fmt"
)

func TestVerifyHttpHandler(t *testing.T) {
	t.Parallel()
	signature := "60b3ef4c26857889adc402f5ad7582b3ed193e94"
	token := "21LcJJVRkFGJ1Wk12FEQQoHDtQA9hrjG"
	timestamp := "aasf"
	nonce := "ccc"
	stringList := []string {token, timestamp, nonce}
	sort.Strings(stringList)
	h := sha1.New()
	t.Logf("source string: %s\n", strings.Join(stringList, ""))
    h.Write([]byte(strings.Join(stringList, "")))
    bs := h.Sum(nil)
	t.Logf("sha1 string: %s\n", string(bs))
	t.Logf("sha1 string [hex]: %x\n", bs)
	if fmt.Sprintf("%x", bs) == signature {
		t.Log("verify pass!")
	}else {
		t.Error("verify failed!")
	}
}