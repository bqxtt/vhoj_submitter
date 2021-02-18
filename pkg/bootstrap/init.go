package bootstrap

import (
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/server"
)

func Init() {
	server.InitTaskServer()
	client.InitTaskClient()
	database.InitDatabase()
	//rpc_service.InitService()
}
