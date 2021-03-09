package test

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/server"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/manager"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSubmitCode(t *testing.T) {
	server.InitTaskServer()
	client.InitTaskClient()
	database.InitDatabase()
	info := &common.SubmissionInfo{
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode: "#include <bits/stdc++.h>\n" +
			"using namespace std;\n" +
			"int main()\n" +
			"{\n" +
			"int a,b;\n" +
			"cin >> a >> b;\n" +
			"cout << a + b << endl;\n" +
			"return 0;\n" +
			"}\n",
		//SourceCode:      "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=",
	}
	manager.SubmitCode(info)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func TestMockHandlerSubmitCode(t *testing.T) {
	bootstrap.Init()
	//handler层的工作
	var submissionId uint = 4
	submission, err := submission_mapper.SubmissionMapper.FindSubmissionById(submissionId)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	info := &common.SubmissionInfo{
		SubmissionID:    submission.ID,
		RemoteOJ:        remote_oj.POJ,
		RemoteProblemId: "1000",
		RemoteLanguage:  "0",
		RealRunId:       "",
		SourceCode:      submission.SubmissionCode.SourceCode,
	}
	manager.SubmitCode(info)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
