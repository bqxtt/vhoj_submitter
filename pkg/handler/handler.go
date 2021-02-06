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
	"github.com/jinzhu/gorm"
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
	s.submitService.SubmitCode(submission)
}
