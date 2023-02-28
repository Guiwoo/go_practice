package goorm

import (
	"fmt"
	"gorm.io/gorm"
)

type UserA struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
	ID   int
	Name string
}

type Member struct {
	gorm.Model
	Name string
	Team []Team `gorm:"foreignkey:MemberId"`
}
type Team struct {
	gorm.Model
	Name     string
	MemberId uint
}

//func (m Member) TableName() string {
//	return "member"
//}

func Start4() {
	var members []Member
	db.Model(Member{}).Preload("Team").Find(&members)
	for _, v := range members {
		fmt.Println(v.Name, v.Team)
	}
}
