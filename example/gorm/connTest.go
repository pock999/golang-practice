package main

// 參考：https://juejin.im/post/6844904090196000775

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id uint `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"size:50"`
	Age int `gorm:"size:3"`
	Birthday  *time.Time
	Email string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

var db *gorm.DB

func main() {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	// 自動遷移資料結構
	// 默認表名（複數）
	// 取消默認表明複數 db.SingularTable(true)
	db.AutoMigrate(&User{})

	// 唯一索引
	db.Model(&User{}).AddUniqueIndex("name_email", "id", "name","email")

	// 新增紀錄
	db.Create(&User{Name:"bgbiao",Age:18,Email:"bgbiao@bgbiao.top"})
	db.Create(&User{Name:"xxb",Age:18,Email:"xxb@bgbiao.top"})

	var user User
	var users []User

	fmt.Println("user元素");
	db.Find(&users);
	fmt.Println(users)

	// 查詢單筆一筆數
	db.First(&user,"name = ?","bgbiao")
	fmt.Println("單筆紀錄：", user);

	// update
	db.Model(&user).Update("name","biaoge")
	fmt.Println("更新後:",user)


	// delete
	db.Delete(&user)

	// find all
	fmt.Println("Find all");
	db.Find(&users);
	fmt.Println(users);

}