package login

import (
	"errors"

	"github.com/PandaPy/pginer/template/initialize/db"
	"github.com/PandaPy/pginer/template/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(params LoginParams) (model *models.UserModel, err error) {
	err = db.DB().Where("username = ?", params.Username).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账户名或密码错误")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(*model.Password), []byte(params.Password))
	if err != nil {
		return nil, errors.New("账户名或密码错误")
	}
	if model.Status == 0 {
		return nil, errors.New("该账户已禁用，如有疑问请联系管理员")
	}
	return model, err
}
