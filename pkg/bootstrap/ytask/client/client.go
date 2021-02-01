package client

import (
	ytask "github.com/gojuukaze/YTask/v2"
	"github.com/gojuukaze/YTask/v2/server"
)

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
