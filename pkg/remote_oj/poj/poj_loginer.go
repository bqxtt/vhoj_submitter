package poj

import (
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/loginer"
	"io/ioutil"
	"net/http"
	"strings"
)

type POJLoginer struct {
	loginer.DefaultLoginerImpl
}

func (P *POJLoginer) GetOJInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (P *POJLoginer) Login() error {
	url := "http://poj.org/login"
	method := "POST"

	payload := strings.NewReader("user_id1=bqx&password1=tcg19981108&B1=login&url=%2F")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	//fmt.Printf("header: %v", res.Header)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
