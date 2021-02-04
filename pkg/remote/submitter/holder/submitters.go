package holder

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/submitter"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/submitter/oj"
)

var Submitters map[remote_oj.RemoteOJ]submitter.ISubmitter

func init() {
	Submitters = make(map[remote_oj.RemoteOJ]submitter.ISubmitter, 0)
	Submitters[remote_oj.POJ] = oj.PojSubmitter
	Submitters[remote_oj.HDU] = oj.HduSubmitter
}
