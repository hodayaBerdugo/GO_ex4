package BusinessLogic

import (
	"database/sql"
)

func saveUserDetails(user userModel) {
	//connect to db
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (name, email, password, location, rooms, price) Values('{0}', '{1}', '{2}', '{3}', {4}, {5})", user.name, user.email, user.password, user.location, user.rooms, user.price)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
