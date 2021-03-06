package rpc

import (
	"github.com/jinzhu/gorm"
	"github.com/tkyatg/ddd_gql-grpc/services/user/adapter/hash"
	usercommandservice "github.com/tkyatg/ddd_gql-grpc/services/user/commands/userCommandService"
	"github.com/tkyatg/ddd_gql-grpc/services/user/domain"
	userqueryservice "github.com/tkyatg/ddd_gql-grpc/services/user/queries/userQueryService"
	definition "github.com/tkyatg/user-definition"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	hash := hash.NewHash()
	// user query service
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConnection)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor, hash)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)

	// user command service
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConnection)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommandservice.NewUsecase(userCommandRepository, hash)
	userCommandServer := usercommandservice.NewServer(userCommandUsecase)

	definition.RegisterUserQueryServiceServer(s.rpc, userQueryServer)
	definition.RegisterUserCommandServiceServer(s.rpc, userCommandServer)
}
