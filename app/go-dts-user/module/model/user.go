package model

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// https://gorm.io/docs/models.html
// bagaimana caranya kita menyambungkan
// golang struct dengan gorm
type User struct {
	// - di json, menandakan golang akan mengabaikan properti
	// - di gorm, menandakan gorm akan mengabaikan properti
	Id        uint64         `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	Name      string         `json:"name" gorm:"column:name;not null"`
	Email     string         `json:"email" gorm:"column:email;uniqueIndex;not null"`
	Dob       time.Time      `json:"dob" gorm:"column:dob;type:date;not null"`
	Delete    bool           `json:"-" gorm:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Photo     *UserPhoto     `json:"photo,omitempty" gorm:"foreignKey:UserId;reference:Id"`
	Photos    []UserPhoto    `json:"photos,omitempty" gorm:"foreignKey:UserId;reference:Id"`
	// gorm.Model
}

// menambahkan HOOK gorm
// begin transaction
// BeforeSave
// BeforeCreate
// AfterCreate
// AfterSave

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if strings.EqualFold(u.Name, "admin") {
		err = errors.New("cannot create user with name admin")
	}
	return
}

type UserPhoto struct {
	Id         uint64         `json:"id" gorm:"column:id;type:serial;primaryKey;autoIncrement;not null"`
	Url        string         `json:"url" gorm:"column:url;not null"`
	UserId     uint64         `json:"user_id" gorm:"column:user_id;type:integer;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	UserDetail *User          `json:"user_detail,omitempty" gorm:"foreignKey:UserId;reference:Id"`
	// gorm.Model
}
