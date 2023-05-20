package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

const (
	issuer = "ozline"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func (c *Claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return jwt.NewValidationError("", jwt.ValidationErrorExpired)
	}
	if !c.VerifyIssuer(issuer, true) {
		return jwt.NewValidationError("", jwt.ValidationErrorIssuer)
	}
	return nil
}

func GenerateToken(userID int64, secret []byte) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func ParseToken(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil && tokenClaims.Valid {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors^jwt.ValidationErrorExpired == 0 {
			return nil, errno.AuthorizationExpiredError
		}
		if ve.Errors^jwt.ValidationErrorIssuer == 0 {
			return nil, errno.AuthorizationFailError
		}
	}

	return nil, errno.NewErrNo(errno.AuthorizationFailedErrCode, err.Error())
}
