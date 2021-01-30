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
	SubmitCode(&common.SubmissionInfo{RemoteProblemId: "1000", RemoteOJ: constants.HDU})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
