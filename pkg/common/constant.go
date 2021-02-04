package common

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/submission_status"
)

type SubmissionInfo struct {
	SubmissionID    uint
	RemoteOJ        remote_oj.RemoteOJ
	RemoteProblemId string
	RemoteLanguage  string
	RealRunId       string
	SourceCode      string
}

type RemoteAccount struct {
	Username string
	Password string
}

type RemoteSubmissionResult struct {
	RawStatus string
	Status    *submission_status.SubmissionStatus
	ExeTime   int64
	ExeMemory int64
	CEInfo    string
}
