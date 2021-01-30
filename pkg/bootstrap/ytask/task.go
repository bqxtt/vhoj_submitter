package ytask

import (
	"github.com/bqxtt/vhoj_submitter/pkg/task_func"
	ytask "github.com/gojuukaze/YTask/v2"
	"github.com/gojuukaze/YTask/v2/server"
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
	ser.Add("code", "submit_code", task_func.SubmitCode)

	ser.Run("code", 30)

	//quit := make(chan os.Signal, 1)
	//
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//ser.Shutdown(context.Background())
}

var Client server.Client

func InitTaskClient() {
	broker := ytask.Broker.NewRedisBroker("139.9.140.136", "6379", "19981108", 0, -1)

	backend := ytask.Backend.NewRedisBackend("139.9.140.136", "6379", "19981108", 0, -1)

	ser := ytask.Server.NewServer(
		ytask.Config.Broker(&broker),
		ytask.Config.Backend(&backend),
		ytask.Config.Debug(true),
		ytask.Config.StatusExpires(60*5),
		ytask.Config.ResultExpires(60*5),
	)

	Client = ser.GetClient()
}
