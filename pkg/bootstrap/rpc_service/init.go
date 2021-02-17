package rpc_service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	submitter "github.com/ecnuvj/vhoj_submitter/pkg/handler"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitService() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpc_config.SubmitterRpc.Address.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := submitter.NewSubmitHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	submitterpb.RegisterSubmitServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
