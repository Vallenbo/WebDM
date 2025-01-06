package dao

import "gorm.io/gorm"

// 定义 FigProject 结构体，对应 fig_project 表
type FigProject struct {
	FigProjectID int    `gorm:"primaryKey;autoIncrement"` // 对应 `fig_project_id` int(11) NOT NULL AUTO_INCREMENT
	ProjectName  string `gorm:"not null"`                 // 对应 `project_name` varchar(50) NOT NULL
	MachineIP    string `gorm:"not null"`                 // 对应 `machine_ip` varchar(20) NOT NULL
	FigDirectory string `gorm:"not null"`                 // 对应 `fig_directory` varchar(200) NOT NULL
	FigParam     string `gorm:"type:text;not null"`       // 对应 `fig_param` text NOT NULL
	FigContent   string `gorm:"type:text;not null"`       // 对应 `fig_content` text NOT NULL
	CreateTime   string `gorm:"not null"`                 // 对应 `create_time` varchar(255) NOT NULL DEFAULT ''
}

// 定义 FigProjectDao 结构体，包含数据库连接对象
type FigProjectDao struct {
	db *gorm.DB
}

// 定义 NewFigProjectDao 方法，用于创建 FigProjectDao 对象
func NewFigProjectDao(db *gorm.DB) *FigProjectDao {
	return &FigProjectDao{
		db: db,
	}
}

// 定义 CreateFigProject 方法，用于创建新的 FigProject 记录
func (dao *FigProjectDao) CreateFigProject(figProject *FigProject) error {
	result := dao.db.Create(figProject)
	return result.Error
}

// 定义 GetFigProjectByID 方法，用于根据 FigProjectID 获取 FigProject 记录
func (dao *FigProjectDao) GetFigProjectByID(figProjectID int) (*FigProject, error) {
	var figProject FigProject
	result := dao.db.First(&figProject, figProjectID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &figProject, nil
}

// 定义 GetAllFigProjects 方法，用于获取所有 FigProject 记录
func (dao *FigProjectDao) GetAllFigProjects() ([]*FigProject, error) {
	var figProjects []*FigProject
	result := dao.db.Find(&figProjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return figProjects, nil
}

// 定义 UpdateFigProject 方法，用于更新 FigProject 记录
func (dao *FigProjectDao) UpdateFigProject(figProject *FigProject) error {
	result := dao.db.Save(figProject)
	return result.Error
}

// 定义 DeleteFigProjectByID 方法，用于根据 FigProjectID 删除 FigProject 记录
func (dao *FigProjectDao) DeleteFigProjectByID(figProjectID int) error {
	result := dao.db.Delete(&FigProject{}, figProjectID)
	return result.Error
}
