package tool

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetJWT(long int, uid, name string) string {

	l := 24
	if long != 0 {
		l = long
	}
	OneDayOfHours := 60 * 60 * l
	expiresTime := time.Now().Unix() + int64(OneDayOfHours)
	//OneDayOfHours用来设置过期时间，这里设置的时一天
	//可以使用int  60*60*24

	claims := jwt.StandardClaims{
		Audience:  uid,               // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        uid,               // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "admin0001",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	//通过 StandardClaims 生成标准的载体，也就是上文提到的七个字段，其中 编号设定为 用户 id。

	var jwtSecret = []byte("yyk*2012")
	//jwtSecret 是我们设定的密钥
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过 HS256 算法生成 tokenClaims ,这就是我们的 HEADER 部分和 PAYLOAD。

	token, err := tokenClaims.SignedString(jwtSecret)
	//生成token

	if err != nil {
		fmt.Println("err:", err.Error())
		//tool.Failed(c,"err:",err.Error())
	}

	token = token
	//将 token 和 Bearer 拼接在一起，同时中间用空格隔开。
	//生成 Bearer Token
	//res:=make(map[string]interface{})
	//res["res"]=200
	//res["token"]=token
	//tool.Success(c,"success",token)

	return token

}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("yyk*2012"), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
