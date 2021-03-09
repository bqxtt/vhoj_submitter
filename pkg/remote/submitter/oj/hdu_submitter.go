package oj

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/loginer/oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/submitter"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

var HduSubmitter submitter.ISubmitter = &HDUSubmitter{}

type HDUSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (H *HDUSubmitter) GetOjInfo() *remote_oj.RemoteOJInfo {
	return remote_oj.HDUInfo
}

func (H *HDUSubmitter) NeedLogin() bool {
	return true
}

func (H *HDUSubmitter) GetMaxRunId(info *common.SubmissionInfo, account *common.RemoteAccount) (string, error) {
	url := remote_oj.HDUInfo.Host + fmt.Sprintf(remote_oj.HDUInfo.StatusUrl, account.Username, info.RemoteProblemId)
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
	runId, err := util.ParseHtmlReg(`<td height=22px>(\d+)`, string(body))
	if err != nil {
		return "", nil
	}
	return runId, nil
}

func (H *HDUSubmitter) SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	url := remote_oj.HDUInfo.Host + remote_oj.HDUInfo.SubmitUrl
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("check", "0")
	code, err := H.Encode(info.SourceCode)
	if err != nil {
		return err
	}
	_ = writer.WriteField("_usercode", code)
	_ = writer.WriteField("problemid", info.RemoteProblemId)
	_ = writer.WriteField("language", info.RemoteLanguage)
	err = writer.Close()
	if err != nil {
		//fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	if H.NeedLogin() {
		req, err = oj.HduLoginer.Login(req, account)
		if err != nil {
			return err
		}
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

	return nil
}

func (H *HDUSubmitter) Encode(code string) (string, error) {
	urlEncode := url.QueryEscape(code)
	encodeString := base64.StdEncoding.EncodeToString([]byte(urlEncode))
	return encodeString, nil
}
