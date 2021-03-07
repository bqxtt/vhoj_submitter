package util

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
)

func ModelSubmissionToRpcSubmission(submission *model.Submission) *submitterpb.Submission {
	if submission == nil {
		return &submitterpb.Submission{}
	}
	return &submitterpb.Submission{
		ProblemId:  uint64(submission.ProblemId),
		UserId:     uint64(submission.UserId),
		Username:   submission.Username,
		Result:     int32(submission.Result),
		TimeCost:   submission.TimeCost,
		MemoryCost: submission.MemoryCost,
		Language:   int32(submission.Language),
	}
}

func ModelSubmissionsToRpcSubmissions(submissions []*model.Submission) []*submitterpb.Submission {
	retSubmissions := make([]*submitterpb.Submission, len(submissions))
	for i, s := range submissions {
		retSubmissions[i] = ModelSubmissionToRpcSubmission(s)
	}
	return retSubmissions
}
