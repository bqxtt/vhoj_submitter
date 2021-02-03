package test

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/server"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/manager"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSubmitCode(t *testing.T) {
	server.InitTaskServer()
	client.InitTaskClient()
	info := &common.SubmissionInfo{
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      "JTIzaW5jbHVkZSUyMCUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzb2FjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSU3QiUwQWludCUyMGElMkNiJTBBcmV0dXJuJTIwMCUzQiUwQSU3RA==",
		//SourceCode:      "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=",
	}
	manager.SubmitCode(info)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
