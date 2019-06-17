package BusinessLogic

import (
	"ex4/DataBase"
	"ex4/DataModel"
)

func addNewPost(post DataModel.Post) {

	DataBase.SaveNewPostToDB(post)
}

func deletePost(post DataModel.Post) {

	DataBase.DeletePostFromDB(post)
}

func updatePostDetails(post DataModel.Post) {

	DataBase.UpdatePostDetailsInDB(post)
}
