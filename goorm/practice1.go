package goorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"<-update"`
	Price uint
	extra string
}
type User struct {
	gorm.Model
	Name     string `gorm:"default:"What is wrong with you'"`
	Age      int    `gorm:"default:18"`
	Birthday time.Time
}

var db *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	var err error
	dsn := "root:rain45bow@tcp(127.0.0.1:3309)/guiwoo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:          newLogger.LogMode(logger.Info),
		CreateBatchSize: 1000,
	})

	//db = db.Session(&gorm.Session{CreateBatchSize: 1000})

	if err != nil {
		panic("Failed to connect database " + err.Error())
	}
	//db.AutoMigrate(&Product{}, &User{})
	fmt.Println("✅ Connected with mysql")
}

func Read() *Product {
	fmt.Println("Read First data of table")
	var product Product
	db.First(&product, 1)

	fmt.Println(product)
	return &product
}

func Update(prod *Product) {
	fmt.Println("Update First data to sibal => fuckyou")
	db.Model(&Product{}).Update("Code", "Holymoly")
}

func Delete() {
	fmt.Println("Delete")
	db.Delete(&Product{}, 1)
}

type User2 struct {
	a string `gorm:"<-:create"` //read and create
	b string `gorm:"<-update"`  //update only
}

func CreateRecord() {
	fmt.Println("✅ Insert Data on User")
	user := User{Name: "박귀우", Age: 31, Birthday: time.Date(1993, time.January, 20, 0, 0, 0, 0, time.UTC)}
	db.Select("Name", "Age", "Birthday").Create(&user)

	user.Name = "변ㄱ"
	db.Select("Birthday", "Age").Create(&user)
}

func generateUser() []User {
	arr := make([]User, 1000)
	for i, _ := range arr {
		arr[i] = User{
			Name: "박귀우",
			Age:  31,
			Birthday: time.Date(1993, time.January, (20+i)%32, i, i, i, i,
				time.UTC),
		}
	}
	return arr
}

func BatchInsert() {
	arr := generateUser()
	db.CreateInBatches(arr, 1000)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("😍 유저 가 만들어지고 있어요~")
	return
}

func CreateByMap() {
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "박귀우", "Age": 100},
		{"Name": "박귀우2", "Age": 200},
	})
}

type CreditCard struct {
	gorm.Model
	Number string
	UserId uint
}

type UserHasCard struct {
	Name       string
	CreditCard CreditCard
}

func CreateFromSqlExpression() {
	db.Create(&UserHasCard{
		Name:       "guiwoo",
		CreditCard: CreditCard{Number: "411234"},
	})
}

func Start() {
	//CreateFromSqlExpression()
}
