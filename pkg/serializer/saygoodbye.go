package serializer

import (
	"context"
	"fmt"
	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	db "github.com/rokane/grpc-demo/pkg/database"
)

type deleteUserWrapper struct {
	*db.DeleteUserResp
}

func (w *deleteUserWrapper) ToV1() (*v1.SayGoodbyeResponse, error) {
	return &v1.SayGoodbyeResponse{
		Message: fmt.Sprintf("Sad to see you go %s, you visited %d times", w.User.FirstName, w.Count),
	}, nil
}

func (w *deleteUserWrapper) ToV2() (*v2.SayGoodbyeResponse, error) {
	return &v2.SayGoodbyeResponse{
		Message: fmt.Sprintf("Sad to see you go %s %s, you visited %d times",
			w.User.FirstName, w.User.LastName, w.Count),
	}, nil
}

func (dbs *dbserializer) DeleteUser(ctx context.Context, criteria db.DeleteUserCriteria) (SayGoodbyeSerializer, error) {
	resp, err := dbs.storer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return &deleteUserWrapper{resp}, nil
}