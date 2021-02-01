package holder

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/adapter"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/adapter/oj"
)

var Adapters map[remote_oj.RemoteOJ]adapter.IAdapter

func init() {
	Adapters = make(map[remote_oj.RemoteOJ]adapter.IAdapter)
	Adapters[remote_oj.HDU] = oj.HduAdapter
}
