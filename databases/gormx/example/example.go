package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 学校的名称
type School struct {
	Name    string `gorm:"column:name;size:16;"`
	Addr    string `gorm:"column:addr;size:8"`
	Comment string `gorm:"column:comment;Comment"`
}

// Student test id, updateAt, createAt, deleteAt
type Student struct {
	ID     int          `gorm:"column:idx;autoIncrement;primaryKey"`
	Create time.Time    `gorm:"column:create;autoCreateTime"`
	Update time.Time    `gorm:"column:update;autoUpdateTime"`
	Delete sql.NullTime `gorm:"column:delete;autoDeleteTime"`
	School School       `gorm:"embedded;embeddedPrefix:school"`
	Age    int          `gorm:"column:age;size:3;check:age>=3"`
}

type Scores struct {
	ID     int `gorm:"column:idx;autoIncrement;primaryKey;"`
	Scores int `gorm:"column:score;"`
}

// 测试联合索引
type StudentScore struct {
	// ID int `gorm:"column:idx"`
	StudentID int `gorm:"index:idx_student_id"`
	ScoreID   int `gorm:"index:idx_score_id"`
}

// User 用户信息
type User struct {
	Name   string `gorm:"unique_index"`
	Number string `gorm:"unique_index"`
}

func ProductT(db *gorm.DB) {
	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	// Read
	var product Product
	db.First(&product, 1) // 根据模型主键查找
	fmt.Printf("%+v\n", product)

	db.First(&product, "code = ?", "D42") // 查找为D42的记录
	fmt.Printf("%+v\n", product)

	// Update - price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("%+v\n", product)

	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Code: "F42", Price: 200}) // 仅仅更新非0值
	fmt.Printf("%+v\n", product)

	db.Model(&product).Updates(map[string]interface{}{
		"Code": "F43", "Price": 200,
	})
	fmt.Printf("%+v\n", product)

	// Delete
	db.Delete(&product, 1)
	fmt.Printf("%+v\n", product)
}

func StudentT(db *gorm.DB) {
	db.AutoMigrate(&Student{}, &School{}, &Scores{}, &StudentScore{}, &User{})

	sch1 := &School{
		Name: "湖北工业大学",
		Addr: "武汉市武昌区",
	}
	db.Create(sch1)

	sch2 := &School{
		Name: "苏州大学",
		Addr: "江苏省苏州市",
	}
	db.Create(sch2)

	// 分数
	sdu1 := &Student{
		School: *sch1,
		Age:    18,
	}
	db.Create(&sdu1)
	fmt.Printf("%+v\n", sdu1)

	sdu2 := &Student{
		School: *sch2,
		Age:    30,
	}
	db.Create(&sdu2)
	fmt.Printf("%+v\n", sdu2)

	scroe1 := Scores{
		Scores: 100,
	}
	db.Create(&scroe1)
	fmt.Printf("%+v\n", scroe1)

	scroe2 := Scores{
		Scores: 50,
	}
	db.Create(&scroe2)
	fmt.Printf("%+v\n", scroe2)

	ss1 := StudentScore{
		// ID:        0,
		StudentID: sdu1.ID,
		ScoreID:   scroe1.ID,
	}
	db.Create(&ss1)
	fmt.Printf("%+v\n", ss1)

	ss2 := StudentScore{
		// ID:        0,
		StudentID: sdu2.ID,
		ScoreID:   scroe2.ID,
	}
	db.Create(&ss2)
	fmt.Printf("%+v\n", ss2)

	ss3 := StudentScore{
		// ID:        0,
		StudentID: sdu2.ID,
		ScoreID:   scroe2.ID,
	}
	err := db.Create(&ss3).Error
	if err != nil {
		fmt.Printf("create err: %v\n", err)
	} else {
		fmt.Printf("%+v\n", ss3)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// ProductT(db)

	StudentT(db)

	// 创建用户
	CreateUser()
}
