package BusinessLogic

import (
	"database/sql"
	"log"
)

func getDataByCriteria(asset assetModel) *Rows {
	//open db connection

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM assets a where a.rooms = {0} and price <= {1} and a.location LIKE '%{2}%' and typeId in {3};", asset.rooms, asset.price, asset.location, asset.types)

	if err != nil {
		panic(err.Error())
	}

	return results
}

func getAllLocations() *Rows {
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

	return results
}

func getUserSavedPosts(user userModel) *Rows {
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

	return results
}

func getAllAssets() *Rows {
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

	return results
}
