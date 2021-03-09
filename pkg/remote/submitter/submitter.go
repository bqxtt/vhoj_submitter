package submitter

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
)

type ISubmitter interface {
	GetOjInfo() *remote_oj.RemoteOJInfo
	NeedLogin() bool
	GetMaxRunId(info *common.SubmissionInfo, account *common.RemoteAccount) (string, error)
	SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) error
	Encode(string) (string, error)
	mustEmbedDefaultSubmitter()
}

type DefaultSubmitterImpl struct{}

func (DefaultSubmitterImpl) GetOjInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (DefaultSubmitterImpl) NeedLogin() bool {
	panic("implement me")
}

func (DefaultSubmitterImpl) GetMaxRunId(*common.SubmissionInfo, *common.RemoteAccount) (string, error) {
	panic("implement me")
}

func (DefaultSubmitterImpl) SubmitCode(*common.SubmissionInfo, *common.RemoteAccount) error {
	panic("implement me")
}

func (DefaultSubmitterImpl) Encode(rawCode string) (string, error) {
	panic("implement me")
}

func (DefaultSubmitterImpl) mustEmbedDefaultSubmitter() {

}
