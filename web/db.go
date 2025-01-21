package web

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit(filepath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&User{}, &Device{}, &Authorization{})
}

type UserType int

const (
	AdminUser UserType = iota
	NormalUser
)

type User struct {
	UserID   uint     `gorm:"primaryKey;autoIncrement"`
	UserName string   `gorm:"unique;not null"`
	UserType UserType `gorm:"not null"`
	Password string   `gorm:"not null"`
}

type DeviceType int

const (
	RotatorDevice DeviceType = iota
	SwitchDevice
	AmplifierDevice
)

type Device struct {
	DeviceID      uint       `gorm:"primaryKey;autoIncrement"`
	DeviceType    DeviceType `gorm:"not null"`
	DeviceName    string     `gorm:"not null"`
	DeviceAddress string     `gorm:"not null"`
}

type Authorization struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null"`
	DeviceID uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Device   Device `gorm:"foreignKey:DeviceID"`
}
