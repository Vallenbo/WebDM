package dao

import (
	"gorm.io/gorm"
)

// 定义 Host 结构体，对应 doconsole_host 表
type Host struct {
	gorm.Model

	Name            string `gorm:"uniqueIndex"` // 对应 UNIQUE KEY `name` (`name`)
	ViaSocket       bool   `gorm:"not null"`    // 对应 `via_socket` tinyint(1) NOT NULL DEFAULT 0
	DockerEngineURL string `gorm:"default:'#'"` // 对应 `docker_engine_url` varchar(191) DEFAULT '#'
	HostIP          string `gorm:"default:'#'"` // 对应 `host_ip` varchar(191) DEFAULT '#'
	TLS             bool   `gorm:"not null"`    // 对应 `tls` tinyint(1) NOT NULL DEFAULT 0
}

// 定义 HostDao 结构体，包含数据库连接对象
type HostDao struct {
	db *gorm.DB
}

// 定义 NewHostDao 方法，用于创建 HostDao 对象
func NewHostDao(db *gorm.DB) *HostDao {
	return &HostDao{
		db: db,
	}
}

// 定义 CreateHost 方法，用于创建新的 Host 记录
func (dao *HostDao) CreateHost(host *Host) error {
	result := dao.db.Create(host)
	return result.Error
}

// 定义 GetHostByName 方法，用于根据 name 获取 Host 记录
func (dao *HostDao) GetHostByName(name string) (*Host, error) {
	var host Host
	result := dao.db.Where("name = ?", name).First(&host)
	if result.Error != nil {
		return nil, result.Error
	}
	return &host, nil
}

// 定义 GetAllHosts 方法，用于获取所有 Host 记录
func (dao *HostDao) GetAllHosts() ([]*Host, error) {
	var hosts []*Host
	result := dao.db.Find(&hosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return hosts, nil
}

// 定义 UpdateHost 方法，用于更新 Host 记录
func (dao *HostDao) UpdateHost(host *Host) error {
	result := dao.db.Save(host)
	return result.Error
}

// 定义 DeleteHostByName 方法，用于根据 name 删除 Host 记录
func (dao *HostDao) DeleteHostByName(name string) error {
	result := dao.db.Delete(&Host{}, "name = ?", name)
	return result.Error
}
