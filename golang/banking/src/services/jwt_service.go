package services

import (
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

type JWTService struct {
	PublicKey jwk.Key
	Log       *zap.Logger
}

func (service JWTService) VerifyToken(request *http.Request) (jwt.Token, error) {
	verifiedToken, err := jwt.ParseRequest(
		request,
		jwt.WithKey(jwa.RS512, service.PublicKey),
		jwt.WithVerify(true),
		jwt.WithValidate(true),
		jwt.WithHeaderKey("Authorization"),
	)

	if err != nil {
		service.Log.Error("Error verifying jwt", zap.Error(err))
		return nil, err
	}
	return verifiedToken, nil
}
