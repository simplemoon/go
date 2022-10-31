package main

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type MyUser struct {
	gorm.Model
	Name string
	Age int
	Birthday time.Time
}


func CreateUser() {
	user := MyUser{
		Name: "yuanzp",
		Age: 33,
		Birthday: time.Now(),
	}

	db, err := ConnectSqlite()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&MyUser{})

	result := db.Create(&user)
	println(result.Error, result.RowsAffected, user.ID)

	user.ID = 0
	db.Omit("Name", "Age", "Birthday").Create(&user)

	var users []MyUser
	for i := 5; i < 205; i++ {
		users = append(users, MyUser{
			Name:     "aaa" + strconv.Itoa(i),
			Age:      18 + i,
		})
	}

	// db.Create(users)
	// 数量为 100
	db.CreateInBatches(users, 100)
	for _, u := range users {
		println("userID: ", u.ID)
	}

	db.Model(&MyUser{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&MyUser{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
}