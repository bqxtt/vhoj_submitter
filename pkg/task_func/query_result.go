package task_func

import (
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/querier/holder"
)

func QueryResult(info *common.SubmissionInfo) {
	querier := holder.Queriers[info.RemoteOJ]
	status, err := querier.Query(info)
	if err != nil {
		//
		return
	}

}
