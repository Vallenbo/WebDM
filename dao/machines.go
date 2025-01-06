package dao

// 在 dao 包下创建 machines.go 文件，定义 Machine 相关结构体和方法

import (
	"gorm.io/gorm"
)

// 定义 Machine 结构体，对应 machines 表
type Machine struct {
	MachineID   int    `gorm:"primaryKey;autoIncrement"` // 对应 `machine_id` int(11) NOT NULL AUTO_INCREMENT
	MachineName string `gorm:"not null"`                 // 对应 `machine_name` varchar(50) NOT NULL
	MachineIP   string `gorm:"not null"`                 // 对应 `machine_ip` varchar(20) NOT NULL
	DockerPort  int    `gorm:"not null"`                 // 对应 `docker_port` int(11) NOT NULL
	IsUse       int    `gorm:"default:0"`                // 对应 `is_use` int(11) DEFAULT 0
	Remark      string `gorm:"type:text"`                // 对应 `remark` varchar(2000) DEFAULT NULL
}

// 定义 MachineDao 结构体，包含数据库连接对象
type MachineDao struct {
	db *gorm.DB
}

// 定义 NewMachineDao 方法，用于创建 MachineDao 对象
func NewMachineDao(db *gorm.DB) *MachineDao {
	return &MachineDao{
		db: db,
	}
}

// 定义 CreateMachine 方法，用于创建新的 Machine 记录
func (dao *MachineDao) CreateMachine(machine *Machine) error {
	result := dao.db.Create(machine)
	return result.Error
}

// 定义 GetMachineByID 方法，用于根据 MachineID 获取 Machine 记录
func (dao *MachineDao) GetMachineByID(machineID int) (*Machine, error) {
	var machine Machine
	result := dao.db.First(&machine, machineID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &machine, nil
}

// 定义 GetAllMachines 方法，用于获取所有 Machine 记录
func (dao *MachineDao) GetAllMachines() ([]*Machine, error) {
	var machines []*Machine
	result := dao.db.Find(&machines)
	if result.Error != nil {
		return nil, result.Error
	}
	return machines, nil
}

// 定义 UpdateMachine 方法，用于更新 Machine 记录
func (dao *MachineDao) UpdateMachine(machine *Machine) error {
	result := dao.db.Save(machine)
	return result.Error
}

// 定义 DeleteMachineByID 方法，用于根据 MachineID 删除 Machine 记录
func (dao *MachineDao) DeleteMachineByID(machineID int) error {
	result := dao.db.Delete(&Machine{}, machineID)
	return result.Error
}
