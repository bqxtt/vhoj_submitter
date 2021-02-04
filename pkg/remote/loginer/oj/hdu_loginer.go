package oj

import (
	"bytes"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/cache"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/loginer"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
	"mime/multipart"
	"net/http"
)

var HduLoginer loginer.ILoginer = &HDULoginer{}

type HDULoginer struct {
	loginer.DefaultLoginerImpl
}

func (H *HDULoginer) GetOJInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (H *HDULoginer) Login(account *common.RemoteAccount) ([]*http.Cookie, error) {

	if cookies := cache.CookieCache.Get(util.GetCookieKey(remote_oj.HDU, account)); cookies != nil {
		fmt.Printf("use cache cookies: %v", cookies)
		return cookies, nil
	}
	url := remote_oj.HDUInfo.Host + remote_oj.HDUInfo.LoginUrl
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", account.Username)
	_ = writer.WriteField("userpass", account.Password)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	cookies := res.Cookies()
	//fmt.Printf("cookie: %v\n", cookies)
	_ = cache.CookieCache.Put(util.GetCookieKey(remote_oj.HDU, account), cookies, 300)

	return cookies, nil
}
