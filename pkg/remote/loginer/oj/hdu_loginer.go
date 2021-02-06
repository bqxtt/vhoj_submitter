package oj

import (
	"bytes"
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

func (H *HDULoginer) Login(req *http.Request, account *common.RemoteAccount) (*http.Request, error) {

	var cookies = cache.CookieCache.Get(util.GetCookieKey(remote_oj.HDU, account))
	if cookies == nil {
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

		r, err := http.NewRequest(method, url, payload)

		if err != nil {
			return nil, err
		}
		r.Header.Set("Content-Type", writer.FormDataContentType())
		res, err := http.DefaultTransport.RoundTrip(r)

		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		cookies = res.Cookies()
		_ = cache.CookieCache.Put(util.GetCookieKey(remote_oj.HDU, account), cookies, 300)
	}
	for _, cookie := range cookies {
		req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)
	}
	return req, nil
}
