package service

import (
	"context"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/user_problem_status"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/contest_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/user_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_remote"
	"github.com/ecnuvj/vhoj_rpc/model/remotepb"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
	"github.com/jinzhu/gorm"
)

type SubmitService struct {
}

func (ss *SubmitService) SubmitCode(submission *model.Submission) (*model.Submission, error) {
	remoteReq, err := ss.chooseSuitableRemoteOJ(submission)
	if err != nil {
		return nil, err
	}
	_ = user_mapper.UserMapper.AddUserSubmitCountById(submission.UserId)
	if submission.ContestId != 0 {
		//更新比赛题目提交记录
		_ = problem_mapper.ProblemMapper.AddContestProblemSubmittedCountById(submission.ContestId, submission.ProblemId)
	} else {
		//更新普通题目提交数
		_ = problem_mapper.ProblemMapper.AddProblemSubmittedCountById(submission.ProblemId)
	}

	//manager.SubmitCode(info)
	resp, err := rpc_remote.RemoteServiceClient.SubmitCode(context.Background(), remoteReq)
	if err != nil {
		return nil, err
	}
	submission.ID = uint(resp.SubmissionId)
	return submission, nil
}

func (ss *SubmitService) ReSubmitCode(submissionId uint) error {
	_ = submission_mapper.SubmissionMapper.ResetSubmissionById(submissionId)
	submission, err := submission_mapper.SubmissionMapper.FindSubmissionById(submissionId)
	if err != nil {
		return err
	}
	_, err = ss.SubmitCode(submission)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SubmitService) chooseSuitableRemoteOJ(submission *model.Submission) (*remotepb.SubmitCodeRequest, error) {
	problemGroups, err := problem_mapper.ProblemMapper.FindGroupProblemsById(submission.ProblemId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			_ = problem_mapper.ProblemMapper.DeleteProblemById(submission.ProblemId)
		}
		return nil, err
	}
	//todo 选择最优oj
	remoteOJ := problemGroups[0].RemoteOJ
	remoteProblemId := problemGroups[0].RemoteProblemId

	//submission.RemoteOJ = remoteOJ
	//submission, err = submission_mapper.SubmissionMapper.AddOrModifySubmission(submission)
	//if err != nil {
	//	return nil, nil, err
	//}
	//return submission, &common.SubmissionInfo{
	//	SubmissionID:    submission.ID,
	//	ProblemID:       submission.ProblemId,
	//	RemoteOJ:        remoteOJ,
	//	RemoteProblemId: remoteProblemId,
	//	RemoteLanguage:  holder.Adapters[remoteOJ].ToOJLanguage(submission.Language),
	//	SourceCode:      submission.SubmissionCode.SourceCode,
	//}, nil
	return &remotepb.SubmitCodeRequest{
		UserId:          uint64(submission.UserId),
		Username:        submission.Username,
		RemoteOj:        int32(remoteOJ),
		RemoteProblemId: remoteProblemId,
		ProblemId:       uint64(submission.ProblemId),
		Language:        int32(submission.Language),
		SourceCode:      submission.SubmissionCode.SourceCode,
		ContestId:       uint64(submission.ContestId),
	}, nil
}

func (SubmitService) ListSubmissions(pageNo int32, pageSize int32, condition *common.SubmissionSearchCondition) ([]*submitterpb.Submission, *common.PageInfo, error) {
	submissions, count, err := submission_mapper.SubmissionMapper.FindSubmissions(pageNo, pageSize, &submission_mapper.SearchSubmissionCondition{
		Username:  condition.Username,
		ProblemId: condition.ProblemId,
		Status:    condition.Result,
		Language:  condition.Language,
	})
	if err != nil {
		return nil, nil, err
	}
	return util.ModelSubmissionsToRpcSubmissions(submissions), &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (SubmitService) CheckUserProblemStatus(userId uint, problemId uint, contestId uint) (user_problem_status.UserProblemStatus, error) {
	submissions, err := submission_mapper.SubmissionMapper.FindSubmissionsGroupByResult(&submission_mapper.UserSubmissionCondition{
		UserId:    userId,
		ProblemId: problemId,
		ContestId: contestId,
	})
	if err != nil {
		return user_problem_status.UNKNOWN, err
	}
	if len(submissions) == 0 {
		return user_problem_status.NOT_SUBMIT, nil
	}
	for _, s := range submissions {
		if s.Result == status_type.AC {
			return user_problem_status.ACCEPTED, nil
		}
	}
	return user_problem_status.ATTEMPTED, nil
}

func (SubmitService) GetSubmission(submissionId uint) (*submitterpb.Submission, error) {
	submission, err := submission_mapper.SubmissionMapper.FindSubmissionById(submissionId)
	if err != nil {
		return nil, err
	}
	return util.ModelSubmissionToRpcSubmission(submission), nil
}

func (SubmitService) GetContestSubmissions(contestId uint) ([]*submitterpb.Submission, error) {
	contest, err := contest_mapper.ContestMapper.FindContestById(contestId)
	if err != nil {
		return nil, err
	}
	submissions, err := submission_mapper.SubmissionMapper.FindSubmissionsByContestId(contestId, contest.StartTime, contest.EndTime)
	if err != nil {
		return nil, err
	}
	return util.ModelSubmissionsToRpcSubmissions(submissions), nil
}
