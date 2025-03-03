package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Permissions []Permission
}

type Permission struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	ModuleID uint
	Module   Module `gorm:"foreignKey:ModuleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Roles    []Role `gorm:"many2many:role_permissions;"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Users struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(100);not null"`
	LastName   string     `gorm:"type:varchar(100)"`
	Email      string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string     `gorm:"type:varchar(100);not null"`
	Businesses []Business `gorm:"many2many:user_businesses;"`
	RoleID     uint
	Role       Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Business struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null"`
	Users []Users `gorm:"many2many:user_businesses;"`
}
