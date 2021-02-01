package adapter

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants/language"
	"github.com/bqxtt/vhoj_common/pkg/common/constants/submission_status"
)

type IAdapter interface {
	ToCommonStatus(string) *submission_status.SubmissionStatus
	ToCommonLanguage(string) language.Language
	ToOJLanguage(language language.Language) string
	mustEmbedDefaultAdapterImpl()
}

type DefaultAdapterImpl struct {
}

func (DefaultAdapterImpl) ToCommonStatus(s string) *submission_status.SubmissionStatus {
	panic("implement me")
}

func (DefaultAdapterImpl) ToCommonLanguage(s string) language.Language {
	panic("implement me")
}

func (DefaultAdapterImpl) ToOJLanguage(language language.Language) string {
	panic("implement me")
}

func (DefaultAdapterImpl) mustEmbedDefaultAdapterImpl() {
	panic("implement me")
}
