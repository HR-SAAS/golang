package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"hr-saas-go/upload-web/global"
	"hr-saas-go/upload-web/models"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登录后操作",
			})
			ctx.Abort()
			return
		}
		j := NewJWT()
		token = strings.Split(token, " ")[1]
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "token 已经过期",
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登录后操作",
			})
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}

func ReadFromJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" {
			j := NewJWT()
			token = strings.Split(token, " ")[1]
			claims, err := j.ParseToken(token)
			if err == nil {
				ctx.Set("claims", claims)
				ctx.Set("userId", claims.ID)
			}
		}
		ctx.Next()
	}
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYe  = errors.New("token not active ye")
	TokenFormatError = errors.New("token err")
	TokenInvalid     = errors.New("could'n handle this token")
)

type JWT struct {
	SingalKey []byte
}

func NewJWT() *JWT {
	return &JWT{[]byte(global.Config.JwtConfig.Key)}
}
func (j *JWT) ParseToken(token string) (*models.CustomClaims, error) {
	t, err := jwt.ParseWithClaims(token, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SingalKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenFormatError
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYe
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if t != nil {
		if claims, ok := t.Claims.(*models.CustomClaims); ok && t.Valid {
			return claims, nil
		}

	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SingalKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
		return j.CreateToken(*claims)
	}
	return "", errors.New("refresh error")
}

func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SingalKey)
}
