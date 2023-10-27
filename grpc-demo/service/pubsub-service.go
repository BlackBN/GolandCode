package service

import (
	"GolandCode/grpc-demo/protobuf"
	"context"
)

type PubSubService struct {
}

func NewPubSubService() *PubSubService {
	return &PubSubService{}
}

func (p *PubSubService) Publish(ctx context.Context, value *protobuf.String) (*protobuf.String, error) {
	return nil, nil
}
func (p *PubSubService) Subscibe(value *protobuf.String) error {
	return nil
}
