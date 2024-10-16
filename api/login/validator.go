package login

import (
	"github.com/PandaPy/pginer/template/utils/validator"
)

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (p LoginParams) GetMessages() validator.ValidatorMessages {
	return validator.ValidatorMessages{
		"username.required": "用户名称不能为空",
		"password.required": "密码不能为空",
	}
}
