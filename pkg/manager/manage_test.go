package manager

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSubmitCode(t *testing.T) {
	bootstrap.Init()
	info := &common.SubmissionInfo{
		RemoteOJ:        constants.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=",
	}
	SubmitCode(info)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
