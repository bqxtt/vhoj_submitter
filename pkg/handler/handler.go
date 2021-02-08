package handler

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/base"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitpb"
	"github.com/ecnuvj/vhoj_submitter/pkg/service"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
)

type SubmitHandler struct {
	submitpb.UnimplementedSubmitServiceServer
	submitService *service.SubmitService
}

func NewSubmitHandler() *SubmitHandler {
	return &SubmitHandler{
		submitService: &service.SubmitService{},
	}
}

func (s *SubmitHandler) SubmitCode(ctx context.Context, request *submitpb.SubmitCodeRequest) (*submitpb.SubmitCodeResponse, error) {
	if request == nil {
		return &submitpb.SubmitCodeResponse{
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
		Language:  language.Language(request.Language),
		ContestId: uint(request.ContestId),
	}
	var err error
	submission, err = s.submitService.SubmitCode(submission)
	if err != nil {
		return &submitpb.SubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "submit code error"),
		}, fmt.Errorf("submit code error: %v", err)
	}
	return &submitpb.SubmitCodeResponse{
		SubmissionId: uint64(submission.ID),
		BaseResponse: util.NewDefaultSuccessReply(),
	}, nil
}

func (s *SubmitHandler) ReSubmitCode(ctx context.Context, request *submitpb.ReSubmitCodeRequest) (*submitpb.ReSubmitCodeResponse, error) {
	if request == nil {
		return &submitpb.ReSubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "re submit code request is nil"),
		}, fmt.Errorf("re submit code request is nil")
	}
	submissionId := request.SubmissionId
	err := s.submitService.ReSubmitCode(uint(submissionId))
	if err != nil {
		return &submitpb.ReSubmitCodeResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "re submit err: %v", err),
		}, fmt.Errorf("re submit err: %v", err)
	}
	return &submitpb.ReSubmitCodeResponse{
		BaseResponse: util.NewDefaultSuccessReply(),
	}, nil
}