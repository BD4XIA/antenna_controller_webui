package main

type User struct {
	UserID   uint   `gorm:"primaryKey;autoIncrement"` // 用户ID，主键，自增
	Username string `gorm:"unique;not null"`          // 用户名，唯一且不为空
	Password string `gorm:"not null"`                 // 密码，不为空
}

type Device struct {
	DeviceID      uint   `gorm:"primaryKey;autoIncrement"` // 设备ID，主键，自增
	DeviceType    string `gorm:"not null"`                 // 设备类型，不为空
	DeviceName    string `gorm:"not null"`                 // 设备名称，不为空
	DeviceAddress string `gorm:"not null"`                 // 设备地址，不为空
}

type Authorization struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"` // 授权ID，主键，自增
	UserID   uint   `gorm:"not null"`                 // 用户ID，外键，不为空
	DeviceID uint   `gorm:"not null"`                 // 设备ID，外键，不为空
	User     User   `gorm:"foreignKey:UserID"`        // 关联 User 表
	Device   Device `gorm:"foreignKey:DeviceID"`      // 关联 Device 表
}
