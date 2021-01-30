package bootstrap

import "github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask"

func Init() {
	ytask.InitTaskServer()
	ytask.InitTaskClient()
}
