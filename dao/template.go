package dao

// 在 dao 包下创建 template.go 文件，定义 Template 相关结构体和方法

import (
	"gorm.io/gorm"
)

// 定义 Template 结构体，对应 template 表
type Template struct {
	TemplateID      int    `gorm:"primaryKey;autoIncrement"` // 对应 `template_id` int(11) NOT NULL AUTO_INCREMENT
	TemplateName    string `gorm:"not null"`                 // 对应 `template_name` varchar(50) NOT NULL
	TemplateType    string `gorm:"not null"`                 // 对应 `template_type` varchar(50) NOT NULL
	TemplateContent string `gorm:"not null"`                 // 对应 `template_content` varchar(5000) NOT NULL
	CreateTime      string `gorm:"not null"`                 // 对应 `create_time` varchar(255) NOT NULL
	Remark          string `gorm:"type:text"`                // 对应 `remark` varchar(2000) DEFAULT NULL
}

// 定义 TemplateDao 结构体，包含数据库连接对象
type TemplateDao struct {
	db *gorm.DB
}

// 定义 NewTemplateDao 方法，用于创建 TemplateDao 对象
func NewTemplateDao(db *gorm.DB) *TemplateDao {
	return &TemplateDao{
		db: db,
	}
}

// 定义 CreateTemplate 方法，用于创建新的 Template 记录
func (dao *TemplateDao) CreateTemplate(template *Template) error {
	result := dao.db.Create(template)
	return result.Error
}

// 定义 GetTemplateByID 方法，用于根据 TemplateID 获取 Template 记录
func (dao *TemplateDao) GetTemplateByID(templateID int) (*Template, error) {
	var template Template
	result := dao.db.First(&template, templateID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &template, nil
}

// 定义 GetAllTemplates 方法，用于获取所有 Template 记录
func (dao *TemplateDao) GetAllTemplates() ([]*Template, error) {
	var templates []*Template
	result := dao.db.Find(&templates)
	if result.Error != nil {
		return nil, result.Error
	}
	return templates, nil
}

// 定义 UpdateTemplate 方法，用于更新 Template 记录
func (dao *TemplateDao) UpdateTemplate(template *Template) error {
	result := dao.db.Save(template)
	return result.Error
}

// 定义 DeleteTemplateByID 方法，用于根据 TemplateID 删除 Template 记录
func (dao *TemplateDao) DeleteTemplateByID(templateID int) error {
	result := dao.db.Delete(&Template{}, templateID)
	return result.Error
}
