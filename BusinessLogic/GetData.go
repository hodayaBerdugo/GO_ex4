package BusinessLogic

import (
	"ex4/DataBase"
	"ex4/DataModel"
)

func getDataByCriteria(asset DataModel.Asset) DataModel.Post {

	return DataBase.GetDataByCriteriaFromDB(asset)
}

func getAllLocations() []string {

	return DataBase.GetAllLocationsFromDB()
}

func getUserSavedPosts(user DataModel.User) DataModel.Post {

	return DataBase.GetUserSavedPostsFromDB(user)
}

func getAllAssets() DataModel.Asset {

	return DataBase.GetAllAssetsFromDB()
}
