package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type LocalDate time.Time

func (t *LocalDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}

func (t *LocalDate) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1] // 去掉引号
	tt, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	*t = LocalDate(tt)
	return nil
}

func (t LocalDate) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type ModelMixin struct {
	CreatedBy   *int            `json:"-"`
	UpdatedBy   *int            `json:"-"`
	DeletedBy   *int            `json:"-"`
	CreatedAt   *LocalTime      `json:"created_at,omitempty"`
	UpdatedAt   *LocalTime      `json:"updated_at,omitempty"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty"`
	Description *string         `json:"description,omitempty"`
	CreatedUser *UserModelMixin `json:"created_user,omitempty" gorm:"-"`
	UpdatedUser *UserModelMixin `json:"updated_user,omitempty" gorm:"-"`
}

type UserModelMixin struct {
	ID       int     `json:"-" gorm:"primary_key"`
	Nickname *string `json:"nickname" gorm:"comment:用户昵称"`
	Username *string `json:"username" gorm:"comment:用户登录名"`
}

func (UserModelMixin) TableName() string {
	return "pd_user"
}
