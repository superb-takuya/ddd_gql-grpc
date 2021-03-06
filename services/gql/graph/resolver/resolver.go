package resolver

import (
	"context"

	authdefinition "github.com/tkyatg/auth-definition"
	authserviceaccessor "github.com/tkyatg/ddd_gql-grpc/services/gql/adapter/rpc/authServiceAccessor"
	userserviceaccessor "github.com/tkyatg/ddd_gql-grpc/services/gql/adapter/rpc/userServiceAccessor"
	"github.com/tkyatg/ddd_gql-grpc/services/gql/graph/generated"
	"github.com/tkyatg/ddd_gql-grpc/services/gql/shared"
	userdefinition "github.com/tkyatg/user-definition"

	"google.golang.org/grpc"
)

type resolver struct {
	userServiceAccessor userserviceaccessor.ServiceAccessor
	authServiceAccessor authserviceaccessor.ServiceAccessor
}

// NewResolver function
func NewResolver(ctx context.Context, env shared.Env) generated.ResolverRoot {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// user service
	userConn, err := grpc.DialContext(ctx, env.GetUserServerName()+":"+env.GetUserServerPort(), opts...)
	if err != nil {
		panic(err)
	}
	userQueryClient := userdefinition.NewUserQueryServiceClient(userConn)
	userCommandClient := userdefinition.NewUserCommandServiceClient(userConn)
	userServiceAccessor := userserviceaccessor.NewUserServiceAccessor(userQueryClient, userCommandClient)

	// auth service
	authConn, err := grpc.DialContext(ctx, env.GetAuthServerName()+":"+env.GetAuthServerPort(), opts...)
	if err != nil {
		panic(err)
	}
	authQueryClient := authdefinition.NewAuthQueryServiceClient(authConn)

	authServiceAccessor := authserviceaccessor.NewAuthServiceAccessor(authQueryClient)

	return &resolver{
		userServiceAccessor,
		authServiceAccessor,
	}
}
