package accounter

import "github.com/bqxtt/vhoj_submitter/pkg/common"

type IAccounter interface {
	FindAccount() (*common.RemoteAccount, error)
	ReleaseAccount(account *common.RemoteAccount) error
	Handle(submission *common.SubmissionInfo) error
	mustEmbedDefaultAccounterImpl()
}

type DefaultAccounterImpl struct {
	WaitingQueue []*common.SubmissionInfo
}

func (DefaultAccounterImpl) FindAccount() (*common.RemoteAccount, error) {
	panic("implement me")
}

func (DefaultAccounterImpl) ReleaseAccount(account *common.RemoteAccount) error {
	panic("implement me")
}

func (DefaultAccounterImpl) Handle(submission *common.SubmissionInfo) error {
	panic("implement me")
}

func (DefaultAccounterImpl) mustEmbedDefaultAccounterImpl() {
	panic("implement me")
}
