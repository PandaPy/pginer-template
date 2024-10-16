package models

type UserModel struct {
	ModelMixin
	ID           int             `json:"id" gorm:"primary_key;comment:主键"`
	Nickname     string          `json:"nickname" gorm:"comment:用户昵称"`
	Username     string          `json:"username" gorm:"comment:用户登录名"`
	Password     *string         `json:"password,omitempty" gorm:"comment:用户登录密码"`
	Email        *string         `json:"email,omitempty" gorm:"comment:电子邮件"`
	Mobile       *string         `json:"mobile,omitempty" gorm:"comment:手机号码"`
	Job          *string         `json:"job,omitempty" gorm:"comment:职位"`
	Sex          *int            `json:"sex,omitempty" gorm:"comment:性别"`
	Status       int             `json:"status,omitempty" gorm:"default:1;comment:账户状态"` // 0 冻结 1 正常
	IsSuperuser  *bool           `json:"is_superuser,omitempty" gorm:"comment:是否超级用户"`
	DepartmentID *int            `json:"department_id,omitempty" gorm:"comment:关联部门"`
	RoleID       *int            `json:"role_id,omitempty" gorm:"comment:关联角色"`
	Department   DepartmentModel `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
	Role         RoleModel       `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
}

func (UserModel) TableName() string {
	return "pd_user"
}
