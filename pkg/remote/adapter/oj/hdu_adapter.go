package oj

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants/language"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/submission_status"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/adapter"
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
	default:
		return submission_status.FAILED_OTHER
	}
}

func (H *HDUAdapter) ToCommonLanguage(s string) language.Language {
	panic("implement me")
}

func (H *HDUAdapter) ToOJLanguage(language language.Language) string {
	panic("implement me")
}
