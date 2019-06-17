package DataBase

import "GO_ex4/DataModel"

func GetUserById(user_id int) DataModel.User {

	return *DataModel.NewUser(1)
}
