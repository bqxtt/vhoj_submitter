package oj

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/cache"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/loginer"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
	"net/http"
	"strings"
)

var PojLoginer loginer.ILoginer = &POJLoginer{}

type POJLoginer struct {
	loginer.DefaultLoginerImpl
}

func (P *POJLoginer) GetOJInfo() *remote_oj.RemoteOJInfo {
	return remote_oj.POJInfo
}

func (P *POJLoginer) Login(req *http.Request, account *common.RemoteAccount) (*http.Request, error) {
	var cookies = cache.CookieCache.Get(util.GetCookieKey(remote_oj.POJ, account))
	if cookies == nil {
		url := remote_oj.POJInfo.Host + remote_oj.POJInfo.LoginUrl
		method := "POST"

		payload := strings.NewReader(fmt.Sprintf("user_id1=%v&password1=%v&B1=login&url=%v", account.Username, account.Password, "\\"))

		r, err := http.NewRequest(method, url, payload)

		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		cookies = res.Cookies()
		_ = cache.CookieCache.Put(util.GetCookieKey(remote_oj.POJ, account), cookies, 300)
	}
	for _, cookie := range cookies {
		req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)
	}
	return req, nil
}
