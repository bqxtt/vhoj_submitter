package task_func

import (
	"fmt"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter/holder"
)

func SubmitCode(submission *common.SubmissionInfo, account *common.RemoteAccount) {
	submitterImpl := holder.Submitters[submission.RemoteOJ]
	submitterImpl.SubmitCode(submission, account)
	fmt.Printf("submit code! submission id: %v\n", submission.RemoteProblemId)
}
