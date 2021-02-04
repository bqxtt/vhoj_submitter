package manager

import (
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/accounter/holder"
)

func SubmitCode(info *common.SubmissionInfo) {
	//1.在handler层选择oj，写入info
	//2.选择对应oj可用账号，委托给对应repo来做
	accounter := holder.Accounters[info.RemoteOJ]
	accounter.HandleSubmit(info)
}
