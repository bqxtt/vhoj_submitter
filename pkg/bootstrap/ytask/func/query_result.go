package _func

import (
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/status_type"
	ytask "github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/querier/holder"
	"time"
)

func QueryResult(info *common.SubmissionInfo) {
	querier := holder.Queriers[info.RemoteOJ]
	status, err := querier.Query(info)
	if err != nil {
		//
		return
	}
	//todo db set result
	fmt.Printf("raw status: %v, type: %v\n", status.RawStatus, status.Status.StatusType)
	if status.Status.StatusType == status_type.CE {
		//todo db set ce info
		fmt.Printf("ce info: %v\n", status.CEInfo)
	}
	if !status.Status.Finished {
		_, err := ytask.Client.SetTaskCtl(ytask.Client.RunAfter, 1*time.Second).Send("code", "query_result", info)
		if err != nil {
			return
		}
	}
}
