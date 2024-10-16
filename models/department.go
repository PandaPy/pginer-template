package models

type DepartmentModel struct {
	ModelMixin
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"  gorm:"comment:部门名称"`
}

func (DepartmentModel) TableName() string {
	return "pd_department"
}
