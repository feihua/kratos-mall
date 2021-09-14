package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"time"
)

var jwtSecret = []byte("test")
var issuer = "test"
var expire time.Duration = 24

type Claims struct {
	UserId   int64
	Username string
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire * time.Hour)

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func DestroyToken(token string) bool {
	return true
}

// AuthMiddleware 根据header读取jwt token
func AuthMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var token string
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				authorization := tr.RequestHeader().Get("Authorization")
				if authorization != "" {
					token = authorization
				}
			}

			md, ok := metadata.FromIncomingContext(ctx)
			if ok && len(md.Get("Authorization")) > 0 {
				token = md.Get("Authorization")[0]
			}

			if token == "" {
				return nil, errors.Unauthorized("token不得为空", "token不得为空")
			}
			ctx = context.WithValue(ctx, "Authorization", token)

			// 解析token
			claims, err := ParseToken(token)
			if err != nil {
				return nil, errors.Unauthorized("token解析错误", "token解析错误")
			} else if time.Now().Unix() > claims.ExpiresAt {
				return nil, errors.Unauthorized("token已过期", "token已过期")
			}
			ctx = context.WithValue(ctx, "x-md-global-uid", claims.UserId)
			return handler(ctx, req)
		}
	}
}
