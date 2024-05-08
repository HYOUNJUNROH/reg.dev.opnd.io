package grpc_auth

import (
	"context"
	"errors"

	"git.dev.opnd.io/gc/backend-admin/pkg/handler/util"
	"git.dev.opnd.io/gc/backend-admin/pkg/model"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
)

type WrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *WrappedStream) Context() context.Context {
	return w.ctx
}

func NewWrappedStream(s grpc.ServerStream, ctx context.Context) grpc.ServerStream {
	return &WrappedStream{s, ctx}
}

func GetUserFromContext(ctx context.Context) (*model.Users, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	userIneterface := ctx.Value("user")
	if userIneterface == nil {
		return nil, errors.New("user ineterface is nil")
	}
	user := userIneterface.(*jwt.Token)
	return util.GetUserFromJWT(user)
}
