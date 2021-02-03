package oj

import (
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/status_type"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/adapter/holder"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/querier"
	"github.com/bqxtt/vhoj_submitter/pkg/util"
)

var HduQuerier querier.IQuerier = &HDUQuerier{}

type HDUQuerier struct {
	querier.DefaultQuerierImpl
}

func (H *HDUQuerier) GetOjInfo() *remote_oj.RemoteOJInfo {
	return remote_oj.HDUInfo
}

func (H *HDUQuerier) Query(info *common.SubmissionInfo) (*common.RemoteSubmissionResult, error) {
	url := remote_oj.HDUInfo.Host + fmt.Sprintf(remote_oj.HDUInfo.ResultUrl, info.RealRunId)
	body, err := util.Download(url, "GET")
	if err != nil {
		return nil, err
	}
	reg := fmt.Sprintf(">%v</td><td>[\\s\\S]*?</td><td>([\\s\\S]*?)</td><td>[\\s\\S]*?</td><td>(\\d*?)MS</td><td>(\\d*?)K</td>", info.RealRunId)
	matches, err := util.ParseHtmlRegSlice(reg, body)
	if err != nil {
		//log
		return nil, err
	}
	var result, exeTime, exeMemory, ceInfo string
	if len(matches) > 1 {
		result = util.HtmlTagFilter(matches[1])
	}
	status := holder.Adapters[remote_oj.HDU].ToCommonStatus(result)
	if status.StatusType == status_type.AC {
		if len(matches) > 2 {
			exeTime = util.HtmlTagFilter(matches[2])
		}
		if len(matches) > 3 {
			exeMemory = util.HtmlTagFilter(matches[3])
		}
	}

	if status.StatusType == status_type.CE {
		url = remote_oj.HDUInfo.Host + fmt.Sprintf(remote_oj.HDUInfo.CompileInfoUrl, info.RealRunId)
		body, err = util.Download(url, "GET")
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
