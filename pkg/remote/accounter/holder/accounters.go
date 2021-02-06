package holder

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/accounter"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/accounter/oj"
)

var Accounters map[remote_oj.RemoteOJ]accounter.IAccounter

func init() {
	Accounters = make(map[remote_oj.RemoteOJ]accounter.IAccounter)
	Accounters[remote_oj.HDU] = oj.HduAccounter
	Accounters[remote_oj.POJ] = oj.PojAccounter
}
