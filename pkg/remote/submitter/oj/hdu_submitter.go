package oj

import (
	"bytes"
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/loginer/oj"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter"
	"mime/multipart"
	"net/http"
)

var HduSubmitter submitter.ISubmitter = &HDUSubmitter{}

type HDUSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (H *HDUSubmitter) GetOjInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (H *HDUSubmitter) NeedLogin() bool {
	return true
}

func (H *HDUSubmitter) GetMaxRunId() int64 {
	panic("implement me")
}

func (H *HDUSubmitter) SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error {

	url := "http://acm.hdu.edu.cn/submit.php?action=submit"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("check", "0")
	_ = writer.WriteField("_usercode", "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=")
	_ = writer.WriteField("problemid", "1000")
	_ = writer.WriteField("language", "0")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if H.NeedLogin() {
		cookies, err := oj.HduLoginer.Login(account)
		if err != nil {
			return err
		}
		if cookies != nil {
			for _, cookie := range cookies {
				req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)
			}
		}
	}
	//fmt.Printf("cookie: %v", req.Header.Get("Cookie"))

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	return nil
}
