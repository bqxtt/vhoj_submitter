package oj

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/submission_status"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter"
)

var HduAdapter adapter.IAdapter = &HDUAdapter{}

type HDUAdapter struct {
	adapter.DefaultAdapterImpl
}

func (H *HDUAdapter) ToCommonStatus(s string) *submission_status.SubmissionStatus {
	switch s {
	case "Accepted":
		return submission_status.AC
	case "Wrong Answer":
		return submission_status.WA
	case "Compilation Error":
		return submission_status.CE
	case "Presentation Error":
		return submission_status.PE
	case "Runtime Error(ACCESS_VIOLATION)":
		return submission_status.RE
	case "Time Limit Exceeded":
		return submission_status.TLE
	case "Memory Limit Exceeded":
		return submission_status.MLE
	case "Output Limit Exceeded":
		return submission_status.OLE
	case "Queuing":
		return submission_status.QUEUEING
	case "Compiling":
		return submission_status.COMPILING
	case "Running":
		return submission_status.JUDGING
	default:
		return submission_status.FAILED_OTHER
	}
}

func (H *HDUAdapter) ToCommonLanguage(s string) language.Language {
	panic("implement me")
}

func (H *HDUAdapter) ToOJLanguage(lang language.Language) string {
	switch lang {
	case language.CPP:
		return "0"
	case language.JAVA:
		return "5"
	default:
		return "-1"
	}
}
