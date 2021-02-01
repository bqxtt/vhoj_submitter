package oj

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUSubmitter_GetMaxRunId(t *testing.T) {
	info := &common.SubmissionInfo{
		RemoteOJ:        constants.HDU,
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
