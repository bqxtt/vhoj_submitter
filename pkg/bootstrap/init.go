package bootstrap

import (
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/server"
)

func Init() {
	server.InitTaskServer()
	client.InitTaskClient()
	InitDatabase()
}
