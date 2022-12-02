package JWt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 1
const FokenExpireDuration = time.Hour * 24

var MySecret = []byte("长亭外古道边")
var FreshToken string = "default"

type MyCalims struct {
	User_ID  int64  `json:"user_id"`
	Password string `json:"password"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return MySecret, nil
}

// 生成 access_token 和 refresh_token
func GenToken2(user_id int64, username, password string) (aToken, rToken string, err error) {
	calims := MyCalims{
		User_ID:  user_id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calims)
	// 生成 aToken
	aToken, err = token.SignedString(MySecret)

	// rToken 不需要存储任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(FokenExpireDuration).Unix(), // 过期时间
		Issuer:    "my-project",                               // 签发人
	}).SignedString(MySecret)
	return aToken, rToken, nil
}
func ParasToken2(access_token string) (claims *MyCalims, err error) {
	var token *jwt.Token
	claims = new(MyCalims)
	token, err = jwt.ParseWithClaims(access_token, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid { // token 是否有效
		err = errors.New("invalid token")
	}
	return claims, nil
}

func RefreshToken(aToken, rToken string) (newToken, newrToken string, err error) {
	// 第一步 : 判断 rToken 格式对的，没有过期的
	if _, err := jwt.Parse(rToken, keyFunc); err != nil {
		return "", "", err
	}

	// todo 第二步：从旧的 aToken 中解析出 cliams 数据   过期了还能解析出来吗
	var claims MyCalims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当 access token 是过期错误，并且 refresh token 没有过期就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken2(claims.User_ID, claims.Username, claims.Password)
	}
	return "", "", err
}
