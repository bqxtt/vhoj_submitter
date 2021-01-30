package oj

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter"
)

var PojSubmitter submitter.ISubmitter = &POJSubmitter{}

type POJSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (P *POJSubmitter) GetOjInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (P *POJSubmitter) NeedLogin() bool {
	panic("implement me")
}

func (P *POJSubmitter) GetMaxRunId() int64 {
	panic("implement me")
}

func (P *POJSubmitter) SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	panic("implement me")
}
