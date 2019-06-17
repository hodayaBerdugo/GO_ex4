package BusinessLogic

import (
	"GO_ex4/DataModel"
	"database/sql"
)

func saveUserDetails(user DataModel.User) {

	//connect to db
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (name, email, password, location, rooms, price) Values('{0}', '{1}', '{2}', '{3}', {4}, {5})",
		user.Name, user.Email, user.Password, user.Location, user.Rooms, user.Price)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
