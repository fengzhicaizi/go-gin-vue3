package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(auth *system.SysAuth) (token string, errR error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	jwtSecret := []byte(GVA_VP.GetString("app.JwtSecret"))

	claims := Claims{
		Id:       auth.Id,
		Username: auth.Username,
		Password: auth.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-blog",
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, errR = tokenClaim.SignedString(jwtSecret)

	return
}

// ParseToken 解析token
func ParseToken(token string) (*system.SysAuth, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, errors.New(`token should begin with 'Bearer'`)
	}

	arr := strings.Split(token, " ")
	token = arr[1]

	jwtSecret := []byte(GVA_VP.GetString("app.JwtSecret"))
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if jwtToken != nil {
		if Claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
			auth := &system.SysAuth{
				Id: Claims.Id,
			}
			return auth, nil
		}
	}

	return nil, errors.New("token解析失败")
}
