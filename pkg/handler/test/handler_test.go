package test

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap"
	handler2 "github.com/ecnuvj/vhoj_submitter/pkg/handler"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSubmitHandler_SubmitCode(t *testing.T) {
	bootstrap.Init()
	handler, _ := handler2.NewSubmitHandler()
	req := &submitterpb.SubmitCodeRequest{
		ProblemId:  2,
		UserId:     1017510,
		Language:   1,
		ContestId:  0,
		SourceCode: "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=",
	}
	resp, err := handler.SubmitCode(context.Background(), req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Println(resp)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func TestSubmitHandler_ReSubmitCode(t *testing.T) {
	bootstrap.Init()
	handler, _ := handler2.NewSubmitHandler()
	request := &submitterpb.ReSubmitCodeRequest{SubmissionId: 7}
	resp, err := handler.ReSubmitCode(context.Background(), request)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(resp)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
