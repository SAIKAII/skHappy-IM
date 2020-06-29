package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret = []byte("skHappy-IM")

var (
	TOKEN_FORMAT_ERROR        = errors.New("token的格式不正确")
	TOKEN_NOT_AVALIABLE_ERROR = errors.New("token不可用")
	TOKEN_UNKNOWN_ERROR       = errors.New("token未知错误")
	TOKEN_INVALID_ERROR       = errors.New("token非法")
)

func NewJWT(meta map[string]interface{}) (string, error) {
	mapClaims := make(jwt.MapClaims)
	for k, v := range meta {
		mapClaims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString(hmacSampleSecret)
}

// VerifyJWT 认证jwt并获取载荷
func VerifyJWT(jwtString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名算法不符合: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, TOKEN_FORMAT_ERROR
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet|jwt.ValidationErrorExpired) != 0 {
			return nil, TOKEN_NOT_AVALIABLE_ERROR
		} else {
			return nil, TOKEN_UNKNOWN_ERROR
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		meta := make(map[string]interface{})
		for k, v := range claims {
			meta[k] = v
		}
		return meta, nil
	}

	return nil, TOKEN_INVALID_ERROR
}
