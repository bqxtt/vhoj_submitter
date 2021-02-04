package oj

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUSubmitter_GetMaxRunId(t *testing.T) {
	info := &common.SubmissionInfo{
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "",
	}
	account := &common.RemoteAccount{
		Username: "bqxtt233",
		Password: "tcg19981108",
	}
	HduSubmitter.GetMaxRunId(info, account)
}
