package business

import (
	"context"
	proto "gomicro/server/pb"
)

type Hello struct{}

func (h Hello) Goodbye(c context.Context, request *proto.Request, response *proto.Response) error {
	response.Message = "Goodbye " + request.Name
	return nil
}

func (h Hello) Say(context context.Context, request *proto.Request, response *proto.Response) error {
	response.Message = "Hello"
	return nil
}