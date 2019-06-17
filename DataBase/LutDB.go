package DataBase

import (
	"database/sql"
	"ex4/DataModel"
	"log"
)

func GetDataByCriteriaFromDB(asset DataModel.Asset) DataModel.Post {
	//open db connection

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM assets a where a.rooms = {0} and price <= {1} and a.location LIKE '%{2}%' and typeId in {3};", asset.Rooms, asset.Price, asset.Location, asset.Types)

	if err != nil {
		panic(err.Error())
	}
	results.Close()
	return *DataModel.NewPost(1, "sss")
}

func GetAllLocationsFromDB() []string {
	//open db connection

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT location FROM assets")

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}
	results.Close()
	return []string{"ds", "dd"}
}

func GetUserSavedPostsFromDB(user DataModel.User) DataModel.Post {
	//open db connection

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT post FROM posts p WHERE p.userSaves = {0}", user)

	if err != nil {
		panic(err.Error())
	}

	results.Close()
	return *DataModel.NewPost(1, "sss")
}

func GetAllAssetsFromDB() DataModel.Asset {
	//open db connection

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM assets")

	if err != nil {
		panic(err.Error())
	}
	results.Close()
	return *DataModel.NewAsset(4, 1250000, 1)
}
