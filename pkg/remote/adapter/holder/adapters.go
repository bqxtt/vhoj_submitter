package holder

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter/oj"
)

var Adapters map[remote_oj.RemoteOJ]adapter.IAdapter

func init() {
	Adapters = make(map[remote_oj.RemoteOJ]adapter.IAdapter)
	Adapters[remote_oj.HDU] = oj.HduAdapter
	Adapters[remote_oj.POJ] = oj.PojAdapter
}
