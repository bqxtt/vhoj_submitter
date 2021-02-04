package manager

import (
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/accounter/holder"
)

//提交代码并获取real runid之后结果查询
//解锁账号
func ResultQuery(info *common.SubmissionInfo, account *common.RemoteAccount) {
	accounter := holder.Accounters[info.RemoteOJ]
	_ = accounter.HandleQuery(info, account)
}
