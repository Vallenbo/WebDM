package dao

import "gorm.io/gorm"

// 在 dao 包下创建 images.go 文件，定义 Image 相关结构体和方法

// 定义 Image 结构体，对应 images 表
type Image struct {
	ImageID    int    `gorm:"primaryKey;autoIncrement"` // 对应 `image_id` int(11) NOT NULL AUTO_INCREMENT
	ImageName  string `gorm:"not null"`                 // 对应 `image_name` varchar(50) NOT NULL
	Creator    string `gorm:"not null"`                 // 对应 `creator` varchar(50) NOT NULL
	CreateTime string `gorm:"not null"`                 // 对应 `create_time` varchar(255) NOT NULL
	Remark     string `gorm:"type:text"`                // 对应 `remark` varchar(2000) DEFAULT NULL
}

// 定义 ImageDao 结构体，包含数据库连接对象
type ImageDao struct {
	db *gorm.DB
}

// 定义 NewImageDao 方法，用于创建 ImageDao 对象
func NewImageDao(db *gorm.DB) *ImageDao {
	return &ImageDao{
		db: db,
	}
}

// 定义 CreateImage 方法，用于创建新的 Image 记录
func (dao *ImageDao) CreateImage(image *Image) error {
	result := dao.db.Create(image)
	return result.Error
}

// 定义 GetImageByID 方法，用于根据 ImageID 获取 Image 记录
func (dao *ImageDao) GetImageByID(imageID int) (*Image, error) {
	var image Image
	result := dao.db.First(&image, imageID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}

// 定义 GetAllImages 方法，用于获取所有 Image 记录
func (dao *ImageDao) GetAllImages() ([]*Image, error) {
	var images []*Image
	result := dao.db.Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}
	return images, nil
}

// 定义 UpdateImage 方法，用于更新 Image 记录
func (dao *ImageDao) UpdateImage(image *Image) error {
	result := dao.db.Save(image)
	return result.Error
}

// 定义 DeleteImageByID 方法，用于根据 ImageID 删除 Image 记录
func (dao *ImageDao) DeleteImageByID(imageID int) error {
	result := dao.db.Delete(&Image{}, imageID)
	return result.Error
}
