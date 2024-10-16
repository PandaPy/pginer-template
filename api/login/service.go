package login

import (
	"fmt"

	"github.com/PandaPy/pginer/template/utils/jwt"
)

func LoginService(params LoginParams) (result jwt.JWTTokenResult, err error) {
	user, err := GetUser(params)
	if err != nil {
		return result, err
	}
	result, err = jwt.GenerateJWTToken(user)
	if err != nil {
		return result, fmt.Errorf("生成 JWT 令牌失败: %v", err)
	}
	return result, nil
}
