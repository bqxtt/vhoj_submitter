package holder

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/accounter"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/accounter/oj"
)

var Accounters map[constants.RemoteOJ]accounter.IAccounter

func init() {
	Accounters = make(map[constants.RemoteOJ]accounter.IAccounter)
	Accounters[constants.HDU] = oj.HduAccounter
}
