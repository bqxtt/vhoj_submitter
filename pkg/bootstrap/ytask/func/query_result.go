package _func

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	ytask "github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/querier/holder"
	"github.com/jinzhu/gorm"
	"time"
)

func QueryResult(info *common.SubmissionInfo, account *common.RemoteAccount) {
	querier := holder.Queriers[info.RemoteOJ]
	status, err := querier.Query(info, account)
	if err != nil {
		//
		return
	}
	_, _ = submission_mapper.SubmissionMapper.UpdateSubmissionById(&model.Submission{
		Model: gorm.Model{
			ID: info.SubmissionID,
		},
		Result: status.Status.StatusType,
	})
	fmt.Printf("raw status: %v, type: %v\n", status.RawStatus, status.Status.StatusType)
	if status.Status.StatusType == status_type.CE {
		_ = submission_mapper.SubmissionMapper.UpdateSubmissionCEInfoById(info.SubmissionID, status.CEInfo)
		fmt.Printf("ce info: %v\n", status.CEInfo)
	}
	if !status.Status.Finished {
		_, err := ytask.Client.SetTaskCtl(ytask.Client.RunAfter, 1*time.Second).Send("code", "query_result", info, account)
		if err != nil {
			return
		}
	}
}
