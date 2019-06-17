package DataModel

import "GO_ex4/DataBase"

type NegativeComment struct {
	user_id       int
	comment       string
	comment_grade int
}

// constr
func InitNegativeComment(user_id int, comment string) *NegativeComment {
	return &NegativeComment{user_id: user_id, comment: comment, comment_grade: 0}
}

//edit Negative comment
func EditNegativeComment(c NegativeComment, newComment string) {
	c.comment = newComment
}

/*
if the user is reliable, set negative comment with high grade
means this comment will show when user look at the attached asset
*/
func SetGradeForNegativeComment(c NegativeComment) {
	user := DataBase.GetUserById(c.user_id)
	if user.Reliable_grade > 75 {
		c.comment_grade = 100
	} else {

		c.comment_grade = 10
	}
}
