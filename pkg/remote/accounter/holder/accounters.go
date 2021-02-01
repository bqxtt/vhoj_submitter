package holder

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/accounter"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/accounter/oj"
)

var Accounters map[remote_oj.RemoteOJ]accounter.IAccounter

func init() {
	Accounters = make(map[remote_oj.RemoteOJ]accounter.IAccounter)
	Accounters[remote_oj.HDU] = oj.HduAccounter
}
