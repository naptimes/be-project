package main

import (
	"be-project/router"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*type Role struct {
	ID         int    `json:"role_id"`
	Descrption string `json:"role_description"`
}*/

func main() {
	/*var dsn = "u5z3elx9bcxq6mf8:niaslitks0xa71a2@tcp(uyu7j8yohcwo35j3.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/iaws6lsi5kzfjyyo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("Success")
	}

	var roles []Role
	db.Find(&roles)
	fmt.Println(roles.ID[0])*/

	router.Router()
}
