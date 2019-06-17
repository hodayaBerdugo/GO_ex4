package DataModel

type Comment struct {
	user_id int
	comment string
}

func InitComment(user_id int, comment string) *Comment {
	return &Comment{user_id: user_id, comment: comment}
}

//edit comment
func EditComment(c Comment, newComment string) {
	c.comment = newComment
}
