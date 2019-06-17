package BusinessLogic

import (
	"GO_ex4/DataModel"
	"database/sql"
)

func addNewPost(post DataModel.Post) {
	//connect to db
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO posts Values('{0}', '{1}', '{2}', '{3}', {4})", post.Name, post.Likes, post.Comments, post.Rank, post.Date)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func deletePost(post DataModel.Post) {
	//connect to db
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	delete, err := db.Query("DELETE FROM posts p WHERE p.postId = {0}", post.Id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func updatePostDetails(post DataModel.Post) {
	//connect to db
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	update, err := db.Query("UPDATE TABLE posts ('{0}', '{1}', '{2}', {3}, {4})", post.Name, post.Likes, post.Comments, post.Rank, post.Date)

	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}
