package submitter

import "github.com/bqxtt/vhoj_common/pkg/common/constants"

type ISubmitter interface {
	GetOjInfo() *constants.RemoteOJInfo
	NeedLogin() bool
	GetMaxRunId() int64
	SubmitCode() error
	mustEmbedDefaultSubmitter()
}

type DefaultSubmitterImpl struct{}

func (DefaultSubmitterImpl) GetOjInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (DefaultSubmitterImpl) NeedLogin() bool {
	panic("implement me")
}

func (DefaultSubmitterImpl) GetMaxRunId() int64 {
	panic("implement me")
}

func (DefaultSubmitterImpl) SubmitCode() error {
	panic("implement me")
}

func (DefaultSubmitterImpl) mustEmbedDefaultSubmitter() {

}
