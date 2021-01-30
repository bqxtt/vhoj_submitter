package common

import "github.com/bqxtt/vhoj_common/pkg/common/constants"

type SubmissionInfo struct {
	RemoteOJ        constants.RemoteOJ
	RemoteProblemId string
	RemoteLanguage  string
	RealRunId       string
	SourceCode      string
}

type RemoteAccount struct {
	Username string
	Password string
}

//todo move to common
type RemoteResult int32

const (
	AC RemoteResult = 1
	CE RemoteResult = 2
)

type Language int32

const (
	CPP  Language = 1
	JAVA Language = 2
)
