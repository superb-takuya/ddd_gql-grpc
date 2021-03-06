package authqueryservice

import (
	"github.com/tkyatg/ddd_gql-grpc/services/auth/shared"
)

type (
	usecase struct {
		token shared.Token
	}
	genTokenRequest struct {
		userUUID string
	}
	genTokenResponse struct {
		accessToken  string
		refreshToken string
	}
	// Usecase interface
	Usecase interface {
		genToken(req genTokenRequest) (genTokenResponse, error)
	}
)

// NewUsecase function
func NewUsecase(token shared.Token) Usecase {
	return &usecase{token}
}

func (uc *usecase) genToken(req genTokenRequest) (genTokenResponse, error) {
	accessToken, refreshToken, err := uc.token.GenTokenPair(req.userUUID)
	if err != nil {
		return genTokenResponse{}, err
	}
	return genTokenResponse{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}, nil
}
