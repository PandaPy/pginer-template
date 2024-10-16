package models

type RoleModel struct {
	ModelMixin
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"  gorm:"comment:角色名称"`
}

func (RoleModel) TableName() string {
	return "pd_role"
}
