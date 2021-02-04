package accounter

import "github.com/ecnuvj/vhoj_submitter/pkg/common"

type IAccounter interface {
	FindAccount() (*common.RemoteAccount, error)
	ReleaseAccount(*common.RemoteAccount) error
	HandleSubmit(*common.SubmissionInfo) error
	HandleQuery(*common.SubmissionInfo, *common.RemoteAccount) error
	mustEmbedDefaultAccounterImpl()
}

type DefaultAccounterImpl struct {
	WaitingQueue []*common.SubmissionInfo
}

func (DefaultAccounterImpl) FindAccount() (*common.RemoteAccount, error) {
	panic("implement me")
}

func (DefaultAccounterImpl) ReleaseAccount(*common.RemoteAccount) error {
	panic("implement me")
}

func (DefaultAccounterImpl) HandleSubmit(*common.SubmissionInfo) error {
	panic("implement me")
}

func (DefaultAccounterImpl) HandleQuery(*common.SubmissionInfo, *common.RemoteAccount) error {
	panic("implement me")
}

func (DefaultAccounterImpl) mustEmbedDefaultAccounterImpl() {
	panic("implement me")
}
