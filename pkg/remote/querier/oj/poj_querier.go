package oj

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter/holder"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/loginer/oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/querier"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
)

var PojQuerier querier.IQuerier = &POJQuerier{}

type POJQuerier struct {
	querier.DefaultQuerierImpl
}

func (P *POJQuerier) GetOjInfo() *remote_oj.RemoteOJInfo {
	return remote_oj.POJInfo
}

func (P *POJQuerier) Query(info *common.SubmissionInfo, account *common.RemoteAccount) (*common.RemoteSubmissionResult, error) {
	url := remote_oj.POJInfo.Host + fmt.Sprintf(remote_oj.POJInfo.ResultUrl, info.RealRunId)
	loginer := &util.LoginFunc{Account: account, Login: oj.PojLoginer.Login}
	body, err := util.Download(url, "GET", loginer)
	if err != nil {
		return nil, err
	}
	result, err := util.ParseHtmlReg(`<b>Result:</b>(.+?)</td>`, body)
	if err != nil {
		return nil, err
	}
	status := holder.Adapters[remote_oj.POJ].ToCommonStatus(result)
	var exeTime, exeMemory string
	if status.StatusType == status_type.AC {
		exeMemory, err = util.ParseHtmlReg(`<b>Memory:</b> ([-\d]+)`, body)
		if err != nil {
			//todo log memory parse error
			fmt.Printf("memory parse error\n")
		}
		exeTime, err = util.ParseHtmlReg(`<b>Time:</b> ([-\d]+)`, body)
		if err != nil {
			//todo log time parse error
			fmt.Printf("time parse error\n")
		}
	}

	var ceInfo string
	if status.StatusType == status_type.CE {
		url = remote_oj.POJInfo.Host + fmt.Sprintf(remote_oj.POJInfo.CompileInfoUrl, info.RealRunId)
		body, err = util.Download(url, "GET", loginer)
		if err != nil {
			//todo log get compile error info error
		} else {
			ceInfo, err = util.ParseHtmlReg(`(<pre>[\s\S]*?</pre>)`, body)
			if err != nil {
				//todo log
			}
		}
	}
	return &common.RemoteSubmissionResult{
		RawStatus: result,
		Status:    status,
		ExeTime:   util.ParseStringToInt64(exeTime),
		ExeMemory: util.ParseStringToInt64(exeMemory),
		CEInfo:    ceInfo,
	}, nil
}
