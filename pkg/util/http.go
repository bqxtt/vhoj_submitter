package util

import (
	"fmt"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"io/ioutil"
	"net/http"
)

func Download(url string, method string, auth *LoginFunc) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	if auth != nil {
		req, err = auth.Login(req, auth.Account)
		if err != nil {
			return "", fmt.Errorf("login failed: %v", err)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}

type LoginFunc struct {
	Account *common.RemoteAccount
	Login   func(*http.Request, *common.RemoteAccount) (*http.Request, error)
}
