package goorm

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

//func init() {
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold:             time.Second,
//			LogLevel:                  logger.Silent,
//			IgnoreRecordNotFoundError: true,
//			Colorful:                  true,
//		},
//	)
//
//	var err error
//	dsn := "root:rain45bow@tcp(127.0.0.1:3309)/guiwoo?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger:          newLogger.LogMode(logger.Info),
//		CreateBatchSize: 1000,
//	})
//
//	if err != nil {
//		panic("Failed to connect database " + err.Error())
//	}
//	fmt.Println("✅ Connected on mysql 127.0.0.1:3309")
//}

func BasicQuery() {
	var user User
	fmt.Println("Select * from user order by id limit 1")
	db.First(&user)

	fmt.Println("Select * from user limit 1")
	db.Take(&user)

	fmt.Println("select * from user order by id desc limit1")
	db.Last(&user)
}
func RetrievingObject() {
	type StrTable struct {
		Id   string `gorm:"primarykey"`
		name string
	}
	db.AutoMigrate(&StrTable{})
	var user User
	db.First(&user, 10)

	var strTable StrTable
	db.Find(&strTable, "id = ?", "holybullshit")

}
func Condition() {
	var user User
	db.Where("name = ?", "박귀우").First(&user)

	var list []User
	db.Where("name = ?", "박귀우").Find(&list)

	db.Where("id in ?", []int{1, 2, 3, 4}).Find(&list)

	db.Where("name like ?", "%woo").Find(&list)

	db.Where(map[string]interface{}{"name": "guiwoo", "age": 31}).Find(&user)

	db.Find(&list, "name like ?", "%woo")

	//Not Condition
	db.Not("name = ?", "guiwoo").First(&user)
}
func SpecificFields() {
	var list []User
	db.Select("name", "age", "id").Find(&list)
	db.Select([]string{"name", "age"}).Find(&list)
	db.Table("users").Select("COALESCE(age,?)", 31).Rows()
}
func Order() {
	var list []User
	db.Order("age desc, name").Find(&list)
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)",
			Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
}
func SubQuery() {
	var list []User
	var rst []map[string]interface{}

	query := db.Table("users").Select("count(*)").Where("age > ?", 100)
	db.Model(&User{}).Where("age > (?)", query).Find(&rst)
	fmt.Println(len(rst))
	for _, v := range rst {
		u := User{}
		if id, ok := v["id"].(uint); ok {
			u.ID = id
		}
		//fmt.Println(u)
	}
	//SELECT * FROM `users` WHERE age > (SELECT count(*) FROM `users` WHERE age > 100) AND `users`.`deleted_at` IS NULL
	//SELECT * FROM `users` WHERE age > (SELECT count(*) FROM `users` WHERE age > 100) AND `users`.`deleted_at` IS NULL

	//query = db.Select("count(name)").Where("name like ?", "%woo").Table("users")
	//db.Select("count(age) as age").Group("name").Having("AVG(age) >= (?)", query).Find(&rst)
	//fmt.Println(rst)

	db.Where("name = @name", sql.Named("name", "박귀우")).Find(&list)
	//fmt.Println(len(list))

	rows, _ := db.Table("users").Where(" name = ?", "박귀우").Rows()
	defer rows.Close()

	for rows.Next() {
		var user User
		db.ScanRows(rows, &user)
		fmt.Println(user.ID)
	}
}
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Name == "박귀우" {
		u.Name += time.Minute.String()
	}
	return
}
func Start2() {
	SubQuery()
}
