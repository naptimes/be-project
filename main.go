package main

import (
	"be-project/handler"
	"be-project/router"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dsn = "u5z3elx9bcxq6mf8:niaslitks0xa71a2@tcp(uyu7j8yohcwo35j3.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/iaws6lsi5kzfjyyo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("Success")
	}

	var dashboard handler.Dashboard
	//var roles []handler.Role
	err = db.Raw("SELECT full_name, role_description, office_longitude, office_latitude, dates, user_longitude, user_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id JOIN attendances AS d ON a.user_id = d.user_id WHERE a.user_id = 1 ORDER BY d.dates DESC LIMIT 1;").Scan(&dashboard).Error
	//err = db.Debug().Find(&roles).Error
	fmt.Println("Result: ", dashboard)
	/*for _, b := range roles {
		fmt.Println("ID: ", b.ID)
		fmt.Println("Roles: ", b.Description)
	}*/

	router.Router()
}
