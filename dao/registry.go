package dao

import (
	"gorm.io/gorm"
)

// 定义 Registry 结构体，对应 doconsole_registry 表
type Registry struct {
	gorm.Model

	Name     string `gorm:"uniqueIndex"` // 对应 UNIQUE KEY `name` (`name`)
	URL      string `gorm:"not null"`    // 对应 `url` varchar(191) NOT NULL DEFAULT '#'
	NeedAuth bool   `gorm:"not null"`    // 对应 `need_auth` tinyint(1) NOT NULL DEFAULT 0
	Username string `gorm:"type:longtext"`
	Password string `gorm:"type:longtext"`
	Comment  string `gorm:"type:longtext"`
}

// 定义 RegistryDao 结构体，包含数据库连接对象
type RegistryDao struct {
	db *gorm.DB
}

// 定义 NewRegistryDao 方法，用于创建 RegistryDao 对象
func NewRegistryDao(db *gorm.DB) *RegistryDao {
	return &RegistryDao{
		db: db,
	}
}

// 定义 CreateRegistry 方法，用于创建新的 Registry 记录
func (dao *RegistryDao) CreateRegistry(registry *Registry) error {
	result := dao.db.Create(registry)
	return result.Error
}

// 定义 GetRegistryByName 方法，用于根据 name 获取 Registry 记录
func (dao *RegistryDao) GetRegistryByName(name string) (*Registry, error) {
	var registry Registry
	result := dao.db.Where("name = ?", name).First(&registry)
	if result.Error != nil {
		return nil, result.Error
	}
	return &registry, nil
}

// 定义 GetAllRegistries 方法，用于获取所有 Registry 记录
func (dao *RegistryDao) GetAllRegistries() ([]*Registry, error) {
	var registries []*Registry
	result := dao.db.Find(&registries)
	if result.Error != nil {
		return nil, result.Error
	}
	return registries, nil
}

// 定义 UpdateRegistry 方法，用于更新 Registry 记录
func (dao *RegistryDao) UpdateRegistry(registry *Registry) error {
	result := dao.db.Save(registry)
	return result.Error
}

// 定义 DeleteRegistryByName 方法，用于根据 name 删除 Registry 记录
func (dao *RegistryDao) DeleteRegistryByName(name string) error {
	result := dao.db.Delete(&Registry{}, "name = ?", name)
	return result.Error
}
