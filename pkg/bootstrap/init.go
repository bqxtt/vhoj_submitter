package bootstrap

import (
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/server"
)

func Init() {
	server.InitTaskServer()
	client.InitTaskClient()
	InitDatabase()
}
