package server

import (
	"github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/func"
	ytask "github.com/gojuukaze/YTask/v2"
)

func InitTaskServer() {
	broker := ytask.Broker.NewRedisBroker("139.9.140.136", "6379", "19981108", 0, 0)
	backend := ytask.Backend.NewRedisBackend("139.9.140.136", "6379", "19981108", 0, 0)
	ser := ytask.Server.NewServer(
		ytask.Config.Broker(&broker),
		ytask.Config.Backend(&backend),
		ytask.Config.Debug(true),
		ytask.Config.StatusExpires(60*5),
		ytask.Config.ResultExpires(60*5),
	)
	ser.Add("code", "submit_code", _func.SubmitCode)
	ser.Add("code", "query_result", _func.QueryResult)

	ser.Run("code", 30, true)

	//ser.Shutdown(context.Background())
}
