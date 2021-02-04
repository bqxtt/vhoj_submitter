package _func

import (
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/status_type"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/manager"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/submitter/holder"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
)

func SubmitCode(info *common.SubmissionInfo, account *common.RemoteAccount) {
	submitterImpl := holder.Submitters[info.RemoteOJ]
	err := submitterImpl.SubmitCode(info, account)
	if err != nil {
		//log
		_ = submission_mapper.SubmissionMapper.UpdateSubmissionResultById(info.SubmissionID, status_type.SUBMIT_FAILED_PERM)
		return
	}
	fmt.Printf("submit code! info problem id: %v\n", info.RemoteProblemId)
	runId, err := submitterImpl.GetMaxRunId(info, account)
	if err != nil {
		//log
		_ = submission_mapper.SubmissionMapper.UpdateSubmissionResultById(info.SubmissionID, status_type.FAILED_OTHER)
		return
	}
	info.RealRunId = runId
	manager.ResultQuery(info, account)
}
