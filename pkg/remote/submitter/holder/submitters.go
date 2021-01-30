package holder

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter/oj"
)

var Submitters map[constants.RemoteOJ]submitter.ISubmitter

func init() {
	Submitters = make(map[constants.RemoteOJ]submitter.ISubmitter, 0)
	Submitters[constants.POJ] = oj.PojSubmitter
	Submitters[constants.HDU] = oj.HduSubmitter
}
