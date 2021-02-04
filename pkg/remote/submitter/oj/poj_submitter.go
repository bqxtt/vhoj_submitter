package oj

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/submitter"
)

var PojSubmitter submitter.ISubmitter = &POJSubmitter{}

type POJSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (P *POJSubmitter) GetOjInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (P *POJSubmitter) NeedLogin() bool {
	panic("implement me")
}

func (P *POJSubmitter) GetMaxRunId(info *common.SubmissionInfo, account *common.RemoteAccount) (string, error) {
	panic("implement me")
}

func (P *POJSubmitter) SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	panic("implement me")
}
