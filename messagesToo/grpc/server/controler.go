package main

import (
	"context"
	"log"

	messagesToo "github.com/EvanFr/messagesToo/messagesToo"
	messagesToogrpc "github.com/EvanFr/messagesToo/messagesToo/grpc"
)

// userServiceController implements the gRPC UserServiceServer interface.
type clientServiceController struct {
	clientService messagesToo.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewClientServiceController(clientService messagesToo.Service) messagesToogrpc.ClientServiceServer {
	return &clientServiceController{
		clientService: clientService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *clientServiceController) SubscribeService(ctx context.Context, req *messagesToogrpc.Subscribe) (resp *messagesToogrpc.SubscribeEvent, err error) {
	resultMap, err := ctlr.clientService.SubscribeService(req.GetClientId())
	if err != nil {
		return
	}

	resp = &messagesToogrpc.SubscribeEvent{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalClient(&u))
	}

	log.Printf("handled GetUsers(%v)\n", req.GetClientId())
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalClient(u *messagesToo.User) *messagesToogrpc.User {
	return &messagesToogrpc.User{Id: u.ID, Name: u.Name}
}
