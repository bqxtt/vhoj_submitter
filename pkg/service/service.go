package service

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/manager"
)

type SubmitService struct {
}

func (ss *SubmitService) SubmitCode(submission *model.Submission) (*model.Submission, error) {
	info, err := ss.chooseSuitableRemoteOJ(submission)
	if err != nil {
		return nil, err
	}
	manager.SubmitCode(info)

}

func (ss *SubmitService) chooseSuitableRemoteOJ(submission *model.Submission) (*common.SubmissionInfo, error) {
	var err error
	//todo
	submission, err = submission_mapper.SubmissionMapper.AddSubmission(submission)
	if err != nil {
		return nil, err
	}
	return &common.SubmissionInfo{
		SubmissionID:    0,
		RemoteOJ:        0,
		RemoteProblemId: "",
		RemoteLanguage:  "",
		RealRunId:       "",
		SourceCode:      "",
	}, nil
}
