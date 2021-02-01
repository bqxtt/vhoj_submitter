package task_func

import (
	"fmt"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/manager"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter/holder"
)

func SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) {
	submitterImpl := holder.Submitters[info.RemoteOJ]
	err := submitterImpl.SubmitCode(info, account)
	if err != nil {
		//log
		//db result submit error
		return
	}
	fmt.Printf("submit code! info problem id: %v\n", info.RemoteProblemId)
	runId, err := submitterImpl.GetMaxRunId(info, account)
	if err != nil {
		//log
		//db result
		return
	}
	info.RealRunId = runId
	manager.ResultQuery(info, account)
}
