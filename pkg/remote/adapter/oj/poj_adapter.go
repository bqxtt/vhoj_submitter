package oj

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/submission_status"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter"
)

var PojAdapter adapter.IAdapter = &POJAdapter{}

type POJAdapter struct {
	adapter.DefaultAdapterImpl
}

func (P *POJAdapter) ToCommonStatus(s string) *submission_status.SubmissionStatus {
	switch s {
	case "Accepted":
		return submission_status.AC
	case "Wrong Answer":
		return submission_status.WA
	case "Compile Error":
		return submission_status.CE
	case "Runtime Error":
		return submission_status.RE
	case "Memory Limit Exceeded":
		return submission_status.MLE
	case "Time Limit Exceeded":
		return submission_status.TLE
	case "Output Limit Exceeded":
		return submission_status.OLE
	case "Presentation Error":
		return submission_status.PE
	case "Waiting":
		return submission_status.QUEUEING
	case "Compiling":
		return submission_status.COMPILING
	case "Running":
		return submission_status.JUDGING
	default:
		return submission_status.FAILED_OTHER
	}
}

func (P *POJAdapter) ToCommonLanguage(s string) language.Language {
	panic("implement me")
}

func (P *POJAdapter) ToOJLanguage(language language.Language) string {
	panic("implement me")
}
