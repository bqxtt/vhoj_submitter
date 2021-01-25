package poj

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/submitter"
)

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

func (P *POJSubmitter) SubmitCode() error {
	panic("implement me")
}
