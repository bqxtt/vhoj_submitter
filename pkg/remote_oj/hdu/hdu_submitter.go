package hdu

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/submitter"
)

type HDUSubmitter struct {
	submitter.DefaultSubmitterImpl
}

func (H *HDUSubmitter) GetOjInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (H *HDUSubmitter) NeedLogin() bool {
	panic("implement me")
}

func (H *HDUSubmitter) GetMaxRunId() int64 {
	panic("implement me")
}

func (H *HDUSubmitter) SubmitCode() error {
	panic("implement me")
}
