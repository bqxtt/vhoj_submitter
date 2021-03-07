package handler

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/base"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"github.com/ecnuvj/vhoj_submitter/pkg/service"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
)

type SubmitHandler struct {
	submitterpb.UnimplementedSubmitServiceServer
	submitService *service.SubmitService
}

func NewSubmitHandler() (*SubmitHandler, error) {
	return &SubmitHandler{
		submitService: &service.SubmitService{},
	}, nil
}

func (s *SubmitHandler) SubmitCode(ctx context.Context, request *submitterpb.SubmitCodeRequest) (*submitterpb.SubmitCodeResponse, error) {
	if request == nil {
		return &submitterpb.SubmitCodeResponse{
			SubmissionId: 0,
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "submit code request is nil"),
		}, fmt.Errorf("submit code request is nil")
	}
	submission := &model.Submission{
		SubmissionCode: &model.SubmissionCode{
			SourceCode: request.SourceCode,
			CodeLength: int64(len(request.SourceCode)),
		},
		ProblemId: uint(request.ProblemId),
		UserId:    uint(request.UserId),
		Username:  request.Username,
		Language:  language.Language(request.Language),
		ContestId: uint(request.ContestId),
	}
	var err error
	submission, err = s.submitService.SubmitCode(submission)
	if err != nil {
		return &submitterpb.SubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "submit code error"),
		}, fmt.Errorf("submit code error: %v", err)
	}
	return &submitterpb.SubmitCodeResponse{
		SubmissionId: uint64(submission.ID),
		BaseResponse: util.NewDefaultSuccessReply(),
	}, nil
}

func (s *SubmitHandler) ReSubmitCode(ctx context.Context, request *submitterpb.ReSubmitCodeRequest) (*submitterpb.ReSubmitCodeResponse, error) {
	if request == nil {
		return &submitterpb.ReSubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "re submit code request is nil"),
		}, fmt.Errorf("re submit code request is nil")
	}
	submissionId := request.SubmissionId
	err := s.submitService.ReSubmitCode(uint(submissionId))
	if err != nil {
		return &submitterpb.ReSubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "re submit err: %v", err),
		}, fmt.Errorf("re submit err: %v", err)
	}
	return &submitterpb.ReSubmitCodeResponse{
		BaseResponse: util.NewDefaultSuccessReply(),
	}, nil
}

func (s *SubmitHandler) ListSubmissions(ctx context.Context, request *submitterpb.ListSubmissionsRequest) (*submitterpb.ListSubmissionsResponse, error) {
	if request == nil {
		return &submitterpb.ListSubmissionsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "rpc request is nil"),
		}, fmt.Errorf("rpc request is nil")
	}
	condition := &common.SubmissionSearchCondition{
		Username:  request.Username,
		ProblemId: uint(request.ProblemId),
		Result:    status_type.SubmissionStatusType(request.Result),
		Language:  language.Language(request.Language),
	}
	submissions, pageInfo, err := s.submitService.ListSubmissions(request.PageNo, request.PageSize, condition)
	if err != nil {
		return &submitterpb.ListSubmissionsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &submitterpb.ListSubmissionsResponse{
		Submissions:  submissions,
		TotalCount:   pageInfo.TotalCount,
		TotalPages:   pageInfo.TotalPages,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}
