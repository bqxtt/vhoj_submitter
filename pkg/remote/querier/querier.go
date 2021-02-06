package querier

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
)

type IQuerier interface {
	GetOjInfo() *remote_oj.RemoteOJInfo
	Query(info *common.SubmissionInfo, account *common.RemoteAccount) (*common.RemoteSubmissionResult, error)
	mustEmbedDefaultQuerierImpl()
}

type DefaultQuerierImpl struct{}

func (DefaultQuerierImpl) GetOjInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (DefaultQuerierImpl) Query(*common.SubmissionInfo, *common.RemoteAccount) (*common.RemoteSubmissionResult, error) {
	panic("implement me")
}

func (DefaultQuerierImpl) mustEmbedDefaultQuerierImpl() {
	panic("implement me")
}
