package BusinessLogic

import (
	"ex4/DataBase"
	"ex4/DataModel"
)

func saveUserDetails(user DataModel.User) {

	DataBase.SaveUserDetailsToDB(user)
}
