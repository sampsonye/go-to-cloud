package repositories

// Org 组织
// 组织拥有基础设施和用户
type Org struct {
	Model
	Name  string  `json:"name" gorm:"column:name;not null;"` // 组织名称
	Users []*User `gorm:"many2many:orgs_users_rel;"`
}

func (m *Org) TableName() string {
	return "org"
}
