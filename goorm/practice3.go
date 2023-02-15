package goorm

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

func UpdateSingle() {
	var user User
	db.First(&user).Where("id = ?", 2)
	db.Model(&User{}).Where("id = ?", 1).Update("name", "first")
	db.Save(&user)
}

func RowSql() {
	var user User
	db.Raw("select id,age,name from users where id = ?", 3).Scan(&user)
	//fmt.Println(user)
	db.Exec("update users set name = @name where id = @id", sql.Named("name", "아씨에어팟1세대 시켰네 개열받게"), sql.Named("id", 3))

	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 3).Statement
	stmt.SQL.String()

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id > ?", 20).Group("name").Find(&[]User{})
	})
	fmt.Println(sql)

}

func Start3() {
	RowSql()
}
