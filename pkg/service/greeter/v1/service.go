package v1

import (
	"context"

	"github.com/rokane/grpc-demo/pkg/serializer"
	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

type service struct {
	ctx        context.Context
	serializer serializer.DatabaseSerializer
}

func NewService() (*v1.GreeterService, error) {
	serializer, err := serializer.NewDBSerializer()
	if err != nil {
		return nil, err
	}
	s := &service{
		ctx:        context.Background(),
		serializer: serializer,
	}
	return &v1.GreeterService{
		SayHello:   s.SayHello,
		SayGoodbye: s.SayGoodbye,
	}, nil
}
