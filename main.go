package main

import (
	"be-project/database"
	"be-project/models"
	"be-project/router"
	"fmt"
)

func main() {
	// create db connection example
	db := database.ConnectDB()

	var dashboard models.Dashboard

	// query for collecting /dashboard data
	db.Raw("SELECT full_name, role_description, office_longitude, office_latitude, dates, user_longitude, user_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id JOIN attendances AS d ON a.user_id = d.user_id WHERE a.user_id = 1 ORDER BY d.dates DESC LIMIT 1;").Scan(&dashboard)

	// check query result
	fmt.Println("Result: ", dashboard)

	router.Router()
}
