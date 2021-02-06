package oj

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUSubmitter_GetMaxRunId(t *testing.T) {
	info := &common.SubmissionInfo{
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "",
	}
	account := &common.RemoteAccount{
		Username: "bqxtt233",
		Password: "tcg19981108",
	}
	HduSubmitter.GetMaxRunId(info, account)
}

func TestHDUSubmitter_SubmitCode(t *testing.T) {
	info := &common.SubmissionInfo{
		SubmissionID:    3,
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=",
	}
	account := &common.RemoteAccount{
		Username: "bqxtt233",
		Password: "tcg19981108",
	}
	HduSubmitter.SubmitCode(info, account)
}

func TestPOJSubmitter_SubmitCode(t *testing.T) {
	info := &common.SubmissionInfo{
		SubmissionID:    3,
		RemoteOJ:        remote_oj.POJ,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "I2luY2x1ZGUgPHN0ZGlvLmg%2BCgppbnQgbWFpbigpCnsKICAgIGludCBhLGI7CiAgICBzY2FuZigiJWQgJWQiLCZhLCAmYik7CiAgICBwcmludGYoIiVkXG4iLGErYik7CiAgICByZXR1cm4gMDsKfQ%3D%3D",
	}
	account := &common.RemoteAccount{
		Username: "bqx",
		Password: "tcg19981108",
	}
	PojSubmitter.SubmitCode(info, account)
}

func TestPOJSubmitter_GetMaxRunId(t *testing.T) {
	info := &common.SubmissionInfo{
		SubmissionID:    3,
		RemoteOJ:        remote_oj.POJ,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "I2luY2x1ZGUgPHN0ZGlvLmg%2BCgppbnQgbWFpbigpCnsKICAgIGludCBhLGI7CiAgICBzY2FuZigiJWQgJWQiLCZhLCAmYik7CiAgICBwcmludGYoIiVkXG4iLGErYik7CiAgICByZXR1cm4gMDsKfQ%3D%3D",
	}
	account := &common.RemoteAccount{
		Username: "bqx",
		Password: "tcg19981108",
	}
	runID, _ := PojSubmitter.GetMaxRunId(info, account)
	fmt.Println(runID)
}
