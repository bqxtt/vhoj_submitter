package util

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"github.com/golang/protobuf/ptypes"
)

func ModelSubmissionToRpcSubmission(submission *model.Submission) *submitterpb.Submission {
	if submission == nil {
		return &submitterpb.Submission{}
	}
	var code string
	if submission.SubmissionCode != nil {
		code = submission.SubmissionCode.SourceCode
	}
	submitTime, _ := ptypes.TimestampProto(submission.CreatedAt)
	return &submitterpb.Submission{
		SubmissionId: uint64(submission.ID),
		ProblemId:    uint64(submission.ProblemId),
		UserId:       uint64(submission.UserId),
		Username:     submission.Username,
		Result:       int32(submission.Result),
		TimeCost:     submission.TimeCost,
		MemoryCost:   submission.MemoryCost,
		Language:     int32(submission.Language),
		Code:         code,
		SubmitTime:   submitTime,
	}
}

func ModelSubmissionsToRpcSubmissions(submissions []*model.Submission) []*submitterpb.Submission {
	retSubmissions := make([]*submitterpb.Submission, len(submissions))
	for i, s := range submissions {
		retSubmissions[i] = ModelSubmissionToRpcSubmission(s)
	}
	return retSubmissions
}
