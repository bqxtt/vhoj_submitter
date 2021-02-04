package oj

import (
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
	panic("implement me")
}

func (P *POJLoginer) Login(account *common.RemoteAccount) ([]*http.Cookie, error) {
	if cookie := cache.CookieCache.Get(util.GetCookieKey(remote_oj.POJ, account)); cookie != nil {
		return cookie, nil
	}

	url := remote_oj.POJInfo.Host + remote_oj.POJInfo.LoginUrl
	method := "POST"

	payload := strings.NewReader("user_id1=bqx&password1=tcg19981108&B1=login&url=%2F")

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	cookies := res.Cookies()
	cache.CookieCache.Put(util.GetCookieKey(remote_oj.POJ, account), cookies, 300)
	//fmt.Printf("cookie: %v", cookies)

	return cookies, nil
}
