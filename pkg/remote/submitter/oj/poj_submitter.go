package oj

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/loginer/oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/submitter"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
	"io/ioutil"
	"net/http"
	"strings"
)

var PojSubmitter submitter.ISubmitter = &POJSubmitter{}

type POJSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (P *POJSubmitter) GetOjInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (P *POJSubmitter) NeedLogin() bool {
	return true
}

func (P *POJSubmitter) GetMaxRunId(info *common.SubmissionInfo, account *common.RemoteAccount) (string, error) {
	url := remote_oj.POJInfo.Host + fmt.Sprintf(remote_oj.POJInfo.StatusUrl, account.Username, info.RemoteProblemId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	runId, err := util.ParseHtmlReg(`<tr align=center><td>(\d+)`, string(body))
	if err != nil {
		return "", nil
	}
	return runId, nil
}

func (P *POJSubmitter) SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	url := remote_oj.POJInfo.Host + remote_oj.POJInfo.SubmitUrl
	method := "POST"

	reqBody := fmt.Sprintf("problem_id=%v&language=%v&source=%v&submit=Submit&encoded=%v", info.RemoteProblemId, info.RemoteLanguage, info.SourceCode, 1)
	//fmt.Println(reqBody)
	payload := strings.NewReader(reqBody)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	if P.NeedLogin() {
		req, err = oj.PojLoginer.Login(req, account)
		if err != nil {
			return err
		}
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (P *POJSubmitter) Encode(code string) (string, error) {
	return code, nil
}
