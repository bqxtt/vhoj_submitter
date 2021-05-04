package bootstrap

import (
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/rpc_client"
	"github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/rpc_service"
)

func Init() {
	//server.InitTaskServer()
	//client.InitTaskClient()
	database.InitDatabase()
	rpc_client.Init()
	rpc_service.InitService()
}
