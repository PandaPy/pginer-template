package login

import (
	"github.com/gin-gonic/gin"

	"github.com/PandaPy/pginer/template/utils/response"
	"github.com/PandaPy/pginer/template/utils/validator"
)

func LoginController(c *gin.Context) {
	var params LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailByValidate(c, validator.GetErrorMsg(params, err))
		return
	}

	token, err := LoginService(params)
	if err != nil {
		response.FailByService(c, err.Error())
		return
	}
	response.Success(c, token)
}
